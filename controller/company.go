package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"github.com/solenovex/it/funcs"
	"github.com/solenovex/it/model"
)

func registerRoutes() {
	http.HandleFunc("/", listCompanies)
	http.HandleFunc("/devices", listCompanies)
	http.HandleFunc("/devices/seed", seed)
	http.HandleFunc("/devices/search", searchDevices)
	http.HandleFunc("/devices/search/", searchDevices) // 不知道为什么form全填提交后，url里search后会多一个“/”
	http.HandleFunc("/devices/add", addCompany)
	http.HandleFunc("/devices/edit/", editCompany)
	http.HandleFunc("/devices/delete/", deleteCompany)

}

func listCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := model.GetAllCompanies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		funcMap := template.FuncMap{"add": funcs.Add}
		t := template.New("companies").Funcs(funcMap)
		t, err = t.ParseFiles("./templates/_layout.html", "./templates/company/list.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		t.ExecuteTemplate(w, "layout", companies)
	}
}

func searchDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := model.GetSearchDevices(r.FormValue("assetno"), r.FormValue("devtype"), r.FormValue("devstatus"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		funcMap := template.FuncMap{"add": funcs.Add}
		t := template.New("devices").Funcs(funcMap)
		t, err = t.ParseFiles("./templates/_layout.html", "./templates/company/list.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		t.ExecuteTemplate(w, "layout", devices)
	}
}

func addCompany(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t := template.New("company-add")
		t, err := t.ParseFiles("./templates/_layout.html", "./templates/company/add.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			t.ExecuteTemplate(w, "layout", nil)
		}
	case http.MethodPost:
		newCompany := model.Device{}
		// newCompany.ID = r.PostFormValue("id")
		newCompany.AssetNo = r.PostFormValue("title")
		newCompany.DevType = r.PostFormValue("belongs")
		// newCompany.DevStaus = r.PostFormValue("devstate")
		// s := r.PostFormValue("devimage")
		// newCompany.Picture = []byte(s)
		err := newCompany.Insert()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			http.Redirect(w, r, "/companies", http.StatusSeeOther)
		}
	}
}

func editCompany(w http.ResponseWriter, r *http.Request) {
	idPattern := regexp.MustCompile(`/companies/edit/([a-zA-Z0-9]*$)`)
	matches := idPattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		id := matches[1]

		switch r.Method {
		case http.MethodGet:
			company, err := model.GetCompany(id)
			if err == nil {
				t := template.New("company-edit")
				t, err := t.ParseFiles("./templates/_layout.html", "./templates/company/edit.html")
				if err == nil {
					t.ExecuteTemplate(w, "layout", company)
					return
				}
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		case http.MethodPost:
			company := &model.Device{}
			company.ID = r.PostFormValue("id")
			company.AssetNo = r.PostFormValue("title")
			company.DevType = r.PostFormValue("belongs")
			// company.DevStaus = r.PostFormValue("devstate")
			// s := r.PostFormValue("devimage")
			// fmt.Printf("%v", s)
			// company.Picture = []byte(s)

			err := company.Update()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				http.Redirect(w, r, "/companies", http.StatusSeeOther)
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func deleteCompany(w http.ResponseWriter, r *http.Request) {
	idPattern := regexp.MustCompile(`/companies/delete/([a-zA-Z0-9]*$)`)
	matches := idPattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		id := matches[1]

		if r.Method == http.MethodDelete {
			err := model.DeleteCompany(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			http.Redirect(w, r, "/companies", http.StatusSeeOther)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func seed(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok")
}
