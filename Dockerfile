FROM golang:1.10

WORKDIR /go/src/app
COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN go install -v ./...

CMD ["docker-events-webhook"]