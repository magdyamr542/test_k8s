# Build
FROM golang:1.19 as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o exec

## Deploy
FROM alpine:latest
COPY --from=builder /app/exec /exec
CMD [ "/exec" ]