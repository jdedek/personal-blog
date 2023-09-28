FROM golang:1.21

WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

RUN go build -o main .

# Expose the port on which your application will run
EXPOSE 42069

# Command to run the executable
CMD ["./main"]

