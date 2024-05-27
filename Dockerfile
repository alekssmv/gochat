FROM golang:latest

WORKDIR /app

COPY . .

RUN cd src/App && go build -o myapp

EXPOSE 8080

CMD ["./src/App/myapp"]