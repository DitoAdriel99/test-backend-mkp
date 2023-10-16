package router

import (
	"go-learn/controller/auth"
	"go-learn/controller/product"
	"go-learn/middleware"
	"go-learn/repositories"
	"go-learn/service"

	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()
	//set dependency
	repo := repositories.NewRepo()
	serv := service.NewService(repo)

	// call controllers auth
	controllerLogin := auth.NewControllerLogin(*serv)
	controllerRegister := auth.NewControllerRegister(*serv)

	// middlewares
	tokenValidator := middleware.NewTokenValidator(*repo)

	//login
	router.HandleFunc("/health", controllerLogin.Health).Methods("GET")
	router.HandleFunc("/login", controllerLogin.HandleLogin).Methods("POST")
	router.HandleFunc("/register", controllerRegister.HandleRegister).Methods("POST")

	// Controller Product
	controllerProduct := product.NewControllerProductCreate(*serv)

	root := router.PathPrefix("").Subrouter()
	root.Use(tokenValidator.ValidateTokenMiddleware())
	root.HandleFunc("/product", controllerProduct.Create).Methods("POST")
	root.HandleFunc("/product", controllerProduct.Get).Methods("GET")

	root.HandleFunc("/sales", controllerProduct.CreateSales).Methods("POST")
	root.HandleFunc("/sales", controllerProduct.GetSales).Methods("GET")

	return router
}
