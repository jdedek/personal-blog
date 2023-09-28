FROM golang:1.21

WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .


# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Build the Go application
RUN go build -o main .

# Expose the port on which your application will run
EXPOSE 42069

# Command to run the executable
CMD ["./main"]

