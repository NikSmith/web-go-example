package main
import (
	"webapi/webgo"
	"webapi/controllers"
)


func main()  {
	webgo.Get("/",&controllers.Test{},nil,nil,"Index")
	webgo.Get("/render",&controllers.Test{},nil,nil,"TestRender")
	webgo.Get("/json",&controllers.Test{},nil,nil,"TestJson")
	webgo.Run(":8081")
}
