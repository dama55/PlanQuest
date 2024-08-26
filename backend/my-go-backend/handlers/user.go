package handlers

import (
    "net/http"
    "my-go-backend/services"
    "github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
    user, err := services.GetUser()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.String(http.StatusOK, "User: %s\n", user)
}