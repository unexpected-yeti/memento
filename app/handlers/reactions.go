package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-zoo/bone"
	"github.com/unexpected-yeti/memento/app/database"
)

type Reactions struct {
	DB database.Datastore
}

func (reactions *Reactions) GetAllReactions(w http.ResponseWriter, r *http.Request) {

	// Get the meme memeId
	memeId, err := strconv.Atoi(bone.GetValue(r, "meme"))

	meme, err := reactions.DB.Get(memeId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	respondJSON(w, http.StatusOK, meme.Reactions)
}

func (reactions *Reactions) GetReaction(w http.ResponseWriter, r *http.Request) {
	memeId := GetValueAsInt(r, "meme")
	meme, err := reactions.DB.Get(memeId)
	if err != nil {
		respondNotFound(w, "meme not found.")
		return
	}

	reactionId := GetValueAsInt(r, "reaction")
	if reactionId > len(meme.Reactions) {
		respondNotFound(w, "reaction not found.")
		return
	}

	reactionId = reactionId - 1
	respondJSON(w, http.StatusOK, meme.Reactions[reactionId])
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
	memeId, err := strconv.Atoi(bone.GetValue(r, "meme"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	reaction := database.Reaction{Value: bone.GetValue(r, "value")}
	reactions.DB.AddReaction(memeId, &reaction)

	location := fmt.Sprintf("/memes/%d/reactions/%d", memeId, reaction.ID)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}
