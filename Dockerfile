FROM golang:latest

WORKDIR /app

COPY /src/GoChat /app

RUN go build -o myapp

EXPOSE 8080

CMD ["./myapp"]