package handlers

import (
	"log"
	"net/http"

	"github.com/AzizAhsaan/GolangMicroServices/data"
)

type Products struct {
	l *log.Logger

}


func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.GetProducts(w,r)
		return
	}
	

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(w http.ResponseWriter, h *http.Request) {

	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil{
		http.Error(w, "Unable to marshal json",http.StatusInternalServerError)
	}
}