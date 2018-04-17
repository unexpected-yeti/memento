package app

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/unexpected-yeti/memento/app/database"
	"github.com/unexpected-yeti/memento/config"

	"github.com/unexpected-yeti/memento/app/handlers"
)

type App struct {
	Handler   http.Handler
	Datastore database.Datastore
}

func (app *App) Initialize(config *config.Config, db database.Datastore) {

	multiplexer := bone.New()

	app.Datastore = db

	memes := handlers.Memes{DB: db}
	reactions := handlers.Reactions{DB: db}

	// GET - Used for retrieving resources
	// Sets of memes
	multiplexer.Get("/memes", http.HandlerFunc(memes.GetAllMemes))
	multiplexer.Get("/memes/top", http.HandlerFunc(memes.GetTopMemes))
	multiplexer.Get("/memes/new", http.HandlerFunc(memes.GetNewMemes))

	// Single memes
	multiplexer.Get("/memes/:meme", http.HandlerFunc(memes.GetMeme))
	multiplexer.Get("/memes/random", http.HandlerFunc(memes.GetRandomMeme))

	// Sets of reactions
	multiplexer.Get("/memes/:meme/reactions", http.HandlerFunc(reactions.GetAllReactions))

	// Single reactions
	multiplexer.Get("/memes/:meme/reactions/:reaction", http.HandlerFunc(reactions.GetReaction))

	// DELETE - Used for deleting resources
	// Single memes
	multiplexer.Delete("/memes/:meme", http.HandlerFunc(memes.DeleteMeme))

	// Single reactions
	multiplexer.Delete("/memes/:meme/reactions/:reaction", http.HandlerFunc(reactions.DeleteReaction))

	// POST - Used for creating resources
	// Single memes
	multiplexer.Post("/memes", http.HandlerFunc(memes.CreateMeme))

	// Single reactions
	multiplexer.Post("/memes:meme/reactions", http.HandlerFunc(reactions.CreateReaction))

	app.Handler = multiplexer
}
