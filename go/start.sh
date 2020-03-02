#!/bin/bash

#docker build -t concurrency:go . && \
docker run -it -p 9001:8080 -e CONCURRENCY_REPO="http://sad_roentgen:8080" -m 256m --memory-swap 256m --cpus 0.5 --link sad_roentgen:sad_roentgen --name go concurrency:go