package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
    "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
    "database/sql"
)
import _ "github.com/go-sql-driver/mysql"

type AuthRequest struct {
	Email string `json:"email"`
    Password string `json:"password"`
}

type AuthResponse struct {
	Status int `json:"status"`
    Message string `json:"message"`
}

type AttributeData struct {
    Id int64 `json:"id"`
	Name string `json:"name"`
    Title string `json:"title"`
    Type string `json:"attributetype"`
    CreatedDate string `json:"createdDate"`
    ModifiedDate string `json:"modifiedDate"`
    Status string `json:"status"`
}

type AttributeList []AttributeData

type CategoryData struct {
    Id int64 `json:"id"`
	Name string `json:"name"`
    CreatedDate string `json:"createdDate"`
    ModifiedDate string `json:"modifiedDate"`
    Status string `json:"status"`
}


type ReadUser struct{
    Oo_id int64 `json:"id"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	//router.HandleFunc("/login", authUser).Methods("POST")
    router.HandleFunc("/attribute/list", attributeList).Methods("GET")
    router.HandleFunc("/category/list", categoryList).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8085", router))
    log.Fatal(http.ListenAndServe(":8085", handlers.CORS()(router)))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func attributeList(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:root@/strykerpoc")
    failOnError(err, "")
    defer db.Close()

    // Open doesn't open a connection. Validate DSN data:
    err = db.Ping()
    failOnError(err, "")
	//log.Println(db)
    var (
    	id int64
    	name string
        title string
        attributetype string
        createdDate string
        modifiedDate string
        status string
    )
    stmtOut, err := db.Prepare("SELECT oo_id, Name, Title, AttributeType, DATE(FROM_UNIXTIME(o_creationDate)) as CreatedDate, DATE(FROM_UNIXTIME(o_modificationDate)) as ModifiedDate, o_published FROM object_5 ")
    failOnError(err, "")
    defer stmtOut.Close()
    rows, err := stmtOut.Query()
    failOnError(err, "")
    
    attributeDataList := []AttributeData{}
    
    for rows.Next() {
        err := rows.Scan(&id, &name, &title, &attributetype, &createdDate, &modifiedDate, &status) 
        failOnError(err, "")
        //log.Println(id, uuid, major, minor, uniqueId)
        attributeDataList = append(attributeDataList, AttributeData{
            id, name, title, attributetype, createdDate, modifiedDate, status,
        })
    }
    err = rows.Err()
    failOnError(err, "")

    w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Accept", "application/json")

	err = json.NewEncoder(w).Encode(attributeDataList)
	failOnError(err, "Issue in json encoding")
}

func categoryList(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:root@/strykerpoc")
    failOnError(err, "")
    defer db.Close()

    // Open doesn't open a connection. Validate DSN data:
    err = db.Ping()
    failOnError(err, "")
	//log.Println(db)
    var (
    	id int64
    	name string
        createdDate string
        modifiedDate string
        status string
    )
    stmtOut, err := db.Prepare("SELECT oo_id, Name, DATE(FROM_UNIXTIME(o_creationDate)) as CreatedDate, DATE(FROM_UNIXTIME(o_modificationDate)) as ModifiedDate, o_published FROM object_3 ")
    failOnError(err, "")
    defer stmtOut.Close()
    rows, err := stmtOut.Query()
    failOnError(err, "")
    
    categoryDataList := []CategoryData{}
    
    for rows.Next() {
        err := rows.Scan(&id, &name, &createdDate, &modifiedDate, &status) 
        failOnError(err, "")
        //log.Println(id, uuid, major, minor, uniqueId)
        categoryDataList = append(categoryDataList, CategoryData{
            id, name, createdDate, modifiedDate, status,
        })
    }
    err = rows.Err()
    failOnError(err, "")

    w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Accept", "application/json")

	err = json.NewEncoder(w).Encode(categoryDataList)
	failOnError(err, "Issue in json encoding")
}

/*func authUser(w http.ResponseWriter, r *http.Request) {
    urlParams := mux.Vars(r)
    email := urlParams["email"]
    password := urlParams["password"]
    
    db, err := sql.Open("mysql", "root:root@/strykerpoc")
    stmt := db.QueryRow("select oo_id from object_12 where Email = ? AND Password = ?", email, password)
    if err != nil {
        log.Fatal(err)
    }

    err = stmt.Scan(&ReadUser.oo_id)
    if err != nil {
        
        authResponse := AuthResponse{
            0, "Login Failed",
        } 
        w.Header().Set("Content-Type", "application/json")
	    //w.Header().Set("Accept", "application/json")

	    err := json.NewEncoder(w).Encode(authResponse)
	    failOnError(err, "Issue in json encoding")
        log.Fatal(err)
    } else {
        authResponse := AuthResponse{
            1, "Login Succeed",
        } 
        w.Header().Set("Content-Type", "application/json")
	    //w.Header().Set("Accept", "application/json")

	    err := json.NewEncoder(w).Encode(authResponse)
	    failOnError(err, "Issue in json encoding")
    }
    
    
}
*/
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
