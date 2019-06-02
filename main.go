package main

import (
	"github.com/riyadennis/realty-tool/ex"
)

func main() {
	urls := ex.GetURLS("config.yaml")
	ex.Search(urls.Urls)
}
