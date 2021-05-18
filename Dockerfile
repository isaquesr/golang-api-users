FROM golang:1.16-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/isaque/Golang-MongoDB-CRUD

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

#RUN swag init -g ./main.go --output docs/echosimple
RUN go build -o ./workdir/svc-users-go .


# This container exposes port 8080 to the outside world
EXPOSE 8000

# Run the binary program produced by `go install`
CMD ["./workir/svc-users-go"]