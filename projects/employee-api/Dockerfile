FROM golang:1.26.2-bookworm

WORKDIR /home-3

COPY . .

RUN go mod tidy

CMD ["go", "run", "main.go"]
