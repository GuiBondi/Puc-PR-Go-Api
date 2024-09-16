# Step 1: Use an official Golang image as the builder
FROM golang:1.23-alpine AS builder

# Step 2: Set the Current Working Directory inside the container
WORKDIR /app

# Step 3: Copy go.mod and go.sum files
COPY go.mod ./

# Step 4: Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Step 5: Copy the source code into the container
COPY . .

# Step 6: Build the Go app
RUN go build -o main .

# Step 7: Use a minimal image for running the application
FROM alpine:latest

# Step 8: Set the working directory in the second stage
WORKDIR /app

# Step 9: Copy the executable from the builder stage
COPY --from=builder /app/main .

# Step 10: Expose the port the app runs on
EXPOSE 8080

# Step 11: Command to run the executable
CMD ["./main"]
