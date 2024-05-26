FROM golang:latest

WORKDIR /app

COPY /src/App /app

RUN go build -o myapp

EXPOSE 8080

CMD ["./myapp"]