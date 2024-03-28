package main

import (
	"context"
	"github.com/gin-gonic/gin"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/openfeature"
	"log"
	"net/http"
)

const defaultMessage = "Hello!"
const newWelcomeMessage = "Hello, welcome to this OpenFeature-enabled website!"

func main() {
	// Use flagd as the OpenFeature provider
	err := openfeature.SetProviderAndWait(flagd.NewProvider())
	if err != nil {
		// If an error occurs, log it and exit
		log.Fatalf("Failed to set the OpenFeature provider: %v", err)
	}

	// Initialize OpenFeature client
	client := openfeature.NewClient("GoStartApp")

	// Initialize Go Gin
	engine := gin.Default()

	// Setup a simple endpoint
	engine.GET("/hello", func(ctx *gin.Context) {

		// Evaluate welcome-message feature flag
		welcomeMessage, _ := client.BooleanValue(
			context.Background(), "welcome-message", false, openfeature.EvaluationContext{},
		)

		if welcomeMessage {
			ctx.JSON(http.StatusOK, newWelcomeMessage)
			return
		} else {
			ctx.JSON(http.StatusOK, defaultMessage)
			return
		}
	})

	engine.GET("/feature", func(ctx *gin.Context) {
		ofctx := openfeature.NewEvaluationContext(
			"email",
			map[string]interface{}{
				"email": ctx.GetHeader("email"),
			},
		)

		feature, _ := client.StringValue(
			context.Background(), "my-feature", "Default feature", ofctx,
		)

		ctx.JSON(http.StatusOK, feature)
	})

	engine.GET("/color", func(ctx *gin.Context) {
		ofctx := openfeature.NewEvaluationContext(
			"email",
			map[string]interface{}{
				"email": ctx.GetHeader("email"),
			},
		)

		color, _ := client.StringValue(
			context.Background(), "color", "#FFFFFF", ofctx,
		)

		ctx.JSON(http.StatusOK, color)
	})

	engErr := engine.Run()
	if engErr != nil {
		log.Fatalf("The engine failed with: %v", engErr)
	}
}
