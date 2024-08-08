FROM golang:1.20-alpine AS builder

# Set environment variable
ENV APP_NAME processout-helloworld
ENV CMD_PATH main.go

# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Build application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# Run Stage
FROM alpine:latest

# Set environment variable
ENV APP_NAME processout-helloworld

# Copy only required data into this image
COPY --from=builder /$APP_NAME .

# Expose application port
EXPOSE 8080

# Start app
CMD ./$APP_NAME
