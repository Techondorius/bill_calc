package main

import "github.com/Techondorius/bill_calc/model"

func main() {

	err := model.InitGorm()
	if err != nil {
		panic(err)
	}

	g := Routers()
	g.Run()
}
