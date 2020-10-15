package handler

import (
	"net/http"
	"time"

	"github.com/petersonsalme/bulletin-api/pkg/bulletin"
	"github.com/petersonsalme/bulletin-api/pkg/entity"

	"github.com/gin-gonic/gin"
)

func BulletinRouteCfg(e *gin.Engine, s *bulletin.BulletinService) {
	e.GET("/board", getAll(s))
	e.POST("/board", add(s))
}

func getAll(s *bulletin.BulletinService) gin.HandlerFunc {
	return func(context *gin.Context) {
		results, err := s.GetAll()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error: " + err.Error()})
			return
		}
		context.JSON(http.StatusOK, results)
	}
}

func add(s *bulletin.BulletinService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var b entity.Bulletin

		if context.Bind(&b) == nil {
			b.CreatedAt = time.Now()
			if err := s.Add(b); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error: " + err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	}
}
