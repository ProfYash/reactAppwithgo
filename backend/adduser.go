package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var jwtKey = []byte("secretkey")
var adminUsers = map[string]string{
	"user1": "pass1",
	"user2": "pass2",
	"user3": "pass3",
}

type adminModel struct {
	gorm.Model
	Username string
	Password string
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string
	jwt.StandardClaims
}
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
	db, err := gorm.Open(sqlite.Open("admins.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&adminModel{})
	var aUs []adminModel
	db.Create(&adminModel{
		Username: "yashshah",
		Password: "hello",
	})
	db.Create(&adminModel{
		Username: "kanan",
		Password: "hello",
	})

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
	log.Fatal(http.ListenAndServe(":4002", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
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

func homePage(w http.ResponseWriter, r *http.Request) {
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
func refreshToken(w http.ResponseWriter, r *http.Request) {
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
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 30)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "tokenRefreshed",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func addUser(w http.ResponseWriter, r *http.Request) {
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
		user = append(user, newUser)

		fmt.Println(newUser)
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	fmt.Println("Inside getUser")
	if isValidCoockie(w, r) {
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside DeleteUSer")
	if isValidCoockie(w, r) {
		params := mux.Vars(r)
		fmt.Println(params)
		if getByRoll(params[("RollNo")]) {
			_deleteUserAtUid(params[("RollNo")])
			json.NewEncoder(w).Encode(user)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
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
	fmt.Println("Inside UpdateUser")
	if isValidCoockie(w, r) {
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

		json.NewEncoder(w).Encode(user)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
func getByRoll(RollNo string) bool {
	for _, u := range user {
		if RollNo == u.RollNo {
			return true
		}
	}
	return false
}
func isValidCoockie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}
