FROM golang:latest

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . ./


# Build
RUN CGO_ENABLED=0 GOOS=linux
RUN go install github.com/mitranim/gow@latest

# Run
# CMD ["bin/go-mock-api"]
CMD ["gow","run","main.go"]