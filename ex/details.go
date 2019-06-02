package ex

import (
	"io"
	"log"
	"net/http"
	"os"
)

func viewProperty(url string) {
	respPid, err := http.Get(url)
	if err != nil {
		log.Fatalf("error while loading the property :: %v", err)
	}
	defer respPid.Body.Close()
	io.Copy(os.Stdout, respPid.Body)
}
