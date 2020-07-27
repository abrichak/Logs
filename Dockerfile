# Start from golang base image
FROM golang:1.14.4-alpine3.12 as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Install Alpine build tools package
RUN apk add build-base

# Set the current working directory inside the container
WORKDIR /api

RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/swaggo/swag/cmd/swag

#Command to run the executable
CMD swag init \
  && CompileDaemon --build="go build main.go"  --command="./main" --color
