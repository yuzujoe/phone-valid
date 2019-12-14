FROM golang:1.13 as builder

ENV GO111MODULE=on

WORKDIR /go/src

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

ENV CGO_ENABLED=0

RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o build ./cmd/main.go

FROM alpine

RUN apk update \
  && apk --no-cache add tzdata \
  && apk add --no-cache ca-certificates

ENV GIN_MODE=release
COPY --from=builder /go/src/build /go/src/build

ENTRYPOINT ["go/src/build"]
