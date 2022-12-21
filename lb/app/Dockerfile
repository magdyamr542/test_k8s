# Build
FROM golang:1.19 as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /app

## Deploy
FROM gcr.io/distroless/base-debian10
COPY --from=builder /app /
USER nonroot:nonroot
ENTRYPOINT [ "/app" ] 