package db

import (
	"log"
	"market/models"
)

// GetProducts "Сканирует и возвращает значение продукта"
func GetProducts() (products []models.Product, err error) {
	rows, err := pgPool.Query("getProducts")
	if err != nil {
		log.Println("GetProducts()", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Cost,
			&product.Picture,
		)
		if err != nil {
			log.Println(err)
			return

		}
		products = append(products, product)
	}
	return
}
