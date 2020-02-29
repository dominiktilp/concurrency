FROM golang:alpine

WORKDIR /var/www/go

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o index ./index.go

ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}

CMD ["./index"]