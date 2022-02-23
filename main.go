package main

import (
	"context"
	"log"
	"net/http"
	"paygateway/controllers"
	"paygateway/middlewares"
	"paygateway/services"

	postgresqldao "paygateway/dao/postgresql"

	"github.com/gorilla/mux"
)

func handlers(c *controllers.Controller, m *middlewares.Middleware) {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/login", c.Login).Methods("POST")
	myRouter.HandleFunc("/authorize", c.Authorize).Methods("POST")
	myRouter.HandleFunc("/capture", c.Capture).Methods("POST")
	myRouter.HandleFunc("/void", c.Void).Methods("POST")
	myRouter.HandleFunc("/refund", c.Refund).Methods("POST")

	myRouter.Use(m.IsAuthorized)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	ctx := context.Background()

	dao := postgresqldao.NewPostgreSql(ctx)
	defer dao.Disconnect()

	services := services.Services{
		Dao: dao,
	}

	controllers := &controllers.Controller{Dao: dao, Services: &services}
	middlewares := &middlewares.Middleware{Dao: dao}

	handlers(controllers, middlewares)
}
