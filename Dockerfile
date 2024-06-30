# syntax=docker/dockerfile:1
FROM golang:1.22.0-alpine AS build

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

RUN go build -ldflags="-s -w" -o ./app ./cmd/server/main.go

# Second stage: lightweight runtime environment
FROM alpine:3.19.1 AS runtime

# Set the working directory inside the container
WORKDIR /app

# Copy environment variable files into the image, if needed
COPY ./env/*.env ./env/

# Copy the binary executable from the build stage
COPY --from=build /app/app /app/

# Expose the port on which your application listens, if needed
EXPOSE 50051

# Specify the command to run when the container starts
ENTRYPOINT ["/app/app"]
