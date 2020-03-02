const express = require('express')
const fetch = require('node-fetch')
const app = express()
const port = process.env.PORT


function fib(n) {
  if (n <= 1) {
    return n
  }
  return fib(n-1) + fib(n-2)
}

app.get(
  '/',
  (req, res) => {
    res.send('Hello World!')
  }
)

app.get(
  '/fib/:n',
  (req, res) => {
    const n = req.params.n
    const fibn = fib(parseInt(n, 10) || 40)
    res.send(`fib(${n}) = ${fibn}`)
  }
)

app.get(
  '/sleep/:n',
  async (req, res) => {
    const n = req.params.n
    await new Promise(r => setTimeout(r, n));
    res.send(`sleep(${n})`)
  }
)

app.get(
  '/products/:id',
  async (req, res) => {
    try {
      const id = req.params.id
      const DATA_HOST = process.env.DATA_HOST || 'http://localhost:9999/';
      const productPromise = fetch(`${DATA_HOST}products/${id}`).then(resp => resp.json());
      const productReviewsPromise = fetch(`${DATA_HOST}productReviews/${id}`).then(resp => resp.json());
      const [productData, productReviewsData] = await Promise.all([productPromise, productReviewsPromise])
      const data = {
        ...productData,
        reviews: productReviewsData
      }

      res.json(data)
    } catch {
      res.status(500)
      res.send('')
    }

  }
)

app.get(
  '/recommendedProducts/:id',
  async (req, res) => {
    try {
      const id = req.params.id
      const DATA_HOST = process.env.DATA_HOST || 'http://localhost:9999/';
      const recommendedProductsData = await fetch(`${DATA_HOST}recommendedProducts/${id}`).then(resp => resp.json());
      const productsPromises = recommendedProductsData.productIds.map(
        productId => fetch(`${DATA_HOST}products/${productId}`).then(resp => resp.json())
      )
      const productsData = await Promise.all(productsPromises)
      const data = {
        ...recommendedProductsData,
        products: productsData
      }

      res.json(data)
    } catch {
      res.status(500)
      res.send('')
    }
  }
)

app.listen(port, () => console.log(`Example app listening on port ${port}!`))