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

func TestGetEmptyMemeList(t *testing.T) {
	server := newTestServer()
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/memes").Expect().Status(http.StatusOK).JSON().Array().Empty()
}

func TestGetAllMemes(t *testing.T) {
	app := getApp()

	meme := database.Meme{Title: "foobar", ImageData: "test"}
	app.Datastore.Store(&meme)

	server := httptest.NewServer(app.Handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/memes").Expect().Status(http.StatusOK).JSON().
		Array().Length().Equal(1)
}

func TestGetMeme(t *testing.T) {
	app := getApp()

	meme := database.Meme{Title: "foobar", ImageData: "test"}
	app.Datastore.Store(&meme)

	server := httptest.NewServer(app.Handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	obj := e.GET("/memes/1").Expect().Status(http.StatusOK).JSON().Object()

	obj.Value("id").Equal(1)
	obj.Value("title").Equal("foobar")
	obj.Value("imageData").Equal("test")
	// TODO(claudio)
	// obj.Value("reactions").Array().Empty()
	obj.Value("reactions").Equal(nil)
}

func TestDeleteMeme(t *testing.T) {
	app := getApp()

	meme := database.Meme{Title: "foobar", ImageData: "test"}
	app.Datastore.Store(&meme)

	server := httptest.NewServer(app.Handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.DELETE("/memes/1").Expect().Status(http.StatusNoContent)
}
