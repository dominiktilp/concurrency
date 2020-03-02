<?php
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;
use \GuzzleHttp\Client;


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

$app->get('/products/{id}', function (Request $request, Response $response, $args) {
    $DATA_HOST = getenv('DATA_HOST');
    if (empty($DATA_HOST)) {
        $DATA_HOST = 'http://localhost:9999/';
    }
    $id = $args["id"];
    $client = new \GuzzleHttp\Client();
    
    $resProduct = $client->request('GET', $DATA_HOST.'products/'.$id);
    $resProductReviews = $client->request('GET', $DATA_HOST.'productReviews/'.$id);

    $data = json_decode($resProduct->getBody());
    $data->reviews = json_decode($resProductReviews->getBody());

    $response->getBody()->write(json_encode($data));
    return $response;
});

$app->get('/recommendedProducts/{id}', function (Request $request, Response $response, $args) {
    $DATA_HOST = getenv('DATA_HOST');
    if (empty($DATA_HOST)) {
        $DATA_HOST = 'http://localhost:9999/';
    }
    $id = $args["id"];
    $client = new \GuzzleHttp\Client();
    
    $resRecomendedProduct = $client->request('GET', $DATA_HOST.'recommendedProducts/'.$id);
    $data = json_decode($resRecomendedProduct->getBody());
    $data->products = [];
    foreach ($data->productIds as $productId) {
        $resProduct = $client->request('GET', $DATA_HOST.'products/'.$productId);
        $productData = json_decode($resProduct->getBody());
        array_push($data->products, $productData);
    }

    $response->getBody()->write(json_encode($data));
    return $response;
});

$app->run();