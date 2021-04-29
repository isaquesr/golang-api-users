FROM golang:1.13-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/isaque/Golang-MongoDB-CRUD

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go get -u .
#RUN swag init -g ./main.go --output docs/echosimple
RUN go build -o ./out/svc-users-go .


# This container exposes port 8080 to the outside world
EXPOSE 8081

# Run the binary program produced by `go install`
CMD ["./out/svc-users-go"]