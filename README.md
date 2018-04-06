docker-events-webhook [![Go Report Card](https://goreportcard.com/badge/github.com/matyunin/docker-events-webhook)](https://goreportcard.com/report/github.com/matyunin/docker-events-webhook)
=====================

Docker events webhook tool allows you to intercept the docker engine events and pipeline them
to specified webhook uri via the HTTP POST request.

#### Settings

Available environment variables:

| Variable name | Description |
| ------------- | ------------- |
|DOCKER_CERT_PATH| |
|DOCKER_TLS_VERIFY| |
|DOCKER_HOST| By default: `unix:///var/run/docker.sock` |
|DOCKER_API_VERSION| |
|WEBHOOK_URI| Webhook URI to handle request. If empty all events will be passed to the `STDOUT`. |

#### Example

Connect to local docker socket and stream events to __requestbin__ endpoint:

```bash
dep ensure
go build
DOCKER_HOST=unix:///var/run/docker.sock WEBHOOK_URI=http://requestbin.fullcontact.com/11956ck1 ./docker-events-webhook
```

Above example using docker image:

```bash
docker run -d -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_HOST=unix:///var/run/docker.sock -e WEBHOOK_URI=http://requestbin.fullcontact.com/11956ck1 matyunin/docker-events-webhook
```

On other side you will receive a messages like this:

```json
{
  "status": "push",
  "id": "matyunin/docker-events-webhook:latest",
  "Type": "image",
  "Action": "push",
  "Actor": {
    "ID": "matyunin/docker-events-webhook:latest",
    "Attributes": {
      "name": "matyunin/docker-events-webhook"
    }
  },
  "time": 1523027703,
  "timeNano": 1523027703594899700
}
```
