package main

import (
	_ "makesweet/docs"
	"makesweet/handlers"
	middleware "makesweet/middlewares"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title			Makesweet golang server
// @Version		1.0
// @Description	A golang server to create gifs from images.
// @Accept			mpfd
// @Produce		json image/gif
// @Host			localhost:8080
// @BasePath		/api
func main() {
	for _, envName := range []string{"SAVE_IMAGE_FOLDER", "SAVE_TEMPLATE_FOLDER"} {
		folderPath := os.Getenv(envName)
		if len(strings.TrimSpace(folderPath)) == 0 {
			log.Fatalf("%s environment variable invalid or not set", envName)
		}
		if err := os.MkdirAll(folderPath, 0o755); err != nil {
			log.Fatalf("Failed to create directory configured by %s: %v", envName, err)
		}
	}

	r := gin.Default()

	r.Use(middleware.RedirectToDocs)

	apiGroup := r.Group("/api")

	gifGroup := apiGroup.Group("/gif")
	gifGroup.POST("/billboard", handlers.CreateBillboardGif)
	gifGroup.POST("/flag", handlers.CreateFlagGif)
	gifGroup.POST("/heart-locket", handlers.CreateHeartLocketGif)
	gifGroup.POST("/circuit", handlers.CreateCircuitGif)
	gifGroup.POST("/nesting-doll", handlers.CreateDollGif)
	gifGroup.POST("/flying-bear", handlers.CreateBearGif)
	gifGroup.POST("/custom", handlers.CreateFromCustom)

	apiGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
