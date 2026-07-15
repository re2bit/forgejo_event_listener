FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dynamic_runner_subsystem main.go

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/dynamic_runner_subsystem /dynamic_runner_subsystem

EXPOSE 8080

CMD ["/dynamic_runner_subsystem", "serve"]
