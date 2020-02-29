# PHP version
This directory contain PHP version of testing API service with these endpoints:
* ```/``` - hello world 
* ```/fib/:n``` - calculate n-th term of fibonacci sequence
* ```/sleep/:n``` - sleep for n microseconds before response

## How to run
```
./start.sh
```
This will build&run Docker container and expose server on port ```:9001```.