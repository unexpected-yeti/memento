package database

// Meme struct is the representation of a Meme in Memento.
type Meme struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	ImageData string     `json:"imageData"`
	Reactions []Reaction `json:"reactions"`
}

// Reaction represents reactions on a Meme.
type Reaction struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}
