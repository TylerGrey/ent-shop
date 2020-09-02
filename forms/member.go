package forms

// JoinForm ...
type JoinForm struct {
	Name    string `form:"name" json:"name" binding:"required"`
	City    string `form:"city" json:"city"`
	Street  string `form:"street" json:"street"`
	Zipcode string `form:"zipcode" json:"zipcode"`
}
