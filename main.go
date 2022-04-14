package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"urlShortener/storage"
)

type Handler struct {
	Router *mux.Router
	Log *log.Logger
	DB *sql.DB
	isStorage string
}

func (h *Handler) InitAndRun() {
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/", h.Home)
	h.Router.HandleFunc("/urlShortener/", h.Post)
	h.Router.HandleFunc("/{key}", h.Get)
	fmt.Printf("Server is starting")

	err := http.ListenAndServe(":8000", h.Router)
	if err != nil {
		h.Log.Fatal()
	}
}

func main() {
	h := Handler{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		param := sc.Text()
		if param == "postgre" {
			h.DB, _ = storage.Connect()
		}
		h.isStorage = param
		h.InitAndRun()
	}
}