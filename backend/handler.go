package main

import (
	//structure "backend/structs"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handleRequests() {
	headersOk := handlers.AllowCredentials()
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "HEAD", "POST", "PUT", "OPTIONS"})
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/login", loginPage)
	router.HandleFunc("/refresh", refreshToken)
	router.HandleFunc("/api/v1/blog/getUser", getUser).Methods("GET")
	router.HandleFunc("/api/v1/blog/adduser", addUser).Methods("POST")
	router.HandleFunc("/api/v1/blog/deleteuser/{RollNo}", deleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/blog/UpdateUser/{RollNo}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/blog/addaddress/{RollNo}", addAddress).Methods("POST")
	router.HandleFunc("/api/v1/blog/addaddress/{RollNo}", getUsersAddress).Methods("GET")
	router.HandleFunc("/api/v1/blog/deleteaddress/{RollNo}/{AddName}", deleteAddress).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4002", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
