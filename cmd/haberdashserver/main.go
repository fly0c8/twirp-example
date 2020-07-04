package main 

import(
	"github.com/fly0c8/haberdasher/internal/haberdashserver"
)


func main() {
	server := &haberdashserver.Server{}
	twirpHandler := haberdasher.NewHaberdashServer(server, nil)
	http.ListenAndServe(":8080", twirpHandler)
}