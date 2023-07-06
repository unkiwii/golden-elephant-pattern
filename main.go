package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/unkiwii/golden-elephant-pattern/infra"
	"github.com/unkiwii/golden-elephant-pattern/usecase"
)

const (
	writeTimeoutSeconds = time.Second * 10
	readTimeoutSeconds  = time.Second * 5
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(200)
			fmt.Fprintf(w, "bad request: invalid user id")
		}

		var repo infra.UserRepo
		user, _ := repo.UserByID(id)

		w.WriteHeader(200)
		fmt.Fprintf(w, "%+v", user)
	}).Methods("GET")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		strIds := r.URL.Query()["id"]
		ids := make([]int64, len(strIds))
		var err error
		for i, strID := range strIds {
			ids[i], err = strconv.ParseInt(strID, 10, 64)
			if err != nil {
				w.WriteHeader(200)
				fmt.Fprintf(w, "bad request: invalid user id")
				return
			}
		}

		var provider usecase.UserProvider
		users, _ := provider.UsersByID(ids...)

		w.WriteHeader(200)
		fmt.Fprintf(w, "%+v", users)
	}).Methods("GET")

	server := http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: writeTimeoutSeconds,
		ReadTimeout:  readTimeoutSeconds,
	}

	log.Fatal(server.ListenAndServe())
}
