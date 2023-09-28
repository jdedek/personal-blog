# Use an official Go runtime as a parent image
FROM golang:1.21.1-alpine3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port on which your application will run
EXPOSE 42069

# Command to run the executable
CMD ["./main"]

