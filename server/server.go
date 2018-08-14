package server

import(
    "log"
    "net/http"
    "github.com/gorilla/handlers"
    "github.com/fiscaluno/institutions-microservice/config"
    "github.com/fiscaluno/institutions-microservice/server/routes"
    _ "github.com/fiscaluno/institutions-microservice/database/migrations"
)

func Start() {
    router := routes.GetRouter()

    origins := handlers.AllowedOrigins([]string{"*"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

    if err := http.ListenAndServe(":" + config.DEV_SERVER_PORT, handlers.CORS(origins, methods)(router));

    err != nil {
        log.Fatal(err)
    }
}
