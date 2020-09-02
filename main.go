package main

import (
	"github.com/TylerGrey/ent-shop/controllers"
	"github.com/TylerGrey/ent-shop/db"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	db.Init()

	hc := new(controllers.HomeController)
	r.GET("/", hc.Home)

	item := r.Group("/items")
	{
		ic := new(controllers.ItemController)
		item.GET("/new", ic.CreateForm)
		item.POST("/new", ic.Create)
		item.GET("/", ic.List)
		item.GET("/edit/:itemId", ic.UpdateItemForm)
		item.POST("/edit/:itemId", ic.UpdateItem)
	}

	member := r.Group("/members")
	{
		mc := new(controllers.MemberController)
		member.GET("/new", mc.CreateForm)
		member.POST("/new", mc.Create)
		member.GET("/", mc.List)
	}

	order := r.Group("/order")
	{
		oc := new(controllers.OrderController)
		order.GET("/", oc.CreateForm)
		order.POST("/", oc.Order)
	}

	orders := r.Group("/orders")
	{
		oc := new(controllers.OrderController)
		orders.GET("/", oc.OrderList)
		orders.POST("/", oc.CancelOrder)
	}

	r.LoadHTMLGlob("./public/templates/**/*")
	r.Static("/css", "./public/static/css")
	r.Static("/js", "./public/static/js")

	r.Run()
}
