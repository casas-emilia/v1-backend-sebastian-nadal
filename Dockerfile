# Build stage
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application and the migration tool
RUN CGO_ENABLED=0 GOOS=linux go build -o main . && CGO_ENABLED=0 GOOS=linux go build -o migrate ./migrate/migrate.go

# Final stage
FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y ca-certificates tzdata && rm -rf /var/lib/apt/lists/*
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .
EXPOSE 8080
# Create a startup script
RUN echo '#!/bin/sh' > start.sh && \
    echo './migrate' >> start.sh && \
    echo './main' >> start.sh && \
    chmod +x start.sh

CMD ["./start.sh"]


