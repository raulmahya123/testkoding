# Use the official Golang image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod file to the working directory
COPY go.mod .

# Copy the main.go file to the working directory
COPY main.go .


# Copy the entire golangsidang directory to the working directory
COPY golangsidang/ ./golangsidang/

# Download dependencies defined in go.mod
RUN go mod download github.com/andybalholm/brotli

# Perform any additional tidy up of dependencies
RUN go mod tidy


# Set environment variables for database connection
ENV DB_HOST="pg-9b31fdb-raulmahya11-dca4.e.aivencloud.com" \
    DB_PORT="19683" \
    DB_USER="avnadmin" \
    DB_PASSWORD="AVNS_ueyJXedpqkj8K4DlNIj" \
    DB_NAME="defaultdb" \
    DB_SSL="disable" \
    JWT_SECRET="your_dynamic_secret_key"

# Build the Go application and output the binary to the bin directory
RUN go build -o bin/myapp .

# Expose port 3000 to the outside world
EXPOSE 3000

# Set the entry point for the container to run the compiled binary
ENTRYPOINT [ "/app/bin/myapp" ]

# Set the default command for the container to run the API server
CMD ["./main"]

