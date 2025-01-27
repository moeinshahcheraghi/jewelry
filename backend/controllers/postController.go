package controllers

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "jewelry/backend/config"
    "jewelry/backend/models"
)

func GetPosts(c *gin.Context) {
    var posts []models.Post
    if err := config.DB.Find(&posts).Error; err != nil {
        log.Println("Error fetching posts:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
        return
    }
    log.Println("Fetched posts successfully")
    c.JSON(http.StatusOK, gin.H{"data": posts})
}

func CreatePost(c *gin.Context) {
    var input struct {
        Title   string `json:"title" binding:"required"`
        Content string `json:"content" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        log.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    post := models.Post{Title: input.Title, Content: input.Content}
    if err := config.DB.Create(&post).Error; err != nil {
        log.Println("Error creating post:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating post"})
        return
    }
    log.Printf("Post created successfully: %+v\n", post)
    c.JSON(http.StatusOK, gin.H{"data": post})
}

func GetPost(c *gin.Context) {
    var post models.Post
    if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
        log.Println("Post not found:", err)
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }
    log.Printf("Fetched post successfully: %+v\n", post)
    c.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
    var post models.Post
    if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
        log.Println("Post not found for update:", err)
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    var input struct {
        Title   string `json:"title"`
        Content string `json:"content"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        log.Println("Error binding JSON for update:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Model(&post).Updates(input).Error; err != nil {
        log.Println("Error updating post:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating post"})
        return
    }

    if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
        log.Println("Error retrieving updated post:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving updated post"})
        return
    }

    log.Printf("Post updated successfully: %+v\n", post)
    c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
    var post models.Post
    if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
        log.Println("Post not found for deletion:", err)
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    if err := config.DB.Delete(&post).Error; err != nil {
        log.Println("Error deleting post:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting post"})
        return
    }

    log.Printf("Post deleted successfully: ID %d\n", post.ID)
    c.JSON(http.StatusOK, gin.H{"data": true})
}

