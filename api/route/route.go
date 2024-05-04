package route

import (
	"net/http"
	"sof/api/controller"
	"time"
    "log"

	"github.com/julienschmidt/httprouter"
)

type middleware struct {
    r *httprouter.Router
}

// 实现ServeHTTP
func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    m.r.ServeHTTP(w, r)
}

// 需要实现ServeHTTP
func NewMiddlewareHandler(r *httprouter.Router) http.Handler {
    return &middleware{
        r: r,
    }
}

func (m *middleware) LoggingMiddlewareHandler(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        end := time.Now()
        log.Printf("[%s] - %q - %v\n", r.Method, r.URL.String(), end.Sub(start))
    }
    return http.HandlerFunc(fn)
}

func (m *middleware) RecoverHandler(next http.Handler) http.Handler {
    fn := func (w http.ResponseWriter, r *http.Request)  {
        defer func() {
            if err := recover(); err != nil {
                http.Error(w, http.StatusText(500), 500)
            }
        }()
        next.ServeHTTP(w, r)
    }

    return http.HandlerFunc(fn)
}

func InitRoutes() *httprouter.Router {
    r := httprouter.New()


    r.POST("/v1/user/created", controller.CreatedUserController)
    return r
}

