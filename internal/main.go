package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	chessGameController := ServiceContainer().ChessGameController()

	r := mux.NewRouter()
	r.HandleFunc("/game/chess/make-move", chessGameController.MakeMove)
	http.Handle("/game/chess/make-move", r)

	fmt.Printf("call ListenAndServe: %v", http.ListenAndServe(":8080", r))
}
