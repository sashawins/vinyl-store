//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./ent/schema", &gen.Config{})
	if err != nil {
		log.Fatalf("entc.Generate failed: %v", err)
	}
}
