package main
import (
	"web-go-example/webgo"
	"web-go-example/controllers"
	"web-go-example/definitions"
)


func main()  {
	webgo.RegisterMiddleware("origin", definitions.Origin{})

	webgo.Get("/",&controllers.Test{},"",nil,"Index")
	webgo.Get("/render",&controllers.Test{},"",nil,"TestRender")
	webgo.Get("/json",&controllers.Test{},"",nil,"TestJson")
	webgo.Post("/post",&controllers.Test{},"",nil,"CheckJsonPost")
	webgo.Post("/simple",&controllers.Test{},"",nil,"CheckPost")
	webgo.Run(":8082")
}
