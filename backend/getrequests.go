package main

import (
	//structure "backend/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	//refreshToken(w, r)
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Hello, %s", claims.Username)
}
func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//refreshToken(w, r)
	fmt.Println("Inside getUser")
	if isValidCoockie(w, r) {
		db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		var usersForDisplay []User
		var addressofuser []Address
		var dto []DTO
		var singledto DTO
		var count = 0
		db.Find(&usersForDisplay)
		db.Find(&addressofuser)
		for _, u := range usersForDisplay {
			singledto.UID = u.UID
			singledto.RollNo = u.RollNo
			singledto.FName = u.FName
			for _, a := range addressofuser {
				if a.UID == u.UID {
					count = count + 1
				}
			}
			singledto.CountOfAddress = strconv.Itoa(count)
			dto = append(dto, singledto)
			count = 0
		}
		db.Create(&dto)
		json.NewEncoder(w).Encode(dto)
		for i, _ := range usersForDisplay {
			usersForDisplay[i] = User{}
		}
		for i, _ := range dto {
			dto[i] = DTO{}
		}
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
func getUsersAddress(w http.ResponseWriter, r *http.Request) {
	//refreshToken(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside getUsersAddress")
	if isValidCoockie(w, r) {
		dbaddress, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect address database")
		}
		var allusers []User
		var addressfordisplay []Address
		var userid = ""
		dbaddress.Find(&allusers)
		params := mux.Vars(r)
		fmt.Println("this is roll", params)
		for _, u := range allusers {
			if u.RollNo == params[("RollNo")] {
				userid = u.UID
				break
			}
		}
		if userid != "" {
			dbaddress.Where("uid = ?", userid).Find(&addressfordisplay)
			json.NewEncoder(w).Encode(addressfordisplay)
			w.WriteHeader(200)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
	}
}
