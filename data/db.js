var faker = require('faker');

module.exports = () => {
  const data = {
    products: [],
    recommendedProducts: [],
    productReviews: []
  };
  const numberOfProducts = 1000;
  for (let i = 0; i < numberOfProducts; i++) {
    data.products.push(
      {
        id: i,
        uuid: faker.random.uuid(),
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
  data.products.map(product => {
    const reviews = [];
    for (let i = 0; i < faker.random.number(15); i++) {
      reviews.push(
        {
          id: faker.random.uuid(),
          user: faker.internet.email(),
          rating: faker.random.number(5),
          comment: faker.lorem.sentence()
        }
      )
    }
    data.productReviews.push(
      {
        id: product.id,
        uuid: product.uuid,
        reviews: reviews
      }
    )
  })
    
  return data
}