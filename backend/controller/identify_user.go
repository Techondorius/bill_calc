package controller

import (
	"log"

	"github.com/Techondorius/bill_calc/model"
	"github.com/Techondorius/bill_calc/view"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	UserID   string `json:"userid"`
	Password string `json:"password"`
}

type ILogin interface {
	CompairHashPWandRow(string) error
}

func Login(c *gin.Context) {
	var lr = loginRequest{}
	c.ShouldBindJSON(&lr)
	log.Println(lr)
	u, err := model.FindUserByID(lr.UserID)
	if err != nil {
		view.Unauthorized(c, "invalid id or password")
		return
	}

	err = u.CompairHashPWandRow(lr.Password)
	if err != nil {
		view.Unauthorized(c, "invalid id or password")
		return
	}

	token := GenerateJWT(u.UserID.ReturnString())
	view.StatusOK(c, gin.H{
		"token": token,
	})
}
