package forms

// BookForm ...
type BookForm struct {
	ID            int    `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	Price         int32  `form:"price" json:"price"`
	StockQuantity int32  `form:"stockQuantity" json:"stockQuantity"`
	Author        string `form:"author" json:"author"`
	Isbn          string `form:"isbn" json:"isbn"`
}
