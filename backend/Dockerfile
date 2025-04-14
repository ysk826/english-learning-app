FROM golang:1.23-alpine

WORKDIR /app

# Install air for hot reload
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum files
COPY go.mod go.sum* ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Copy air configuration
COPY .air.toml .

EXPOSE 8080

# Use air for hot reload in development
CMD ["air", "-c", ".air.toml"]