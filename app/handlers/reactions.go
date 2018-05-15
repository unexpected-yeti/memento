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

// CreateReaction creates a reaction and attaches it to given Meme.
func (reactions *Reactions) CreateReaction(w http.ResponseWriter, r *http.Request) {
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

// DeleteReaction deletes a reaction from a given Meme.
// TODO
func (reactions *Reactions) DeleteReaction(w http.ResponseWriter, r *http.Request) {
	memeId, err := strconv.Atoi(bone.GetValue(r, "meme"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	reactionId, err := strconv.Atoi(bone.GetValue(r, "reaction"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = reactions.DB.RemoveReaction(memeId, reactionId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetReaction returns a reaction by ID of given Meme.
//
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

// GetAllReactions returns all reactions on a given Meme.
func (reactions *Reactions) GetAllReactions(w http.ResponseWriter, r *http.Request) {
	memeId, err := strconv.Atoi(bone.GetValue(r, "meme"))
	meme, err := reactions.DB.Get(memeId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	respondJSON(w, http.StatusOK, meme.Reactions)
}
