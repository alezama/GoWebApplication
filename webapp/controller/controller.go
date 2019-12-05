package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	shopController shop
)

func Startup(template map[string]*template.Template) {
	homeController.hometemplate = template["home.html"]
	homeController.registerRoutes()
	shopController.shopTemplate = template["shop.html"]
	shopController.categoryTemplate = template["shop_details.html"]
	shopController.productTemplate = template["shop_detail.html"]
	shopController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
