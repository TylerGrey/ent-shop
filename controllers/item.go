package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/TylerGrey/ent-shop/forms"
	"github.com/TylerGrey/ent-shop/models"
	"github.com/gin-gonic/gin"
)

// ItemController ...
type ItemController struct {
	itemModel models.ItemModel
}

// CreateForm ...
func (ic ItemController) CreateForm(c *gin.Context) {
	c.HTML(http.StatusOK, "createItemForm.html", nil)
}

// Create ...
func (ic ItemController) Create(c *gin.Context) {
	var bookForm forms.BookForm

	if c.ShouldBind(&bookForm) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "invalid form"})
		return
	}

	_, err := ic.itemModel.SaveItem(bookForm)
	if err != nil {
		log.Println(err.Error())
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

// List ...
func (ic ItemController) List(c *gin.Context) {
	items, err := ic.itemModel.FindItems()
	if err != nil {
		log.Println(err.Error())
	}
	c.HTML(http.StatusOK, "itemList.html", gin.H{
		"Items": items,
	})
}

// UpdateItemForm ...
func (ic ItemController) UpdateItemForm(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("itemId"))
	item, _ := ic.itemModel.FindOne(id)
	c.HTML(http.StatusOK, "updateItemForm.html", gin.H{
		"Form": item,
	})
}

// UpdateItem ...
func (ic ItemController) UpdateItem(c *gin.Context) {
	var bookForm forms.BookForm

	if c.ShouldBind(&bookForm) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "invalid form"})
		return
	}

	_, err := ic.itemModel.UpdateItem(bookForm)
	if err != nil {
		log.Println(err.Error())
	}

	c.Redirect(http.StatusMovedPermanently, "/items")
}
