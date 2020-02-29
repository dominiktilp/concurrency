# Data for concurrency test
This directory faker node server with data for concurrency test
* ```/``` - json-server home page
* ```/recommendedProducts``` - 100 variant of 10 recommended products
* ```/recommendedProducts/:id```
* ```/products``` - 1000 products information
* ```/products/:uuid```

## How to run
```
./start.sh
```
This will build&run Docker container and expose server on port ```:9999```.