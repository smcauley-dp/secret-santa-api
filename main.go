package main

import (
	"database/sql"
	"log"

	"github.com/KoseSoftware/secret-santa-api/config"
	"github.com/KoseSoftware/secret-santa-api/controllers"
	"github.com/KoseSoftware/secret-santa-api/repositories"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func main() {
	v := render.New()

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3307)/secret_santa_dev")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	lr := repositories.NewListRepository(db)

	hc := controllers.NewHomepageController(v)
	lc := controllers.NewListsController(lr, v)

	r := mux.NewRouter()
	r.HandleFunc("/", hc.Index).Methods("GET").Name("homepage")

	r.HandleFunc("/lists", lc.GetLists).Methods("GET").Name("get_lists")
	r.HandleFunc("/lists", lc.PostLists).Methods("POST").Name("post_lists")
	r.HandleFunc("/lists/{id:[0-9]+}", lc.GetList).Methods("GET").Name("get_list")

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(r)

	log.Fatal(config.GetServer(n).ListenAndServe())
}
