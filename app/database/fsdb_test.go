package app

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

type TestDB struct {
	*FSDatastore
}

// setupTestDB creates a temporary database for testing purposes
func setupTestDB() *TestDB {
	dir, err := ioutil.TempDir("", "testfiles")
	if err != nil {
		log.Fatal(err)
	}

	db := NewFSDatastore(dir)
	return &TestDB{db}
}

// Clean up test database
func (db *TestDB) Close() {
	os.RemoveAll(db.RootDir)
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func TestGet(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	meme := Meme{ImageData: "binary...", Title: "MyTitle"}
	db.Store(&meme)

	stored, err := db.Get(meme.ID)
	if err != nil {
		log.Fatal(err)
	}

	if !reflect.DeepEqual(stored, meme) {
		t.Errorf("Actual and stored memes were not equal, \ngot: %+v,\nwant %+v", stored, meme)
	}

}

// If a Meme can not be found, return error
func TestGetNotExistingMeme(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	if _, err := db.Get(1337); err == nil {
		t.Errorf("err must not be nil")
	}
}

func TestStore(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	meme := Meme{ImageData: "asdf", Title: "MyTitle"}
	db.Store(&meme)

	if meme.ID != 1 {
		t.Errorf("meme ID was incorrect, got: %d, want: %d", meme.ID, 1)
	}

	var fname = db.getFilename(meme.ID)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		t.Errorf("meme file at %s does not exist", fname)
	}
}

func TestGetNextID(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	db.Store(&Meme{ImageData: "asdf", Title: "MyTitle"})

	nextID := db.getNextID()
	if nextID != 2 {
		t.Errorf("getNextID was incorrect, got: %d, want: %d", nextID, 2)
	}
}

func TestGetNextID_EmptyDatastore(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	if id := db.getNextID(); id != 1 {
		t.Errorf("getNextID must start at 1.")
	}
}

func TestDelete(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	meme := Meme{ImageData: "asdf", Title: "MyTitle"}

	db.Store(&meme)
	db.Delete(meme.ID)

	fname := db.getFilename(meme.ID)
	if exists, _ := exists(fname); exists {
		t.Errorf("file must not exist")
	}
}

func TestAddReaction(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	meme := Meme{ImageData: "asdf", Title: "MyTitle"}
	db.Store(&meme)

	reaction := Reaction{Value: "High Potential"}
	db.AddReaction(meme.ID, &reaction)

	if reaction.ID != 1 {
		t.Errorf("reaction ID was incorrect, got: %d, want: %d", reaction.ID, 1)
	}

	storedMeme, err := db.Get(meme.ID)
	if err != nil {
		panic(err)
	}

	storedReaction := storedMeme.Reactions[0]
	if !reflect.DeepEqual(reaction, storedReaction) {
		t.Errorf("Reactions were not equal, \ngot: %+v,\nwant %+v", storedReaction, reaction)
	}
}

func TestRemoveReaction(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	meme := Meme{ImageData: "asdf", Title: "MyTitle"}
	reaction := Reaction{Value: "High Potential"}

	db.Store(&meme)
	db.AddReaction(meme.ID, &reaction)

	db.RemoveReaction(meme.ID, reaction.ID)

	meme, _ = db.Get(meme.ID)
	if len(meme.Reactions) > 0 {
		t.Errorf("Reactions must be empty!")
	}
}
