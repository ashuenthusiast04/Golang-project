package routes

import (
	"github.com/gorilla/mux"
)


var registerbookstoreroutes=func(router *mux.Router){
	router.HandleFunc("/books",controllers.createbooks).Methods("POST")
	router.HandleFunc("/books",controllers.getbooks).Methods("GET")
		router.HandleFunc("/books/{bookid}",controllers.getbook).Methods("GET")
	router.HandleFunc("/books/{bookid}",controllers.updatebook).Methods("PUT")
	router.HandleFunc("/books/{bookid}",controllers.deletebook).Methods("DELETE")
	

}