package handlers

import (
	"net/http"

	"tes-database-pq/config"
	"tes-database-pq/models"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @Summary Mendapatkan semua users
// @Description Mengambil daftar semua user dari database
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.GetDB().Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

// GetUser godoc
// @Summary Mendapatkan user berdasarkan ID
// @Description Mengambil detail satu user berdasarkan ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := config.GetDB().First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

// CreateUser godoc
// @Summary Membuat user baru
// @Description Menambahkan user baru ke database
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "Data User"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var input models.CreateUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	user := models.User{
		Name:  input.Name,
		Email: input.Email,
	}

	result := config.GetDB().Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User berhasil dibuat",
		"data":    user,
	})
}

// UpdateUser godoc
// @Summary Update user
// @Description Mengupdate data user berdasarkan ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "Data User"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := config.GetDB().First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User tidak ditemukan",
		})
		return
	}

	var input models.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	config.GetDB().Model(&user).Updates(models.User{
		Name:  input.Name,
		Email: input.Email,
	})

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User berhasil diupdate",
		"data":    user,
	})
}

// DeleteUser godoc
// @Summary Hapus user
// @Description Menghapus user berdasarkan ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	result := config.GetDB().First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "User tidak ditemukan",
		})
		return
	}

	config.GetDB().Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User berhasil dihapus",
	})
}
