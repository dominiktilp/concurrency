const express = require('express')
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

app.listen(port, () => console.log(`Example app listening on port ${port}!`))