package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"urlShortener/storage"
	"urlShortener/utils"
)

//Post принимает оригинальную ссылку, возвращает сокращенную
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	inputUrl := r.URL.Query().Get("URL")
	if inputUrl != "" {
		key := utils.ByteMaskGen()

		switch h.isStorage {
		case "postgre":
			err := storage.Insert(key, inputUrl, h.DB)
			if err != nil {
				if err == storage.AlreadyExists {
					send(w,err.Error(), http.StatusOK)
					return
				} else if err == storage.InertFail {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		case "memory":
			err := storage.PostStorage(key, inputUrl)
			switch err {
			case nil:
				send(w,"500", http.StatusInternalServerError)
				return
			case storage.NotFound:
				send(w,err.Error(), http.StatusOK)
				return
			}
		default:
			h.Log.Fatal()
		}

		send(w, r.Host+"/"+key, http.StatusOK)
	} else {
		send(w, "enter the url", http.StatusOK)
	}
}

//Get принимает сокращенную, возвращает оригинальную
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["key"]

	switch h.isStorage {
	case "postgre":
		lookup, err := storage.Lookup(short, h.DB)
		if err != nil {
			return
		}
		http.Redirect(w, r, lookup, http.StatusSeeOther)
	case "memory":
		inputUrl, err := storage.GetStorage(short)
		if err != nil {
			send(w, "already exists", http.StatusOK)
		}
		http.Redirect(w, r, inputUrl, http.StatusSeeOther)
	default:
		h.Log.Fatal()
	}

}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	send(w, "use ?URL=site.ru", http.StatusOK)
}

func send(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(msg))
	if err != nil {
		return
	}
}

