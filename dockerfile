
FROM golang:1.23-alpine AS builder

WORKDIR /src/api

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# --------------> production image
FROM gcr.io/distroless/base-debian10
COPY --from=builder /src/api/main /
CMD ["/main"]
