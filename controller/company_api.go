package controller

import (
	"encoding/json"
	"net/http"

	"github.com/solenovex/it/model"
)

func registerAPIRoutes() {
	http.HandleFunc("/api/companies", getCompanies)
}

func getCompanies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		companies, err := model.GetAllCompanies()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			enc := json.NewEncoder(w)
			err = enc.Encode(companies)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		c := model.Device{}
		err := dec.Decode(&c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			err = c.Insert()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}
	}
}
