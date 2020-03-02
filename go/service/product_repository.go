package service

import (
	"concurrency/products"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ProductRepository struct {
	path *url.URL
}

func NewProductRepository(path string) (*ProductRepository, error) {
	parsedPath, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("parsing URL failed: %w", err)
	}
	return &ProductRepository{
		path: parsedPath,
	}, nil
}

func (this ProductRepository) GetAll() (prods []*products.Product, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/products", this.path.String()))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &prods)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling failed: %w", err)
	}
	return
}

func (this ProductRepository) Get(id string) (prod *products.Product, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/products/%s", this.path.String(), id))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &prod)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling failed: %w", err)
	}
	return
}

func (this ProductRepository) GetRecommendedAll() (recs []*products.Recommended, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/recommendedProducts", this.path.String()))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &recs)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling failed: %w", err)
	}
	return
}

func (this ProductRepository) GetRecommended(id int64) (rec *products.Recommended, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/recommendedProducts/%d", this.path.String(), id))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &rec)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling failed: %w", err)
	}
	return
}
