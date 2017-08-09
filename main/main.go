package main

import (
    "log"
    "runtime"

    "net/http"
    "database/sql"

    "route"
    "db_handler"
)

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())

	var err error
	db_handler.Db, err = sql.Open("mysql", "root:mel@tcp(localhost:3306)/kppm_staging?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
    router := route.NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}