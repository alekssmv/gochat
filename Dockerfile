FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o myapp src/main.go

EXPOSE 8080

CMD ["./myapp"]