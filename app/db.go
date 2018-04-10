package app

// Datastore interface abstracts persistence layer.
type Datastore interface {
	// Retrieve a Meme
	Get(memeID int) Meme
	// Remove a Meme
	Delete(memeID int)
	// Store a Meme and return its new ID
	Store(meme Meme) int
	// Overwrite existing Meme
	Update(memeID int, meme Meme)
	// Add new reaction to a Meme and return the Reaction ID
	AddReaction(memeID int, reaction Reaction) int
	// Remove Reaction from a Meme
	RemoveReaction(memeID int, reactionID int)
}
