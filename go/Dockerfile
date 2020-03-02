FROM golang:1.14.0-stretch

RUN \
  apt-get update \
  && apt-get -y install git \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .
RUN go build -o ./bin/concurrency ./cmd/*

ARG PORT=8080
ENV CONCURRENCY_PORT=${PORT}
EXPOSE ${PORT}

ARG REPO="http://localhost:9999"
ENV CONCURRENCY_REPO=${REPO}

ENTRYPOINT ["./bin/concurrency"]

CMD ["run"]
