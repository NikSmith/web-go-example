package controllers
import (
	"webapi/webgo"
	"webapi/models"
	"strconv"
)

type Users struct {
	webgo.Controller
}


func (c *Users) GetUser() {
	var err error
	var id int

	id, err = strconv.Atoi(c.Ctx.Params["id"])
	if err !=nil {
		c.Error(404,"")
		return
	}

	user:=models.User{}
	user.GetUser(id)

	if user.UserId == 0 {
		c.Error(404,"")
		return
	}

	c.Json(user)
}