# filepath: /Users/arunmittal/Desktop/meeting-scheduler/Dockerfile
# Stage 1: Build
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build for Linux explicitly (required for Alpine final image)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o meeting-scheduler ./cmd/server

# Stage 2: Run
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/meeting-scheduler .

# Ensure it's executable
RUN chmod +x ./meeting-scheduler

EXPOSE 8080
CMD ["./meeting-scheduler"]
