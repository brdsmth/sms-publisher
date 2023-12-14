# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o sms-publisher

# Expose the port that your Publisher service listens on
EXPOSE 8080

# Define the command to run your Publisher service
CMD ["./sms-publisher"]
