package server

import(
    "os"
    "log"
    "net/http"
    "institutions/config"
    "institutions/server/routes"
)

func Start() {
    port := os.Getenv(config.ENV_PORT)
    router := routes.GetRouter()
    
    if err := http.ListenAndServe(":" + port, router); err != nil {
        log.Fatal(err)
    }
}
