package app

import (
	"task/controller"
	"task/exception"

	"github.com/julienschmidt/httprouter"
	group "github.com/mythrnr/httprouter-group"
)

func NewRouter(auth controller.AuthUserHandler) *httprouter.Router {
	router := httprouter.New()

	group.New("/api").Middleware()

	router.POST("/api/auth/register", auth.Create)
	router.POST("/api/auth/login", auth.Login)
	router.GET("/api/dans/career", auth.FindCareer)

	router.GET("/api/dans/career/:id", auth.FindCareerById)

	router.PanicHandler = exception.ErrorHandler
	return router
}
