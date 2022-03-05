package main

import (
	//structure "backend/structs"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	router.HandleFunc("/api/v1/blog/UpdateAddress/{RollNo}/{AddName}", UpdateAddress).Methods("PUT")
	router.HandleFunc("/api/v1/blog/UpdateUser/{RollNo}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/blog/addaddress/{RollNo}", addAddress).Methods("POST")
	router.HandleFunc("/api/v1/blog/addaddress/{RollNo}", getUsersAddress).Methods("GET")
	router.HandleFunc("/api/v1/blog/deleteaddress/{RollNo}/{AddName}", deleteAddress).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4002", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside Update Address")
	if isValidCoockie(w, r) {
		db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		var address []Address
		var addressofUser []Address
		var addresstoupdate Address
		var addressofUsertoupdate Address
		var user User
		_ = json.NewDecoder(r.Body).Decode(&addressofUsertoupdate)
		params := mux.Vars(r)
		Rollno := params[("RollNo")]
		Addname := params[("AddName")]
		db.Where("roll_no = ?", Rollno).Find(&user)

		db.Where("uid = ?", user.UID).Find(&addressofUser)
		for _, a := range addressofUser {
			if a.AddressName == Addname {
				addresstoupdate = a
				break
			}
		}
		fmt.Println("addresstoupdate:", addresstoupdate.AddID)
		addresstoupdate.City = addressofUsertoupdate.City
		addresstoupdate.Pincode = addressofUsertoupdate.Pincode
		addresstoupdate.FirstLineAddress = addressofUsertoupdate.FirstLineAddress
		db.Model(&address).Where("add_id = ?", addresstoupdate.AddID).Updates(addresstoupdate)
		json.NewEncoder(w).Encode(addresstoupdate)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
