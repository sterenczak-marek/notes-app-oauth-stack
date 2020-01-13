FROM golang:1.13.0-alpine3.10

ENV CGO_ENABLED 0

RUN apk --no-cache add \
    git bash

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download && \
    go mod verify && \
    go get github.com/cosmtrek/air/... \
	github.com/go-delve/delve/cmd/dlv

COPY ./ ./app/

CMD ["air"]
