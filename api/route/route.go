package route

import (
	"net/http"
	"sof/api/controller"

	"github.com/julienschmidt/httprouter"
)

type middlewareHandler struct {
    r *httprouter.Router
}

// 实现ServeHTTP
func (m *middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    m.r.ServeHTTP(w, r)
}

// 需要实现ServeHTTP
func NewMiddlewareHandler(r *httprouter.Router) http.Handler {
    return &middlewareHandler{
        r: r,
    }
}

func InitRoutes() *httprouter.Router {
    r := httprouter.New()

    r.POST("/v1/user/created", controller.CreatedUserController)
    return r
}

