#!/bin/bash

docker build -t concurrency:data . && \
docker run -it -p 9999:8080 -m 256m --memory-swap 256m --cpus 0.5 concurrency:data