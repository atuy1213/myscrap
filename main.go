package main

import (
	"net/http"

	"github.com/atuy1213/myscrap/myscrap-api/driver/db"
	"github.com/atuy1213/myscrap/myscrap-api/internal/di"
)

func main() {
	db.InitDB()

	authenticationController := di.InjectAuthenticationController()
	http.HandleFunc("/signup", authenticationController.Signup)

	clipController := di.InjectClipController()
	http.HandleFunc("/clip", clipController.ClipArticle)

	http.ListenAndServe(":8080", nil)
}
