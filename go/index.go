package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	port := os.Getenv("PORT")
	dataHost := os.Getenv("DATA_HOST")

	r := httprouter.New()
	r.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		processResponse(w, "Hello World!")
	})

	r.GET("/fib/:n", func(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		n, _ := strconv.Atoi(params.ByName("n"))
		fibn := fib(n)
		processResponse(w, fmt.Sprintf("fib(%d)=%d", n, fibn))
	})

	r.GET("/sleep/:n", func(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		n, _ := strconv.Atoi(params.ByName("n"))
		time.Sleep(time.Duration(n) * time.Millisecond)
		processResponse(w, fmt.Sprintf("sleep(%d)", n))
	})

	r.GET("/products/:n", func(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		id, _ := strconv.ParseInt(params.ByName("n"), 10, 64)
		wg := sync.WaitGroup{}

		prod := &product{}
		wg.Add(1)
		go func(prod *product) {
			resp, _ := http.Get(fmt.Sprintf("%sproducts/%d", dataHost, id))
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &prod)
			wg.Done()
		}(prod)

		reviewsResp := &dataReviewResponse{
			Reviews: make([]*review, 0),
		}
		wg.Add(1)
		go func(reviewsResp *dataReviewResponse) {
			resp, _ := http.Get(fmt.Sprintf("%sproductReviews/%d", dataHost, id))
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(body, &reviewsResp)
			wg.Done()
		}(reviewsResp)

		wg.Wait()
		prod.Reviews = reviewsResp.Reviews
		processResponse(w, prod)
	})

	r.GET("/recommendedProducts/:n", func(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		id, _ := strconv.ParseInt(params.ByName("n"), 10, 64)

		rcmd := &recommended{}
		resp, _ := http.Get(fmt.Sprintf("%srecommendedProducts/%d", dataHost, id))
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &rcmd)

		wg := sync.WaitGroup{}
		prods := make(chan *product, len(rcmd.ProductIds))
		for _, v := range rcmd.ProductIds {
			wg.Add(1)
			go func(id int64, prods chan *product) {
				defer wg.Done()
				resp, _ := http.Get(fmt.Sprintf("%sproducts/%d", dataHost, id))
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				prod := &product{}
				json.Unmarshal(body, &prod)
				prods <- prod
			}(v, prods)
		}

		done := make(chan struct{})
		products := make([]*product, 0)
		go func() {
			for {
				select {
				case prod := <-prods:
					products = append(products, prod)
				case <-done:
					break
				}
			}
		}()
		wg.Wait()
		done <- struct{}{}

		processResponse(w, products)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func processResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(body)
}

type product struct {
	Id      int64     `json:"id"`
	Uuid    string    `json:"uuid"`
	Title   string    `json:"title"`
	Image   string    `json:"image"`
	Color   string    `json:"color"`
	Price   string    `json:"price"`
	Reviews []*review `json:"reviews"`
}

type review struct {
	Id      string `json:"id"`
	User    string `json:"user"`
	Rating  int8   `json:"rating"`
	Comment string `json:"comment"`
}

type dataReviewResponse struct {
	Reviews []*review `json:"reviews"`
}

type recommended struct {
	Id         int64   `json:"id"`
	ProductIds []int64 `json:"productIds"`
}
