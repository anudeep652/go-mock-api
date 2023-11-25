FROM golang:latest

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY *.go ./
COPY Makefile ./ 

# Build
RUN CGO_ENABLED=0 GOOS=linux make build

# Run
CMD ["bin/go-mock-api"]