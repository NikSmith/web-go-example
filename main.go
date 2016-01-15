package main
import (
	"webapi/webgo"
	"webapi/controllers"
	"webapi/definitions"
)


func main()  {
	webgo.RegisterMiddleware("origin", definitions.Origin{})

	webgo.Get("/",&controllers.Test{},"origin",nil,"Index")
	webgo.Get("/render",&controllers.Test{},"",nil,"TestRender")
	webgo.Get("/json",&controllers.Test{},"",nil,"TestJson")
	webgo.Run(":8081")
}
