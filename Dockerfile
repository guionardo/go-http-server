FROM golang:1.16

WORKDIR /go/src/github.com/guionardo/go-http-server/
COPY go.* ./
RUN go mod download
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-http-server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ARG GIN_MODE
COPY --from=0 /go/src/github.com/guionardo/go-http-server/go-http-server .
CMD ["./go-http-server"]