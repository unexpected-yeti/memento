package database

// Datastore interface abstracts persistence layer.
type Datastore interface {
	// Retrieve a Meme
	Get(memeID int) (Meme, error)
	// Retrieve all stored Memes
	GetAll() ([]Meme, error)
	// Remove a Meme
	Delete(memeID int) error
	// Store a Meme and return its new ID
	Store(meme *Meme) error
	// Overwrite existing Meme
	Update(memeID int, meme Meme) error
	// Add new reaction to a Meme and return the Reaction ID
	AddReaction(memeID int, reaction *Reaction) error
	// Remove Reaction from a Meme
	RemoveReaction(memeID int, reactionID int) error
}
