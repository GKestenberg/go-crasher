# Dockerfile
FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server main.go

EXPOSE 3000 
CMD ["./server"]
