package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"github.com/unexpected-yeti/memento/app/database"

	"github.com/go-zoo/bone"
)

type Memes struct {
	DB database.Datastore
}

func (memes *Memes) GetAllMemes(w http.ResponseWriter, r *http.Request) {
	all, err := memes.DB.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	js, err := json.Marshal(all)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (memes *Memes) GetTopMemes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}

func (memes *Memes) GetNewMemes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}

func (memes *Memes) GetMeme(w http.ResponseWriter, r *http.Request) {

	// Get the meme memeId
	memeId, err := strconv.Atoi(bone.GetValue(r, "meme"))

	// TODO(claudio): Check if meme was actually found
	meme, err := memes.DB.Get(memeId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	js, err := json.Marshal(meme)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (memes *Memes) GetRandomMeme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}

func (memes *Memes) DeleteMeme(w http.ResponseWriter, r *http.Request) {

	// Get the meme memeId
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	fmt.Fprint(w, r.URL.EscapedPath())
}

func (memes *Memes) CreateMeme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}
