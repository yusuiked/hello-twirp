package main

import (
	"net/http"

	"github.com/yusuiked/hello-twirp/internal/haberdasherserver"
	"github.com/yusuiked/hello-twirp/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{}
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
