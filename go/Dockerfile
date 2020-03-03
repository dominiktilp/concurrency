FROM golang:1.14.0-stretch

RUN \
  apt-get update \
  && apt-get -y install git \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .
RUN go build -o ./bin/index ./index.go

ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}

ARG DATA_HOST="http://localhost:9999"
ENV DATA_HOST=${DATA_HOST}

CMD ["./bin/index"]
