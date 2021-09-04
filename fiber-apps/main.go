package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"time"
)

var recipes []Recipe

func init()  {
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

type Recipe struct {
	ID				string		`json:"id"`
	Name			string		`json:"name"`
	Tags			[]string	`json:"tags"`
	Ingredients		[]string	`json:"ingredients"`
	Instructions	[]string	`json:"instructions"`
	PublisedAt		time.Time	`json:"publised_at"`
}



func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(recipes)
	})

	app.Listen(":3000")
}