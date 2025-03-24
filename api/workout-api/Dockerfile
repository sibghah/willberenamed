FROM golang:1.24.1 AS builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

# Copy the source code into the container
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o api .


# Stage 2: Create the final minimal image for production
FROM alpine:latest

#install necessary dependencies
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]