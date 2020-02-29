<?php
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;

require __DIR__ . '/vendor/autoload.php';

$app = AppFactory::create();

function fib(int $n) {
    if ($n <= 1) {
        return $n;
    }
    return fib($n-1) + fib($n-2);
}

$app->get('/', function (Request $request, Response $response, $args) {
    $response->getBody()->write("Hello world!");
    return $response;
});

$app->get('/fib/{n}', function (Request $request, Response $response, $args) {
    $n = $args["n"];
    $fibn = fib($n);
    $response->getBody()->write("fib(".$n.") = ".$fibn);
    return $response;
});

$app->get('/sleep/{n}', function (Request $request, Response $response, $args) {
    $n = $args["n"];
    usleep($n * 1000);
    $response->getBody()->write("sleep(".$n.")");
    return $response;
});

$app->run();