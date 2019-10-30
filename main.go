package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	"github.com/riyadennis/realty-tool/ex"
)

func main() {
	urls := ex.GetURLS("config.yaml")
	if urls != nil {
		records := ex.Search(urls.Urls)
		if len(records) == 0 {
			panic("no records")
		}
		err := ex.CheckAndUpdate(records)
		if err != nil {
			panic(err)
		}
	}
}
