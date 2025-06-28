package controllers

import (
	"net/http"
	"poxa_rafa/database"
	"poxa_rafa/models"

	"github.com/gin-gonic/gin"
)

// GetPosts godoc
// @Summary Lista todos os posts
// @Description Retorna todos os posts salvos
// @Tags Posts
// @Produce json
// @Success 200 {array} models.Post
// @Router /posts [get]
func GetPosts(c *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

// GetPostByID godoc
// @Summary Retorna um post por ID
// @Tags Posts
// @Produce json
// @Param id path int true "ID do post"
// @Success 200 {object} models.Post
// @Failure 404 {object} map[string]string
// @Router /posts/{id} [get]
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// CreatePost godoc
// @Summary Cria um novo post
// @Tags Posts
// @Accept json
// @Produce json
// @Param post body models.Post true "Dados do post"
// @Success 201 {object} models.Post
// @Failure 400 {object} map[string]string
// @Router /posts [post]
// @Security BearerAuth
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&post)
	c.JSON(http.StatusCreated, post)
}

// UpdatePost godoc
// @Summary Atualiza um post existente
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path int true "ID do post"
// @Param post body models.Post true "Novos dados do post"
// @Success 200 {object} models.Post
// @Failure 400,404 {object} map[string]string
// @Router /posts/{id} [put]
// @Security BearerAuth
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var updatedData models.Post
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Title = updatedData.Title
	post.Content = updatedData.Content
	post.Tags = updatedData.Tags
	post.CoverImage = updatedData.CoverImage

	database.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

// DeletePost godoc
// @Summary Deleta um post
// @Tags Posts
// @Produce json
// @Param id path int true "ID do post"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /posts/{id} [delete]
// @Security BearerAuth
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	database.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
