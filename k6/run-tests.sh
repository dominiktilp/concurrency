#!/bin/bash

BASE_URL=$1

if [ -z "${BASE_URL}" ]; then
  BASE_URL=http://192.168.68.103:9001/
fi;


TAG=$2

if [ -z "${TAG}" ]; then
  TAG=$(date +%s)
fi;


docker run -e "URL=${BASE_URL}" -v "$PWD:/work" -i loadimpact/k6 \
run --out json=./work/reports/${TAG}_.json - <index.js > ./reports/${TAG}

docker run -e "URL=${BASE_URL}fib/20" -v "$PWD:/work" -i loadimpact/k6 \
run --out json=./work/reports/${TAG}_fib.json - <index.js > ./reports/${TAG}_fib

docker run -e "URL=${BASE_URL}sleep/200" -v "$PWD:/work" -i loadimpact/k6 \
run --out json=./work/reports/${TAG}_sleep.json - <index.js > ./reports/${TAG}_sleep