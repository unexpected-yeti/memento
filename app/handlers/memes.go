package handlers

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/unexpected-yeti/memento/app/database"

	"github.com/go-zoo/bone"
)

type Memes struct {
	DB database.Datastore
}

// CreateMeme creates a new Meme and persists it.
func (memes *Memes) CreateMeme(w http.ResponseWriter, r *http.Request) {
	title := bone.GetValue(r, "title")
	imageData := bone.GetValue(r, "imageData")

	meme := database.Meme{Title: title, ImageData: imageData}
	err := memes.DB.Store(&meme)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/memes/%d", meme.ID))
	w.WriteHeader(http.StatusCreated)
}

// DeleteMeme deletes a Meme from the datastore.
func (memes *Memes) DeleteMeme(w http.ResponseWriter, r *http.Request) {
	memeId, err := strconv.Atoi(bone.GetValue(r, "meme"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = memes.DB.Delete(memeId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetMeme returns a Meme by ID.
func (memes *Memes) GetMeme(w http.ResponseWriter, r *http.Request) {
	memeId, err := strconv.Atoi(bone.GetValue(r, "meme"))

	// TODO(claudio): Check if meme was actually found, else return 404
	meme, err := memes.DB.Get(memeId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	respondJSON(w, http.StatusOK, meme)
}

// GetAllMemes will return all stored Memes.
func (memes *Memes) GetAllMemes(w http.ResponseWriter, r *http.Request) {
	all, err := memes.DB.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	respondJSON(w, http.StatusOK, all)
}

// TODO
func (memes *Memes) GetTopMemes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}

// TODO
func (memes *Memes) GetNewMemes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}

// TODO
func (memes *Memes) GetRandomMeme(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.EscapedPath())
}
