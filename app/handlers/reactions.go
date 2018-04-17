package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/unexpected-yeti/memento/app/database"
)

type Reactions struct {
	DB database.Datastore
}

func (reactions *Reactions) GetAllReactions(w http.ResponseWriter, r *http.Request) {

	// Get the meme memeId
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	fmt.Fprint(w, r.URL.EscapedPath())
}

func (reactions *Reactions) GetReaction(w http.ResponseWriter, r *http.Request) {

	// Get the meme id
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	// Get the reaction id
	reactionId := bone.GetValue(r, "reaction")

	w.Write([]byte(reactionId))

	fmt.Fprint(w, r.URL.EscapedPath())
}

func (reactions *Reactions) DeleteReaction(w http.ResponseWriter, r *http.Request) {

	// Get the meme id
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	// Get the reaction id
	reactionId := bone.GetValue(r, "reaction")

	w.Write([]byte(reactionId))

	fmt.Fprint(w, r.URL.EscapedPath())
}

func (reactions *Reactions) CreateReaction(w http.ResponseWriter, r *http.Request) {

	// Get the meme memeId
	memeId := bone.GetValue(r, "meme")

	w.Write([]byte(memeId))

	fmt.Fprint(w, r.URL.EscapedPath())
}
