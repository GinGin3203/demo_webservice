package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whitesource_home_assingment/internal/movie"
	"whitesource_home_assingment/pkg/must"
)

type Server struct {
	MovieService movie.Service
}

func (s *Server) Run(router *gin.Engine, addr, certPath, keyPath string) {
	must.NoError(s.NewRoutesV1(router).RunTLS(addr, certPath, keyPath))
}

func (s *Server) NewRoutesV1(router *gin.Engine) *gin.Engine {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, msgOK)
	})
	v1 := router.Group("/v1/")

	s.addGroupMovieID(v1)
	s.addGroupMovie(v1)

	return router
}

func (s *Server) addGroupMovie(rg *gin.RouterGroup) {
	rg.POST("/movie", movieInsertHandler(&s.MovieService))
}

func (s *Server) addGroupMovieID(rg *gin.RouterGroup) {
	rg = rg.Group("/movie/")
	rg.GET("/:id", movieGetHandler(&s.MovieService))
	rg.PUT("/:id", movieUpdateHandler(&s.MovieService))
	rg.DELETE("/:id", movieDeleteHandler(&s.MovieService))
}
