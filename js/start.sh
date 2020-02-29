#!/bin/bash

docker build -t concurrency:js . && \
docker run -it -p 9001:8080 -m 256m --memory-swap 256m --cpus 0.5 concurrency:js