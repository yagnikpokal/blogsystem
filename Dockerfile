# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files (go.mod and go.sum)
COPY go.mod .
COPY go.sum .

# Download and cache Go modules
RUN go mod download

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o main ./

# Expose the port if needed (if your Go program listens on a specific port)
# EXPOSE 8080

# Command to run the executable
CMD ["./main"]
