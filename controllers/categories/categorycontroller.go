package categorycontroller

import (
	categoryentities "go-web-native/entities"
	"go-web-native/models/categorymodel"
	"net/http"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()

	// Jika akan melempar data ke sebuah file/views, memerlukan variabel yg tipe datanya adalah map dengan keynya string dan
	// valuenya adalah any atau interface kosong yg berasal dari categories atas itu
	data := map[string]any{
		"categories": categories, //"categories" dalam petik adalah key yg dipanggil di views
	}

	//variabel begini itu return 2, data yg disimpan ke variabel yg kita declare, atau error, jadi declarenya bentuk begini
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data) //mengeksekusi pengiriman ke tamplate, yg dikirim response dan data, data return query dari DB
}

func Add(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "GET") {
		temp, err := template.ParseFiles("views/category/create.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if(r.Method == "POST") {
		var category categoryentities.CategoryEntities  //var category tipe struct dari entities

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok { //categorymodel.Create(category) -> mengirimkan var category ke func Create di model
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
