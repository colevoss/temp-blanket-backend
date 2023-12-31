FROM golang:1.21.1-alpine as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/local/bin/app /usr/src/app/cmd/tempblanket/main.go

FROM alpine

ENV PORT=8080
ENV ENV=production

RUN apk add --no-cache tzdata

WORKDIR /

COPY --from=builder /usr/local/bin/app /bin/app

# EXPOSE 8080

# USER nonroot:nonroot

CMD ["/bin/app"]
