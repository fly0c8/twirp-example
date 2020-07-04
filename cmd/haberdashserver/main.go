package main 

import(
	"github.com/fly0c8/twirp-example/internal/haberdashserver"
	"github.com/fly0c8/twirp-example/rpc/haberdasher"
	"net/http"

)


func main() {
	server := &haberdashserver.Server{}
	twirpHandler := haberdasher.NewHaberdasherServer(server, nil)
	http.ListenAndServe(":8080", twirpHandler)
}