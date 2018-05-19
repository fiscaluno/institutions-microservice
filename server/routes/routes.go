package routes

import(
    "github.com/gorilla/mux"
)

func GetRouter() (router *mux.Router) {
    router = mux.NewRouter()
    return
}
