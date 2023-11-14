package main

import (
	"go-web/config"
	homecontroller "go-web/controllers"
	categorycontroller "go-web/controllers/categories"
	productcontroller "go-web/controllers/product"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/pagination", categorycontroller.IndexPagination)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)
	http.HandleFunc("/categories/export", categorycontroller.ExportExcel)

	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
