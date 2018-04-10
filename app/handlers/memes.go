package handlers

import (
	"net/http"
	"fmt"
	"github.com/go-zoo/bone"
)

func GetAllMemes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}

func GetTopMemes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}

func GetNewMemes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}

func GetMeme(w http.ResponseWriter, r *http.Request) {

	// Get the meme memeId
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	fmt.Fprint(w, r.URL.EscapedPath())
}

func GetRandomMeme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}

func DeleteMeme(w http.ResponseWriter, r *http.Request) {

	// Get the meme memeId
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	fmt.Fprint(w, r.URL.EscapedPath())
}

func CreateMeme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}