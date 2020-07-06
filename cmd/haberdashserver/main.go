package main 

import(
	"github.com/fly0c8/twirp-example/internal/haberdashserver"
	"github.com/fly0c8/twirp-example/rpc/haberdasher"
	"net/http"

)


func main() {
	server := &haberdashserver.Server{}
	twirpHandler := haberdasher.NewHaberdasherServer(server, nil)	

	// note: FileServer handles index.html in a different way than other html files
	// => curl localhost:8080/index.html => redirects to ./
	fs := http.FileServer(http.Dir("./www"))
	
	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	mux.Handle("/", http.StripPrefix("/", fs))
	
	http.ListenAndServe(":8080", mux)

}