package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/riyadennis/realty-tool/ex"
)

func main() {
	urls := ex.GetURLS("config.yaml")
	ex.Search(urls.Urls)
}
