package controllers
import (
	"webapi/webgo"
)

type Test struct{
	webgo.Controller
}

func (c *Test) Index() {
	c.SetHeader("header-test","test");
	c.Plain("Hello World")
}

func (c *Test) TestRender () {
	type Person struct {
		UserName string
	}
	p := Person{UserName: "Astaxie"}
	c.Render("welcome",p)
}

func (c *Test) TestJson () {
	f:= []string{"id","1"}
	c.Json(f)
}