# Use an official Golang runtime as a parent image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o hospitalApp

# Expose the port the application will run on
EXPOSE 8080

# Run the application
CMD ["./hospitalApp"]