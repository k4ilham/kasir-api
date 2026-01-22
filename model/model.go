package model

type Produk struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	Stok       int    `json:"stok"`
	CategoryID int    `json:"category_id"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Products = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10, CategoryID: 1},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40, CategoryID: 2},
}

var Categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Kategori makanan dan snack"},
	{ID: 2, Name: "Minuman", Description: "Kategori minuman"},
}

func CategoryExists(id int) bool {
	for _, c := range Categories {
		if c.ID == id {
			return true
		}
	}
	return false
}
