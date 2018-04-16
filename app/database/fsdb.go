package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// FSDatastore uses the local filesystem for persistence.
type FSDatastore struct {
	RootDir string
}

// NewFSDatastore initializes a new FSDatastore
func NewFSDatastore(rootPath string) *FSDatastore {
	return &FSDatastore{
		RootDir: rootPath,
	}
}

// Get returns a Meme from the filesystem
func (db *FSDatastore) Get(memeID int) (Meme, error) {
	var meme Meme
	fname := db.getFilename(memeID)

	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		return meme, err
	}

	err = json.Unmarshal(dat, &meme)
	if err != nil {
		return meme, err
	}

	return meme, nil
}

// Delete removes a Meme from the filesystem.
func (db *FSDatastore) Delete(memeID int) error {
	fname := db.getFilename(memeID)
	err := os.Remove(fname)
	if err != nil {
		return err
	}

	return nil
}

// Store persists a Meme as JSON to the local filesystem.
func (db *FSDatastore) Store(meme *Meme) error {
	meme.ID = db.getNextID()
	fname := db.getFilename(meme.ID)
	err := db.write(fname, *meme)
	if err != nil {
		return err
	}
	return nil
}

// Update will overwrite an existing Meme.
func (db *FSDatastore) Update(meme Meme) error {
	fname := db.getFilename(meme.ID)
	err := db.write(fname, meme)
	if err != nil {
		return err
	}

	return nil
}

// AddReaction adds a reaction to a Meme
func (db *FSDatastore) AddReaction(memeID int, reaction *Reaction) error {
	meme, err := db.Get(memeID)
	if err != nil {
		return err
	}

	reaction.ID = len(meme.Reactions) + 1
	meme.Reactions = append(meme.Reactions, *reaction)

	db.Update(meme)
	return nil
}

// RemoveReaction will remove a reaction from a Meme.
func (db *FSDatastore) RemoveReaction(memeID int, reactionID int) error {
	index := (reactionID - 1)
	meme, err := db.Get(memeID)
	if err != nil {
		return err
	}
	meme.Reactions = append(meme.Reactions[:index], meme.Reactions[index+1:]...)
	db.Update(meme)
	return nil
}

// write stores a Meme as JSON file to disk
func (db *FSDatastore) write(fname string, meme Meme) error {
	data, err := json.Marshal(meme)
	if err != nil {
		log.Fatalf("Unable to marshal Meme: %s", err.Error())
		return err
	}

	err = ioutil.WriteFile(fname, data, 0644)
	if err != nil {
		log.Fatalf("Unable to write Meme to %s", fname)
		return err
	}

	log.Printf("Wrote Meme with ID %d to %s\n", meme.ID, fname)
	return nil
}

// getNextID returns next available ID.
func (db *FSDatastore) getNextID() int {
	files, err := ioutil.ReadDir(db.RootDir)
	if err != nil {
		log.Fatal(err)
	}

	var highestID = 0
	for _, f := range files {
		s := strings.Split(f.Name(), ".")
		id, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}

		if id > highestID {
			highestID = id
		}
	}
	return (highestID + 1)
}

// getFilename constructs the filepath for a persisted Meme.
func (db *FSDatastore) getFilename(memeID int) string {
	return fmt.Sprintf("%s/%d.json", db.RootDir, memeID)
}
