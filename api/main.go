package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", LoginUser)
	return router
}


func main() {
	fmt.Println("video start sever")
	router := RegisterHandlers()
	http.ListenAndServe(":8000", router)
	fmt.Println("video end sever")
}