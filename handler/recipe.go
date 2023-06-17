package handler

// import (
// 	"net/http"
// 	"webapp1/apis"
// 	"webapp1/apis/spoonacular"
// 	"webapp1/util"

// 	"github.com/gin-gonic/gin"
// )

// func HandlerGetRecipe(c *gin.Context) {
// 	currency := c.Query("currency")
// 	randomRecipe, err := spoonacular.GetRandomRecipes()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, util.BuildResponse("error", nil))
// 		return
// 	}

// 	if currency == "IDR" {
// 		idrValue, err := apis.GetCurrencyExchange()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, util.BuildResponse("error", nil))
// 			return
// 		}

// 		priceIdr := int(randomRecipe.PricePerServing * idrValue)
// 		randomRecipe.PricePerServing = float64(priceIdr)
// 	}

// 	c.JSON(http.StatusOK, util.BuildResponse("success get all recipes", randomRecipe))
// }