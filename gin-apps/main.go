package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	"time"
)

var recipes []Recipe

func init()  {
	recipes = make([]Recipe, 0)
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

func NewRecipeHandler(c *gin.Context)  {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublisedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func ListRecipesHandler(c *gin.Context){
	c.JSON(http.StatusOK, recipes)
}

func main(){
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.Run()
}