package app

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gopkg.in/gavv/httpexpect.v1"

	"github.com/unexpected-yeti/memento/app/database"
	"github.com/unexpected-yeti/memento/config"
)

func newTestServer() *httptest.Server {
	app := getApp()
	return httptest.NewServer(app.Handler)
}

func getApp() App {
	configuration := config.Config{}
	application := App{}

	// Setup temporary fs-datastore
	dir, err := ioutil.TempDir("", "testfiles")
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewFSDatastore(dir)

	application.Initialize(&configuration, db)

	return application
}

func readFileToBase64(path string) (string, error) {
	handle, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer handle.Close()

	// Create a new buffer based on the file size
	info, _ := handle.Stat()
	size := info.Size()
	buf := make([]byte, size)

	// Read file into the buffer
	reader := bufio.NewReader(handle)
	reader.Read(buf)

	return base64.StdEncoding.EncodeToString(buf), nil
}

func TestAcceptPngMeme(t *testing.T) {
	server := newTestServer()
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	image, err := readFileToBase64("../resources/example.png")
	if err != nil {
		log.Fatal(err)
	}

	payload := map[string]interface{}{
		"title":     "test",
		"imageData": image,
	}

	// Ensure the URL to the created Meme is returned in the Location header
	e.POST("/memes").
		WithJSON(payload).
		Expect().
		Status(http.StatusCreated).
		Header("Location").
		Match("/memes/1")
}

func TestAcceptGifMeme(t *testing.T) {
	server := newTestServer()
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	image, err := readFileToBase64("../resources/example.gif")
	if err != nil {
		log.Fatal(err)
	}

	payload := map[string]interface{}{
		"title":     "test",
		"imageData": image,
	}

	e.POST("/memes").WithJSON(payload).Expect().Status(http.StatusCreated)
}

// GET /memes
func TestGetAllMemes(t *testing.T) {
	app := getApp()
	server := httptest.NewServer(app.Handler)
	defer server.Close()

	meme := database.Meme{Title: "foobar", ImageData: "test"}
	app.Datastore.Store(&meme)

	e := httpexpect.New(t, server.URL)

	e.GET("/memes").
		Expect().
		Status(http.StatusOK).
		JSON().Array().Length().Equal(1)
}

func TestGetEmptyMemeList(t *testing.T) {
	server := newTestServer()
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/memes").
		Expect().
		Status(http.StatusOK).
		JSON().Array().Empty()
}

// GET /meme/:id
func TestGetMeme(t *testing.T) {
	app := getApp()
	server := httptest.NewServer(app.Handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	app.Datastore.Store(&database.Meme{
		Title:     "foobar",
		ImageData: "test",
		Reactions: make([]database.Reaction, 0),
	})

	schema, _ := ioutil.ReadFile("schemas/meme.json")
	e.GET("/memes/1").
		Expect().
		Status(http.StatusOK).
		JSON().Schema(schema)
}

// DELETE /meme/:id
func TestDeleteMeme(t *testing.T) {
	app := getApp()
	server := httptest.NewServer(app.Handler)
	defer server.Close()

	meme := database.Meme{Title: "foobar", ImageData: "test"}
	app.Datastore.Store(&meme)

	e := httpexpect.New(t, server.URL)

	e.DELETE("/memes/1").
		Expect().
		Status(http.StatusNoContent)

	if meme, _ := app.Datastore.Get(1); meme.ID == 1 {
		t.Error("Meme should be deleted, but is present!")
	}
}

// POST /memes/:id/reactions
func TestAddReactionToMeme(t *testing.T) {
	app := getApp()
	server := httptest.NewServer(app.Handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	app.Datastore.Store(&database.Meme{
		Title:     "foobar",
		ImageData: "test",
	})

	payload := map[string]interface{}{
		"value": "high potential",
	}

	e.POST("/memes/1/reactions").WithJSON(payload).
		Expect().
		Status(http.StatusCreated).
		Header("Location").Match("/memes/1/reactions/1")
}

func TestGetReaction(t *testing.T) {
	app := getApp()
	server := httptest.NewServer(app.Handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	meme := database.Meme{
		Title:     "foobar",
		ImageData: "test",
	}

	app.Datastore.Store(&meme)
	app.Datastore.AddReaction(meme.ID, &database.Reaction{
		Value: "high potential",
	})

	obj := e.GET("/memes/1/reactions/1").
		Expect().
		JSON().Object()

	obj.Value("value").Equal("high potential")
}

func TestGetReaction404(t *testing.T) {
	app := getApp()
	server := httptest.NewServer(app.Handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	meme := database.Meme{
		Title:     "foobar",
		ImageData: "test",
	}

	app.Datastore.Store(&meme)

	e.GET("/memes/1/reactions/10").
		Expect().
		Status(http.StatusNotFound)
}

// DELETE /memes/:meme/reactions/:reaction
func TestDeleteReactionFromMeme(t *testing.T) {
	app := getApp()
	server := httptest.NewServer(app.Handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	meme := database.Meme{
		Title:     "foobar",
		ImageData: "test",
	}

	app.Datastore.Store(&meme)
	app.Datastore.AddReaction(meme.ID, &database.Reaction{Value: "-1"})

	e.DELETE("/memes/1/reactions/1").
		Expect().
		Status(http.StatusNoContent)

	meme, _ = app.Datastore.Get(1)
	if len(meme.Reactions) > 0 {
		t.Error("Reaction should be deleted, but is present!")
	}
}

// GET /memes/:meme/reactions
func TestGetReactions(t *testing.T) {
	app := getApp()
	server := httptest.NewServer(app.Handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	meme := database.Meme{
		Title:     "foobar",
		ImageData: "test",
	}

	app.Datastore.Store(&meme)

	reactions := []database.Reaction{
		database.Reaction{Value: "-1"},
		database.Reaction{Value: "2"},
		database.Reaction{Value: "high potential"},
	}

	for _, reaction := range reactions {
		app.Datastore.AddReaction(meme.ID, &reaction)
	}

	e.GET("/memes/1/reactions").
		Expect().
		JSON().Array().NotEmpty()
}
