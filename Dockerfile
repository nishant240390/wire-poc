FROM golang:latest

LABEL maintainer="Nishant Agrawal <nishantagrawal009@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
