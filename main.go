package main

import (
	"context"
	"fmt"
	"log"

	"github.com/prochac/huge-test-deps/supersql"
)

func main() {
	db, err := supersql.Connect(context.TODO(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println(db.Stats())
}
