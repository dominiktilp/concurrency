package products

type Repository interface {
	GetAll() ([]*Product, error)
	Get(id string) (*Product, error)
	GetRecommendedAll() ([]*Recommended, error)
	GetRecommended(id int64) (*Recommended, error)
}

type ProductHandler struct {
	repository Repository
}

func NewProductHandler(repo Repository) *ProductHandler {
	return &ProductHandler {
		repository: repo,
	}
}

func (this *ProductHandler) GetAll() ([]*Product, error) {
	return this.repository.GetAll()
}

func (this *ProductHandler) Get(id string) (*Product, error) {
	return this.repository.Get(id)
}

func (this *ProductHandler) GetRecommendedAll() ([]*Recommended, error) {
	return this.repository.GetRecommendedAll()
}

func (this *ProductHandler) GetRecommended(id int64) (*Recommended, error) {
	return this.repository.GetRecommended(id)
}
