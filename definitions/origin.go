package definitions
import (
	"web-go-example/webgo"
)


type Origin struct {
	webgo.Middleware
}

func (h Origin) Handler (ctx *webgo.Context) bool {
	origin := ctx.Request.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	ctx.Response.Header().Set("Access-Control-Allow-Headers", "accept, content-type");
	ctx.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS");
	ctx.Response.Header().Set("Access-Control-Allow-Origin", origin);
	ctx.Response.Header().Set("Access-Control-Allow-Credentials", "true");
	ctx.Response.Header().Set("Access-Control-Expose-Headers", "Date");

	return true
}

