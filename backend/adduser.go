package main

import (
	//structure "backend/structs"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	flagfordelete = 0
	db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&adminModel{})
	var aUs []adminModel
	// db.Create(&adminModel{
	// 	Username: "yashshah",
	// 	Password: "hello",
	// })
	// db.Create(&adminModel{
	// 	Username: "kanan",
	// 	Password: "hello",
	// })

	db.Find(&aUs)

	for _, a := range aUs {
		fmt.Println("Username: ", a.Username, "Password: ", a.Password)
		adminUsers[a.Username] = a.Password
	}
	for key, b := range adminUsers {
		fmt.Println("The Map")
		fmt.Println("Username: ", key, "Password: ", b)
	}
	handleRequests()
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	refreshToken(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside DeleteUSer")
	if isValidCoockie(w, r) {
		db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		var usersForDisplay []User
		params := mux.Vars(r)
		fmt.Println(params)
		//if getByRoll(params[("RollNo")]) {
		rollToDelete := params[("RollNo")]
		fmt.Println(rollToDelete)
		db.Find(&usersForDisplay)
		for _, u := range usersForDisplay {
			fmt.Println(u.RollNo)
		}
		db.Where("roll_no = ?", rollToDelete).Delete(&usersForDisplay)
		for index, _ := range usersForDisplay {
			usersForDisplay[index] = User{}
		}
		db.Find(&usersForDisplay)
		// _deleteUserAtUid(params[("RollNo")])
		// json.NewEncoder(w).Encode(user)
		//} else {
		//w.WriteHeader(http.StatusNotFound)
		//}
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
func deleteAddress(w http.ResponseWriter, r *http.Request) {
	//refreshToken(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside deleteAddress")
	if isValidCoockie(w, r) {
		dbaddress, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect address database")
		}
		var allusers []User
		var addressfordelete []Address
		var addresstodelete Address
		var userid = ""

		dbaddress.Find(&allusers)

		params := mux.Vars(r)
		fmt.Println(params[("RollNo")])
		fmt.Println(params[("AddName")])
		for _, u := range allusers {

			if u.RollNo == params[("RollNo")] {
				userid = u.UID
				break
			}
		}
		fmt.Println(userid)
		if userid != "" {
			dbaddress.Where("uid = ?", userid).Find(&addressfordelete)
			// dbaddress.Find(&usersaddress)
			for _, a := range addressfordelete {

				if a.AddressName == params[("AddName")] {
					fmt.Println(a.AddressName)
					addresstodelete = a
					break
				}

			}
			fmt.Println(addresstodelete.City)
			dbaddress.Where("add_id = ?", addresstodelete.AddID).Delete(&addressfordelete)
			//dbaddress.Delete(&addresstodelete)
			w.WriteHeader(200)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
	}
}
