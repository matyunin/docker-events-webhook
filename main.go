package main

import (
	"fmt"
	"os"

	"context"
	"github.com/docker/docker/api/types"
	"github.com/kelseyhightower/envconfig"
	"github.com/moby/moby/client"
	"encoding/json"
	url "net/url"
	"bytes"
	"net/http"
	"time"
)

// TODO: Implement worker pool for HTTP POST jobs
// TODO: Incapsulate logic to separate methods
// TODO: Handle OS signals properly
// TODO: Use logrus for logging
// TODO: Unit testing (https://travis-ci.org)
// TODO: Go report card (https://goreportcard.com/)

type Config struct {
	WebhookUri string `envconfig:"webhook_uri"`
}

func main() {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}

	uri, err := url.ParseRequestURI(c.WebhookUri)
	if err != nil {
		panic(err)
	}

	h := &http.Client{
		Timeout: 2 * time.Second,
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// TODO: Properly get and handle event stream options (since, until, filter)
	events, _ := cli.Events(context.Background(), types.EventsOptions{})

	for e := range events {
		//fmt.Printf("%+v\n", e)

		j, err := json.Marshal(e)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot marshal event to json")
			continue
		}

		req, err := http.NewRequest("POST", uri.String(), bytes.NewBuffer(j))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create HTTP POST request: %s", err.Error())
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := h.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to send HTTP POST request: %s", err.Error())
		}
		defer resp.Body.Close()

		fmt.Printf("%+v\n", string(j[:]))
	}
}
