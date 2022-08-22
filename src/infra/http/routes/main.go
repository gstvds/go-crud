package routes

import "github.com/gorilla/mux"

func Generate() *mux.Router {
	muxRouter := mux.NewRouter()
	return Configure(muxRouter)
}
