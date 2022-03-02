package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type User struct {
	UID     string `json:"UID"`
	FName   string `json:"FName"`
	RollNo  string `json:"RollNo"`
	Contact string `json:"Contact"`
}

var user []User
var flagfordelete int

func main() {
	flagfordelete = 0
	genrateUID := uuid.New()

	user = append(user, User{
		UID:     genrateUID.String(),
		FName:   "cheese",
		RollNo:  "54",
		Contact: "989898998",
	})
	user = append(user, User{
		UID:     genrateUID.String(),
		FName:   "yashshah",
		RollNo:  "55",
		Contact: "12348",
	})

	handleRequests()
}

func handleRequests() {
	headersOk := handlers.AllowCredentials()
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "HEAD", "POST", "PUT", "OPTIONS"})
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")

	router.HandleFunc("/api/v1/blog/getUser", getUser).Methods("GET")
	router.HandleFunc("/api/v1/blog/adduser", addUser).Methods("POST")
	router.HandleFunc("/api/v1/blog/deleteuser/{RollNo}", deleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/blog/UpdateUser/{RollNo}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":4002", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpont Called homePage:")
}
func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	genrateUID := uuid.New()
	newUser.UID = genrateUID.String()
	user = append(user, newUser)
	fmt.Println("Inside addUser")
	fmt.Println(newUser)
	json.NewEncoder(w).Encode(user)

}
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	fmt.Println("Inside getUser")
	json.NewEncoder(w).Encode(user)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	params := mux.Vars(r)
	fmt.Println(params)
	if getByRoll(params[("RollNo")]) {
		_deleteUserAtUid(params[("RollNo")])
		fmt.Println("Inside DeleteUSer")
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
func _deleteUserAtUid(RollNo string) {
	for index, u := range user {
		if u.RollNo == RollNo {
			user = append(user[:index], user[index+1:]...)
			flagfordelete = 1
			break
		}
	}
	if flagfordelete == 1 {
		flagfordelete = 0
	} else {
		fmt.Println("Invalid Roll No")
		flagfordelete = 2
	}
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	params := mux.Vars(r)
	fmt.Println(params)
	_deleteUserAtUid(params[("RollNo")])
	if flagfordelete == 0 {
		genrateUID := uuid.New()
		newUser.UID = genrateUID.String()
		user = append(user, newUser)
	} else if flagfordelete == 2 {
		flagfordelete = 0
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Println("Inside UpdateUser")
	json.NewEncoder(w).Encode(user)
}
func getByRoll(RollNo string) bool {
	for _, u := range user {
		if RollNo == u.RollNo {
			return true
		}
	}
	return false
}
