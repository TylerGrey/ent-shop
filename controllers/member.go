package controllers

import (
	"log"
	"net/http"

	"github.com/TylerGrey/ent-shop/forms"
	"github.com/TylerGrey/ent-shop/models"
	"github.com/gin-gonic/gin"
)

// MemberController ...
type MemberController struct {
	memberModel models.MemberModel
}

// CreateForm ...
func (mc MemberController) CreateForm(c *gin.Context) {
	c.HTML(http.StatusOK, "createMemberForm.html", nil)
}

// Create ...
func (mc MemberController) Create(c *gin.Context) {
	var joinForm forms.JoinForm

	if c.ShouldBind(&joinForm) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "invalid form"})
		return
	}

	mc.memberModel.Join(joinForm)

	c.Redirect(http.StatusMovedPermanently, "/")
}

// List ...
func (mc MemberController) List(c *gin.Context) {
	members, err := mc.memberModel.FindMembers()
	if err != nil {
		log.Println(err.Error())
	}
	c.HTML(http.StatusOK, "memberList.html", gin.H{
		"Members": members,
	})
}
