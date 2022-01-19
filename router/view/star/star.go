package star

import (
	"meta-go/util/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router fiber.Router) {
	router.Get("index", starHandler.list(false))
	router.Get("list", starHandler.list(false))
	router.Get("detail", starHandler.detail())

	router.Post("media/upload", starHandler.mediaUpload())
	router.Get("media/:name", starHandler.getMedia())
	//
	my := router.Group("my/", middleware.LoginRequired())
	my.Get("list", starHandler.list(true))
	my.Get("add", starHandler.createStar())
	my.Post("add", starHandler.createStar())
	my.Get("edit", starHandler.updateStar())
	my.Post("edit", starHandler.updateStar())
}
