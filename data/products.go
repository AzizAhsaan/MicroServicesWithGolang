package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
 	Price float32 `json:"price"`
	SKU string `json:"sku"` 
	CreatedOn string `json:"-"` 
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`

}

type Products []*Product

func (p *Products) ToJSON (w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
}
var productList = []*Product {
	&Product{
		ID: 1,
		Name: "Coffe",
		Description:"Nice Coffer" ,
		Price:1.2 ,
		SKU:"sky" ,
		CreatedOn:time.Now().UTC().String() ,
		UpdatedOn:time.Now().UTC().String() ,
		DeletedOn:time.Now().UTC().String() ,
	},
	&Product{
		ID: 2,
		Name: "Tea",
		Description:"excellett Tea" ,
		Price:1.2 ,
		SKU:"sku2" ,
		CreatedOn:time.Now().UTC().String() ,
		UpdatedOn:time.Now().UTC().String() ,
		DeletedOn:time.Now().UTC().String() ,
	},
}