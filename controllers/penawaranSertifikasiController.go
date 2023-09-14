package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/finaenno/sertifikasi-go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var penawaranSertifikasi []models.PenawaranSertifikasi

	models.DB.Find(&penawaranSertifikasi)
	c.JSON(http.StatusOK, gin.H{"penawaran_sertifikasi": penawaranSertifikasi})
}

func Show(c *gin.Context) {
	var penawaranSertifikasi models.PenawaranSertifikasi

	id := c.Param("id_penawaran_sertifikasi")

	if err := models.DB.First(&penawaranSertifikasi, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"penawaran_sertifikasi": penawaranSertifikasi})
}

func Create(c *gin.Context) {
	var penawaranSertifikasi models.PenawaranSertifikasi

	if err := c.ShouldBindJSON(&penawaranSertifikasi); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&penawaranSertifikasi)
	c.JSON(http.StatusOK, gin.H{"penawaran_sertifikasi": penawaranSertifikasi})
}

func Update(c *gin.Context) {
	var penawaranSertifikasi models.PenawaranSertifikasi

	id := c.Param("id_penawaran_sertifikasi")

	if err := c.ShouldBindJSON(&penawaranSertifikasi); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&penawaranSertifikasi).Where("id_penawaran_sertifikasi = ?", id).Updates(&penawaranSertifikasi).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"penawaran_sertifikasi": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var penawaranSertifikasi models.PenawaranSertifikasi

	var input struct {
		id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.id.Int64()
	if models.DB.Delete(&penawaranSertifikasi, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
