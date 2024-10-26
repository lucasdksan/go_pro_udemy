FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o ./server ./cmd/api/

FROM scratch

WORKDIR /bin

COPY --from=builder /app/server /bin/server

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY .env /bin/.env

COPY cer.cer /bin/cer.cer
COPY cer.key /bin/cer.key

CMD ["/bin/server"]