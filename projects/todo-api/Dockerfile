FROM golang:1.26.2-bookworm

WORKDIR /app

COPY . .

RUN go build -o /app/todo-api .

CMD ["/app/todo-api"]