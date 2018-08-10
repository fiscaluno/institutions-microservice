package server

import(
    "os"
    "log"
    "net/http"
    "github.com/gorilla/handlers"
    "github.com/fiscaluno/hyoga/config"
    "github.com/fiscaluno/hyoga/server/routes"
    _ "github.com/fiscaluno/hyoga/database/migrations"
)

func Start() {
    router := routes.GetRouter()

    origins := handlers.AllowedOrigins([]string{"*"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

    port := os.Getenv("PORT")
    if port == "" {
        port = config.DEV_SERVER_PORT
    }

    if err := http.ListenAndServe(":" + port, handlers.CORS(origins, methods)(router));

    err != nil {
        log.Fatal(err)
    }
}
