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
    new_institution := map[string] interface{} {
        "Name": r.FormValue("name"),
        "Address": r.FormValue("address"),
        "Cnpj": r.FormValue("cnpj"),
    	"Email": r.FormValue("email"),
    	"Website": r.FormValue("website"),
    	"Phone": r.FormValue("phone"),
    	"ImageUri": r.FormValue("image"),
    }
    institution := institution.Create(new_institution)
    institution.Save()

    w.Write([]byte("Success"))
}
