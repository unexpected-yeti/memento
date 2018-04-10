package test

import (
	"bufio"
	"encoding/base64"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gopkg.in/gavv/httpexpect.v1"

	"github.com/unexpected-yeti/memento/app"
	"github.com/unexpected-yeti/memento/config"
)

func getAppHandler() http.Handler {

	configuration := config.Config{}
	application := app.App{}

	application.Initialize(&configuration)

	return application.Handler
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

	handler := getAppHandler()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	image, err := readFileToBase64("./example.png")
	if err != nil {
		log.Fatal(err)
	}

	payload := map[string]interface{}{
		"title":     "test",
		"imageData": image,
	}

	e.POST("/memes").WithJSON(payload).Expect().Status(http.StatusOK)
}

func TestAcceptGifMeme(t *testing.T) {

	handler := getAppHandler()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	image, err := readFileToBase64("./example.gif")
	if err != nil {
		log.Fatal(err)
	}

	payload := map[string]interface{}{
		"title":     "test",
		"imageData": image,
	}

	e.POST("/memes").WithJSON(payload).Expect().Status(http.StatusOK)
}
