package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/goccy/go-json"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type bunnyFactResponse struct {
	Fact string `json:"fact"`
}

var bunnyFacts = [][]byte{
	jsonMustMarshal("A group of bunnies is called a fluffle!"),
	jsonMustMarshal("Happy bunnies do a jump-twist called a 'binky'!"),
	jsonMustMarshal("Bunnies can see nearly 360 degrees around them."),
	jsonMustMarshal("A baby bunny is called a kitten or kit."),
	jsonMustMarshal("Bunnies purr when they're happy, just like cats!"),
	jsonMustMarshal("A bunny's teeth never stop growing."),
	jsonMustMarshal("Bunnies can hop up to 3 feet high!"),
	jsonMustMarshal("Bunnies have over 100 million scent cells."),
	jsonMustMarshal("Bunnies can run up to 35 mph!"),
	jsonMustMarshal("Bunnies groom themselves like cats do."),
	jsonMustMarshal("Bunnies can learn to recognize their own names."),
	jsonMustMarshal("A bunny's ears can rotate 270 degrees!"),
	jsonMustMarshal("Bunnies sometimes sleep with their eyes open!"),
	jsonMustMarshal("The world's longest rabbit was over 4 feet long!"),
	jsonMustMarshal("Bunnies do zoomies when they're extra happy!"),
}

func jsonMustMarshal(fact string) []byte {
	b, err := json.Marshal(bunnyFactResponse{Fact: fact})
	if err != nil {
		panic(fmt.Errorf("could not json marshal struct: %w", err))
	}
	return b
}

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/x/bunny-fact", func(re *core.RequestEvent) error {
			re.Response.Header().Set("Content-Type", "application/json")
			fact := bunnyFacts[rand.Intn(len(bunnyFacts))]
			_, err := re.Response.Write(fact)
			return err
		})

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
