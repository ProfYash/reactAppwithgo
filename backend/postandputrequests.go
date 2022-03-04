package main

import (
	//structure "backend/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	fmt.Println("Inside Login")
	fmt.Println(credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPassword, ok := adminUsers[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 20)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

func addUser(w http.ResponseWriter, r *http.Request) {
	refreshToken(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside addUser")
	if isValidCoockie(w, r) {
		var newUser User
		_ = json.NewDecoder(r.Body).Decode(&newUser)
		genrateUID := uuid.New()
		newUser.UID = genrateUID.String()
		//user = append(user, newUser)
		db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db.AutoMigrate(&User{})
		db.Create(&newUser)

		fmt.Println(newUser)
		json.NewEncoder(w).Encode(newUser)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	refreshToken(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside UpdateUser")
	db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var usersForDisplay []User

	if isValidCoockie(w, r) {
		var newUser User
		_ = json.NewDecoder(r.Body).Decode(&newUser)
		params := mux.Vars(r)
		fmt.Println(params)
		db.Model(&usersForDisplay).Where("roll_no = ?", params[("RollNo")]).Updates(newUser)
		// _deleteUserAtUid(params[("RollNo")])
		// if flagfordelete == 0 {
		// 	genrateUID := uuid.New()
		// 	newUser.UID = genrateUID.String()
		// 	user = append(user, newUser)
		// } else if flagfordelete == 2 {
		// 	flagfordelete = 0
		// 	w.WriteHeader(http.StatusNotFound)
		// }

		json.NewEncoder(w).Encode(newUser)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
func addAddress(w http.ResponseWriter, r *http.Request) {
	refreshToken(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside addAddress")
	if isValidCoockie(w, r) {
		dbaddress, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect address database")
		}
		var allusers []User
		var address Address
		_ = json.NewDecoder(r.Body).Decode(&address)
		params := mux.Vars(r)
		fmt.Println(params)
		dbaddress.Find(&allusers)
		var uidofuser = ""
		for _, u := range allusers {
			if u.RollNo == params[("RollNo")] {
				uidofuser = u.UID
				break
			}
		}
		fmt.Println("Found UID")
		if uidofuser != "" {
			address.UID = uidofuser
			genrateUID := uuid.New()
			address.AddID = genrateUID.String()
			fmt.Println("Address Created")
			dbaddress.Create(&address)
			fmt.Println("Address Inserted")
			w.WriteHeader(200)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
	}

}
