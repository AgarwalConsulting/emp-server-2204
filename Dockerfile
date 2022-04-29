FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED 0
RUN go build -o app ./cmd/server

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
ENV PORT 8000
ENV DB_URL "postgres://host.docker.internal:5432/emp-demo?sslmode=disable"
CMD [ "/app/app" ]
