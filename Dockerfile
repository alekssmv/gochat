#Backend
FROM golang:latest as backend

WORKDIR /app

COPY src/Backend .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gochat .

#Frontend
FROM node:latest as frontend

WORKDIR /app

COPY src/Frontend .

RUN npm install && npm run build

#Final
FROM alpine:latest

WORKDIR /app

COPY --from=backend /app/gochat /app/gochat

COPY --from=frontend /app/dist /app/dist

EXPOSE 8080

# Run the binary
CMD ["./gochat"]