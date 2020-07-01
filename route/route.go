package route

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/v12"
	"iris-gorm-demo/controllers"
	"iris-gorm-demo/middleware"
	"net/http"
)

func InitRouter(app *iris.Application) {
	//app.Use(CrossAccess)
	bathUrl := "/api"
	mvc.New(app.Party(bathUrl + "/user")).Handle(controllers.NewUserController())
	app.Use(middleware.GetJWT().Serve) // jwt
	mvc.New(app.Party(bathUrl + "/book")).Handle(controllers.NewBookController())
}

func CrossAccess11(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
func CrossAccess(ctx iris.Context) {
	ctx.ResponseWriter().Header().Add("Access-Control-Allow-Origin", "*")
}
