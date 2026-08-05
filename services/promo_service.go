package services

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"rumah-restaurant/config"
	"rumah-restaurant/models"
	"rumah-restaurant/repositories"

	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func CreatePromo(file *multipart.FileHeader, caption string, c *gin.Context) (gin.H, error) {

	ext := filepath.Ext(file.Filename)

	filename := fmt.Sprintf(
		"promo_%d%s",
		time.Now().Unix(),
		ext,
	)

	path := "uploads/" + filename

	err := c.SaveUploadedFile(file, path)

	if err != nil {
		return nil, err
	}

	err = repositories.CreatePromo(path, caption)

	if err != nil {
		return nil, err
	}

	return gin.H{

		"message": "Promo berhasil ditambahkan",

		"image": config.AppURL() + "/" + path,
	}, nil
}

func GetAllPromo() (interface{}, error) {

	promos, err := repositories.GetAllPromo()

	if err != nil {
		return nil, err
	}

	for i := range promos {

		promos[i].Img =
			config.AppURL() +
				"/" +
				promos[i].Img
	}

	return promos, nil
}

func GetPromoByID(id string) (models.Promo, error) {

	promo, err := repositories.GetPromoByID(id)

	if err != nil {
		return promo, err
	}

	promo.Img = config.AppURL() + "/" + promo.Img

	return promo, nil
}

func UpdatePromo(id string, file *multipart.FileHeader, caption string, c *gin.Context) error {

	// ambil data lama
	promo, err := repositories.GetPromoByID(id)

	if err != nil {
		return err
	}

	imagePath := promo.Img

	// jika upload gambar baru
	if file != nil {

		// hapus gambar lama
		if _, err := os.Stat(imagePath); err == nil {
			os.Remove(imagePath)
		}

		// ekstensi
		ext := filepath.Ext(file.Filename)

		// nama file baru
		filename := fmt.Sprintf("promo_%d%s", time.Now().Unix(), ext)

		imagePath = "uploads/" + filename

		// simpan file
		err = c.SaveUploadedFile(file, imagePath)

		if err != nil {
			return err
		}
	}

	// update database
	err = repositories.UpdatePromo(
		id,
		imagePath,
		caption,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeletePromo(id string) error {

	// ambil data lama
	promo, err := repositories.GetPromoByID(id)

	if err != nil {
		return err
	}

	// hapus gambar
	if _, err := os.Stat(promo.Img); err == nil {
		os.Remove(promo.Img)
	}

	// hapus database
	err = repositories.DeletePromo(id)

	if err != nil {
		return err
	}

	return nil
}
