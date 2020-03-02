package main

import (
	"concurrency/products"
	"concurrency/service"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Run: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag("port", cmd.Flags().Lookup("port"))
			if len(viper.GetString("port")) == 0 {
				log.Fatalln("HTTP port has not been set")
			}
			viper.BindPFlag("repo", cmd.Flags().Lookup("repo"))
			if len(viper.GetString("repo")) == 0 {
				log.Fatalln("URL to the repository has not been set")
			}

			handler, err := newProductHandler(viper.GetString("repo"))
			if err != nil {
				log.Fatal(err)
			}

			router := httprouter.New()
			router.GET("/products", handler.GetProducts)
			router.GET("/products/:id", handler.GetProduct)
			router.GET("/recommendedProducts", handler.GetRecommendedProducts)
			router.GET("/recommendedProducts/:id", handler.GetRecommended)

			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("port")), router))
		},
	}
	cmd.Flags().StringP("port", "p", "", "HTTP port")
	cmd.Flags().StringP("repo", "r", "", "URL to the repository")

	return cmd
}

type productHandler struct {
	handler *products.ProductHandler
}

func newProductHandler(repoPath string) (*productHandler, error) {
	repo, err := service.NewProductRepository(repoPath)
	if err != nil {
		return nil, fmt.Errorf("could not init the repository: %w", err)
	}
	handler := products.NewProductHandler(repo)
	return &productHandler {
		handler: handler,
	}, nil
}

func (this *productHandler) GetProducts(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	result, err := this.handler.GetAll()
	this.processResponse(w, result, err)
}

func (this *productHandler) GetProduct(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	result, err := this.handler.Get(params.ByName("id"))
	this.processResponse(w, result, err)
}

func (this *productHandler) GetRecommendedProducts(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	result, err := this.handler.GetRecommendedAll()
	this.processResponse(w, result, err)
}

func (this *productHandler) GetRecommended(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		this.processResponse(w, nil, err)
		return
	}
	result, err := this.handler.GetRecommended(id)
	this.processResponse(w, result, err)
}

func (this *productHandler) processResponse(w http.ResponseWriter, result interface{}, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}
