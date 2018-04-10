package handlers

import (
	"net/http"
	"fmt"
	"github.com/go-zoo/bone"
)

func GetAllReactions(w http.ResponseWriter, r *http.Request) {

	// Get the meme memeId
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	fmt.Fprint(w, r.URL.EscapedPath())
}

func GetReaction(w http.ResponseWriter, r *http.Request) {

	// Get the meme id
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	// Get the reaction id
	reactionId := bone.GetValue(r, "reaction")

	w.Write([]byte(reactionId))

	fmt.Fprint(w, r.URL.EscapedPath())
}

func DeleteReaction(w http.ResponseWriter, r *http.Request) {

	// Get the meme id
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	// Get the reaction id
	reactionId := bone.GetValue(r, "reaction")

	w.Write([]byte(reactionId))

	fmt.Fprint(w, r.URL.EscapedPath())
}

func CreateReaction(w http.ResponseWriter, r *http.Request) {

	// Get the meme memeId
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	fmt.Fprint(w, r.URL.EscapedPath())
}