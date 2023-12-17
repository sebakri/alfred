FROM golang:1.21 AS builder
COPY . /app
WORKDIR /app
RUN go build -o bin/alfred ./cmd/alfred

FROM scratch
COPY --from=builder /app/bin/alfred /alfred
ENTRYPOINT ["/alfred"]
