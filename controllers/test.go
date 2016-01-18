package controllers
import (
	"webapi/webgo"
//	"fmt"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	DB:=webgo.MODULES("db").GetInstance().(*mgo.Session)

	db := DB.Copy()
	fmt.Println(db.DB("test").C("test").Find(bson.M{"efef":333}))

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