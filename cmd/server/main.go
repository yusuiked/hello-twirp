package main

import (
    "net/http"

    "gitlab.com/yukung/hello-twirp/internal/haberdasherserver"
    "gitlab.com/yukung/hello-twirp/rpc/haberdasher"
)

func main() {
    server := &haberdasherserver.Server{}
    twirpHandler := haberdasher.NewHaberdasherServer(server)

    http.ListenAndServe(":8080", twirpHandler)
}
