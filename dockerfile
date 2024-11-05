# Use Go 1.18 instead of 1.17
FROM golang:1.18-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# List the contents of the working directory for debugging
RUN ls -la

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]