package controllers

import (
	"github.com/revel/revel"
	_ "database/sql"
	_"log"
	_"github.com/revel/modules/db/app"
	_ "github.com/go-sql-driver/mysql"
	_"github.com/mattn/go-oci8"
	"fmt"
	"myapp/app/database"
)

type App struct {
	*revel.Controller
}



func (c App) Index(age int) revel.Result {
	//greeting := "명바기 짱!!!"
	fmt.Printf("age is %d", age )
	return c.RenderJson(age)
}


func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("name is requered!")
	c.Validation.MinSize(myName,3).Message("name is not long enough!")
	c.Validation.MaxSize(myName,10).Message("name is too long!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)

	}

	var resultMap map[string]string

	err, resultMap:= database.TestSelect();

	if err != nil {
		errmsg := "Error connecting to the database"
		fmt.Println(err)
		return c.Redirect(errmsg)

	}

	return c.Render(resultMap)
}



