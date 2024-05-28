# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install fresh for live reloading
RUN go install github.com/pilu/fresh@latest

# Command to run the executable
CMD ["fresh"]
