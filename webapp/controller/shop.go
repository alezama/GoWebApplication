package controller

import (
	"FirstWebApplication/webapp/model"
	"FirstWebApplication/webapp/viewmodel"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
	productTemplate  *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	http.HandleFunc("/shop/", s.handleShop)
	http.HandleFunc("/products/", s.handleProducts)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		s.handleCategory(w, r, categoryID)
	} else {
		categories := model.GetCategories()
		vm := viewmodel.NewShop(categories)
		s.shopTemplate.Execute(w, vm)
	}
}

func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, id int) {
	products := model.GetProductsForCategory(id)
	vm := viewmodel.NewShopDetail(products)
	s.categoryTemplate.Execute(w, vm)
}

func (s shop) handleProducts(w http.ResponseWriter, r *http.Request) {
	productPattern, _ := regexp.Compile(`/products/(\d+)`)
	matches := productPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		productID, _ := strconv.Atoi(matches[1])
		product, err := model.GetProduct(productID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		vm := viewmodel.NewProduct(product)
		s.productTemplate.Execute(w, vm)

	} else {
		http.Error(w, "Mal formed URL", http.StatusBadRequest)
	}
}
