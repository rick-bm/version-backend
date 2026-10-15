package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppVersion 是后端硬编码的语义化版本号（semver）。
const AppVersion = "1.0.0"

func versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": AppVersion})
}