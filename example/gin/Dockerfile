FROM golang:1.17 as builder

WORKDIR /go/src/app
COPY *.go .

# COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

RUN go get -d -v ./...
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app /
CMD ["/app"]