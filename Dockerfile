FROM golang:1.25-alpine3.23

LABEL authors="chunan"

WORKDIR /app

# Avoid VCS errors inside containers
ENV GOFLAGS="-buildvcs=false"

# Install air
RUN go install github.com/air-verse/air@latest

# Copy dependency files first (better cache usage)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

EXPOSE 6969

CMD ["air", "-c", ".air.toml"]
