package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"whitesource_home_assingment/internal/movie"
)

type ginHandler = func(c *gin.Context)

func movieDeleteHandler(s *movie.Service) ginHandler {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := s.Delete(c, id)
		switch {
		case
			errors.Is(err, movie.ErrWrongIDFormat),
			errors.Is(err, movie.ErrNotFound),
			errors.Is(err, nil):
			c.JSON(http.StatusNoContent, msgOK)
			return
		default:
			c.JSON(http.StatusInternalServerError, msgErr(err))
			return
		}
	}
}

func movieInsertHandler(s *movie.Service) ginHandler {
	return func(c *gin.Context) {
		var m movieJSON
		if err := c.ShouldBindJSON(&m); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		mov, err := s.Add(c, m.toMovie())
		switch {
		case errors.Is(err, nil):
			c.JSON(http.StatusCreated, newMovieJSON(mov))
			return
		default:
			c.JSON(http.StatusInternalServerError, msgErr(err))
			return
		}
	}
}

func movieGetHandler(s *movie.Service) ginHandler {
	return func(c *gin.Context) {
		id := c.Param("id")
		mov, err := s.GetOne(c, id)
		switch {
		case errors.Is(err, nil):
			c.JSON(http.StatusOK, newMovieJSON(mov))
			return
		case
			errors.Is(err, movie.ErrNotFound),
			errors.Is(err, movie.ErrWrongIDFormat):
			c.JSON(http.StatusNotFound, msgErr(movie.ErrNotFound))
			return
		default:
			c.JSON(http.StatusInternalServerError, msgErr(err))
			return
		}
	}
}

func movieUpdateHandler(s *movie.Service) ginHandler {
	return func(c *gin.Context) {
		id := c.Param("id")
		var m movieJSON
		if err := c.ShouldBindJSON(&m); err != nil {
			c.JSON(http.StatusBadRequest, msgErr(err))
			return
		}
		err := s.Update(c, id, m.toMovie())
		switch {
		case errors.Is(err, nil):
			c.JSON(http.StatusOK, msgOK)
			return
		case
			errors.Is(err, movie.ErrNotFound),
			errors.Is(err, movie.ErrWrongIDFormat):
			c.JSON(http.StatusNotFound, msgErr(movie.ErrNotFound))
			return
		default:
			c.JSON(http.StatusInternalServerError, msgErr(err))
			return
		}
	}
}
