# Uses Alpine image
FROM golang:1.18-alpine
WORKDIR /app

# Copy the mod and download dependencies
COPY go.* ./
RUN go mod download

# Build the project
COPY . .
RUN go build -o ./dist .

# Run the executable
CMD ./dist