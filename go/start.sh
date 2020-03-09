#!/bin/bash

docker build -t concurrency:go . && \
docker run -it -p 9001:8080 -m 256m --memory-swap 256m --cpus 0.5 concurrency:go
#docker run -it -p 9001:8080 -m 1024m --memory-swap 1024m --cpus 2 concurrency:go
