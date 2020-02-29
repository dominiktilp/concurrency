var faker = require('faker');

module.exports = () => {
  const data = {
    products: [],
    recommendedProducts: [],
  };
  const numberOfProducts = 1000;
  for (let i = 0; i < numberOfProducts; i++) {
    data.products.push(
      {
        id: faker.random.uuid(),
        title: faker.commerce.productName(),
        image: faker.image.animals(),
        color: faker.commerce.color(),
        price: faker.commerce.price()
      }
    )
  }
  for (let i = 0; i < 10; i++) {
    data.recommendedProducts.push(
      {
        id: i,
        productIds: Array(10).fill().map(_ => data.products[Math.ceil(Math.random() * (numberOfProducts-1))].id)
      }
    )
  }
  return data
}