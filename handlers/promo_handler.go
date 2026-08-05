package handlers

import (
	"net/http"

	"rumah-restaurant/services"

	"github.com/gin-gonic/gin"
)

func CreatePromo(c *gin.Context) {

	image, err := c.FormFile("img")

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Image wajib diupload",
		})

		return
	}

	caption := c.PostForm("caption")

	result, err := services.CreatePromo(image, caption, c)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetAllPromo(c *gin.Context) {

	promos, err := services.GetAllPromo()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data",
		"data":    promos,
	})
}

func GetPromoByID(c *gin.Context) {

	id := c.Param("id")

	promo, err := services.GetPromoByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Promo tidak ditemukan",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data",
		"data":    promo,
	})
}

func UpdatePromo(c *gin.Context) {

	id := c.Param("id")

	caption := c.PostForm("caption")

	file, _ := c.FormFile("img")

	err := services.UpdatePromo(
		id,
		file,
		caption,
		c,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Promo berhasil diupdate",
	})
}

func DeletePromo(c *gin.Context) {

	id := c.Param("id")

	err := services.DeletePromo(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Promo berhasil dihapus",
	})
}
