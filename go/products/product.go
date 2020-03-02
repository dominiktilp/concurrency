package products

type Product struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Color string `json:"color"`
	Price string `json:"price"`
}

type Recommended struct {
	Id int64 `json:"id"`
	ProductIds []string `json:"productIds"`
}
