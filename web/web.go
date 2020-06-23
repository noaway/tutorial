package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noaway/tutorial/service"
)

func InitWeb(us *service.UserService, is *service.ItemService) (http.Handler, error) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.Any("/", func(c *gin.Context) {
		fmt.Println(is.Name())
		c.JSON(200, us.GetUser())
	})

	return r, nil
}
