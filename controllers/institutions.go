package controllers

import(
    "encoding/json"
    "net/http"
    "github.com/fiscaluno/hyoga/models/institution"
)

func All(w http.ResponseWriter, r *http.Request) {
    institutions := institution.All()

    response, _ := json.Marshal(institutions)
    w.Write([]byte(response))
}

func New(w http.ResponseWriter, r *http.Request) {
    // _ := "teste"
    w.Write([]byte(r.FormValue("name")))
}

func Delete(w http.ResponseWriter, r *http.Request) {
    // institution := "institution"
}

func Update(w http.ResponseWriter, r *http.Request) {
    // institution := "institution"
}
