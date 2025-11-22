# --------------- BUILD -----------------
FROM golang:1.25.4-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /out/server ./cmd/server/main.go

# --------------- RUN -------------------
FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=build /out/server .

EXPOSE 8080

CMD ["./server"]