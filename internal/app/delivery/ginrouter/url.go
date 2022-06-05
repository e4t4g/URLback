package ginrouter

import (
	"fmt"
	"github.com/e4t4g/URLback/cmd/app/config"
	"net/http"
	"strconv"

	"github.com/e4t4g/URLback/internal/app/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Delivery interface {
	Create() gin.HandlerFunc
	Redirect() gin.HandlerFunc
	GetStat() gin.HandlerFunc
}

type delivery struct {
	url    usecase.UseCase
	logger *zap.SugaredLogger
}

type URLData struct {
	ID       int    `json:"id" form:"id"`
	FullURL  string `json:"full_url" form:"full_url"`
	ShortURL string `json:"short_url" form:"short_url"`
	Counter  int64  `json:"counter" form:"counter"`
}

func New(url usecase.UseCase, logger *zap.SugaredLogger) Delivery {
	return delivery{url: url, logger: logger}
}

func (d delivery) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		d.logger.Info("Create")
		cfg := config.Config{}

		err := cfg.ReadFromFile(d.logger)
		if err != nil {
			return
		}

		host := cfg.Host
		port := cfg.Port

		var newURL *URLData
		if err = c.ShouldBind(&newURL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		}

		result, err := d.url.Create(c.Request.Context(), (*usecase.URLData)(newURL))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "unable to create"})
		}

		shortURL := fmt.Sprintf("%s:%d/%s", host, port, result.ShortURL)
		statURL := fmt.Sprintf("%s:%d/stat/%d", host, port, result.ID)

		p := URLData{FullURL: statURL, ShortURL: shortURL}
		c.JSON(http.StatusOK, p)
	}
}

func (d delivery) Redirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		d.logger.Info("Redirect")
		token := c.Param("token")

		redirectStruct, err := d.url.Redirect(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "incorrect token format"})
		}
		c.Redirect(http.StatusMovedPermanently, redirectStruct.FullURL)
	}
}

func (d delivery) GetStat() gin.HandlerFunc {
	return func(c *gin.Context) {
		d.logger.Info("GetStat")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "incorrect ID"})
		}

		redirectStruct, err := d.url.GetStat(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})

		}

		p := URLData{Counter: redirectStruct.Counter}
		c.JSON(http.StatusOK, p.Counter)
	}
}



