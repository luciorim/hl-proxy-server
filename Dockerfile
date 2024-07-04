FROM golang:1.22.0

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "/app/cmd", "."]