package controller

import (
	"fmt"
	"time"

	"github.com/Techondorius/bill_calc/model"
	"github.com/Techondorius/bill_calc/view"
	"github.com/gin-gonic/gin"
)

func NewBillHandlerFunc(c *gin.Context) {
	nbr := NewBillRequest{}
	c.ShouldBindJSON(&nbr)
	fmt.Println(nbr)
	if len(nbr.UIDPrice) == 0 {
		view.BadRequest(c, "uIDPrice is required")
		return
	}
	t, err := time.Parse(time.RFC3339, nbr.Date)
	if err != nil {
		view.BadRequest(c, err.Error())
		return
	}

	b, err := model.NewBill(nbr.Subject, t, nbr.UIDPrice, nbr.Note)
	if err != nil {
		view.BadRequest(c, err.Error())
		return
	}

	for _, uidp := range nbr.UIDPrice {
		b.AddUserToBillingUsers(uidp.UserID, uidp.Price)
	}

	fmt.Println(b)
	err = b.SaveToDB()
	if err != nil {
		view.BadRequest(c, err.Error())
		return
	}
}

type NewBillRequest struct {
	Subject  string           `json:"subject"`
	Date     string           `json:"date"`
	UIDPrice []model.UIDPrice `json:"uIDPrice"`
	Note     string           `json:"note"`
}
