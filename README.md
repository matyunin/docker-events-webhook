docker-events-webhook
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
docker run -d \
    -e DOCKER_HOST=unix:///var/run/docker.sock \ 
    -e WEBHOOK_URI=http://requestbin.fullcontact.com/11956ck1 \
    matyunin/docker-events-webhook
```