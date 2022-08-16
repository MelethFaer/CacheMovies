package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMovies(ctx *gin.Context) {
	data, _ := json.MarshalIndent(cacheObject.Movies, "", "    ")
	ctx.String(http.StatusOK, string(data))
}
