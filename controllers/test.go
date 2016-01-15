package controllers
import (
	"web-go-example/webgo"
//	"fmt"
	"fmt"
)

type Test struct{
	webgo.Controller
}

func (c *Test) Index() {
	c.Plain("Hello World")
}

func (c *Test) TestRender () {
	c.SetHeader("header-test","test");
	type Person struct {
		UserName string
	}
	p := Person{UserName: "Astaxie"}
	c.Render("index",p)
}

func (c *Test) TestJson () {
	f:= []string{"id","1"}
	c.Json(f)
}

func (c *Test) CheckPost () {
	var t map[string]int
	err := c.Ctx.ValidateSchema(&t)
	if err != nil {
		fmt.Println(err)
	}
	c.Json(c.Ctx.Body)
}

func (c *Test) CheckJsonPost () {
	var t map[string][]int
	err := c.Ctx.ValidateSchema(&t)
	if err != nil {

	}
	c.Json(t)
}