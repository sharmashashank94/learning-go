package main

import (
    "fmt"
    "database/sql"
    "net/http"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    _ "github.com/lib/pq"
)

var router *chi.Mux
var db *sql.DB
var val string

const (
  host     = "localhost"
  port     = 5432
  user     = "shashanksharma"
  dbname   = "development8"
)

type Post struct {
    ID      int    `json: "id"`
    FirstName   string `json: "first_name"`
    Email string `json: "email"`
}


func routers() *chi.Mux {
    fmt.Println("Database at router:")
    fmt.Println(db)
    router.Get("/posts", AllPosts)
    //router.Get("/posts/{id}", DetailPost)
    //router.Post("/posts", CreatePost)
    //router.Put("/posts/{id}", UpdatePost)
    //router.Delete("/posts/{id}", DeletePost)
    return router
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
    errors := []error{}
    payload := []Post{}

    fmt.Println("Database:")
    fmt.Println(db)
    fmt.Println(val)

    rows, err := db.Query("Select id, first_name, email from users limit 10")
    fmt.Println("I am coming here")

    catch(err)

    defer rows.Close()

    for rows.Next() {
        data := Post{}

        er := rows.Scan(&data.ID, &data.FirstName, &data.Email)

        if er != nil {
            errors = append(errors, er)
        }
        payload = append(payload, data)
    }

    respondwithJSON(w, http.StatusOK, payload)
}

func init() { 
    router = chi.NewRouter() 
    router.Use(middleware.Recoverer)
}

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "dbname=%s sslmode=disable",
        host, port, user, dbname)

    val = "220"

    db, err := sql.Open("postgres", psqlInfo)
    
    catch(err)
    defer db.Close()

    err = db.Ping()
    catch(err)

    fmt.Println("Successfully connected!")
    
    fmt.Println("Database before router call")
    fmt.Println(db)

    routers()

    fmt.Println("Database in main:")
    fmt.Println(db)
    http.ListenAndServe(":8005", Logger())
}