package aipaper

import (
	"github.com/zachturing/util/config"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func GetProductList() ([]Product, error) {
	cfg, err := config.Get(config.Common)
	if err != nil {
		return nil, err
	}
	var productList []Product
	err = cfg.GetWithUnmarshal("product", &productList, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return productList, nil
}
