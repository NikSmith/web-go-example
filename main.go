package main
import (
	"webapi/webgo"
	"webapi/controllers"
	"webapi/definitions"
	"webapi/modules"
)


func main()  {
	webgo.RegisterMiddleware("origin", definitions.Origin{})
	webgo.RegisterModule("db", &modules.Mongo{})


	webgo.Get("/",&controllers.Test{},"","Index")
	webgo.Get("/render",&controllers.Test{},"","TestRender")
	webgo.Get("/json",&controllers.Test{},"","TestJson")
	webgo.Post("/post",&controllers.Test{},"","CheckJsonPost")
	webgo.Post("/simple",&controllers.Test{},"","CheckPost")
	webgo.Get("/users/:id",&controllers.Users{},"","GetUser")
	webgo.Run()
}
