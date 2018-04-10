package app

import (
	"net/http"
	"github.com/go-zoo/bone"
	"github.com/unexpected-yeti/memento/app/handlers"
	"github.com/unexpected-yeti/memento/config"
)

type App struct {
	Handler http.Handler
}

func (app *App) Initialize(config *config.Config) {

	multiplexer := bone.New()

	// GET - Used for retrieving resources
	// Sets of memes
	multiplexer.Get("/memes", http.HandlerFunc(handlers.GetAllMemes))
	multiplexer.Get("/memes/top", http.HandlerFunc(handlers.GetTopMemes))
	multiplexer.Get("/memes/new", http.HandlerFunc(handlers.GetNewMemes))

	// Single memes
	multiplexer.Get("/memes/:meme", http.HandlerFunc(handlers.GetMeme))
	multiplexer.Get("/memes/random", http.HandlerFunc(handlers.GetRandomMeme))

	// Sets of reactions
	multiplexer.Get("/memes/:meme/reactions", http.HandlerFunc(handlers.GetAllReactions))

	// Single reactions
	multiplexer.Get("/memes/:meme/reactions/:reaction", http.HandlerFunc(handlers.GetReaction))

	// DELETE - Used for deleting resources
	// Single memes
	multiplexer.Delete("/memes/:meme", http.HandlerFunc(handlers.DeleteMeme))

	// Single reactions
	multiplexer.Delete("/memes/:meme/reactions/:reaction", http.HandlerFunc(handlers.DeleteReaction))

	// POST - Used for creating resources
	// Single memes
	multiplexer.Post("/memes", http.HandlerFunc(handlers.CreateMeme))

	// Single reactions
	multiplexer.Post("/memes:meme/reactions", http.HandlerFunc(handlers.CreateReaction))

	app.Handler = multiplexer
}