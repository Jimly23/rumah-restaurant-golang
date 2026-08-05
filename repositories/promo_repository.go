package repositories

import (
	"rumah-restaurant/config"
	"rumah-restaurant/models"
)

func CreatePromo(img string, caption string) error {

	query := "INSERT INTO promo(img, caption) VALUES(?, ?)"

	_, err := config.DB.Exec(query, img, caption)

	return err
}

func GetAllPromo() ([]models.Promo, error) {

	query := "SELECT id, img, caption FROM promo"

	rows, err := config.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var promos []models.Promo

	for rows.Next() {

		var promo models.Promo

		err := rows.Scan(
			&promo.ID,
			&promo.Img,
			&promo.Caption,
		)

		if err != nil {
			return nil, err
		}

		promos = append(promos, promo)
	}

	return promos, nil
}

func GetPromoByID(id string) (models.Promo, error) {

	var promo models.Promo

	query := "SELECT id, img, caption FROM promo WHERE id=?"

	err := config.DB.QueryRow(query, id).Scan(
		&promo.ID,
		&promo.Img,
		&promo.Caption,
	)

	return promo, err
}

func UpdatePromo(id string, img string, caption string) error {

	query := `
	UPDATE promo
	SET img=?, caption=?
	WHERE id=?
	`

	_, err := config.DB.Exec(
		query,
		img,
		caption,
		id,
	)

	return err
}

func DeletePromo(id string) error {

	query := "DELETE FROM promo WHERE id=?"

	_, err := config.DB.Exec(query, id)

	return err
}
