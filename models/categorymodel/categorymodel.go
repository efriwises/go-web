package categorymodel

import (
	"go-web-native/config"
	categoryentities "go-web-native/entities"
)

// func GetAll() () { awalnya fungsi biasa, tapi kalau ada return jadi dibawahnya
func GetAll() []categoryentities.CategoryEntities { //[]categoryentities.CategoryEntities disamakan sama return di akhir fungsinya
	rows, err := config.DB.Query(`Select * from categories`) //query mengembalikan 2 param, err dan rows

	if err != nil {
		panic(err)
	}

	defer rows.Close() //tutup var rows

	//variabel untuk menampung data kategorinya, [] karena akan menampung banyak data
	var categories []categoryentities.CategoryEntities //packgae.func

	//looping rows
	for rows.Next() {
		var category categoryentities.CategoryEntities // query bawah akan disimpan ke var ini

		// err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)

		// if err != nil {
		// 	panic(err)
		// }

		//kode diatas diringkas jadi dibawah ini
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category) // dikirim ke categories yg atas itu
	}

	return categories
}

func Create(category categoryentities.CategoryEntities) bool {
	result, err := config.DB.Exec(`
	INSERT INTO categories (name, created_at, updated_at)
	VALUE (?,?,?)`,
	category.Name, category.CreatedAt, category.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}
