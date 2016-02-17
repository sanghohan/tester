package controllers

import (
	"github.com/revel/revel"
	"fmt"
)

type TestApi struct {
	*revel.Controller
}


func (c TestApi) Index(age int) revel.Result {

	fmt.Printf("age is %d", age )

	return c.RenderJson(age)
}
