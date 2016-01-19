package controllers
import (
	"webapi/webgo"
	"fmt"
)

type Test struct{
	webgo.Controller
}


func (c *Test) Index() {
	c.Plain("Hello World")
}

func (c *Test) TestRender () {
	c.Ctx.SetCookie("efef","eee",100, "/","127.0.0.1",false,false)

	//cookieValue:=c.Ctx.GetCookie("efef")

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
	var err error
	var t map[string]int
	err = c.Ctx.ValidateSchema(&t)
	if err != nil {
		fmt.Println(err)
	}
	c.Json(c.Ctx.Body)
}

func (c *Test) CheckJsonPost () {
	var err error
	var t map[string][]int
	err = c.Ctx.ValidateSchema(&t)
	if err != nil {

	}
	c.Json(t)
}