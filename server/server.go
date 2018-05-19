package server

import(
    "os"
    "log"
    "net/http"
    "github.com/fiscaluno/hyoga/config"
    "github.com/fiscaluno/hyoga/server/routes"
)

func Start() {
    port := os.Getenv(config.ENV_PORT)
    router := routes.GetRouter()

    if err := http.ListenAndServe(":" + port, router); err != nil {
        log.Fatal(err)
    }
}
