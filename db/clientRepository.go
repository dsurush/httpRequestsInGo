package db

import (
	"fmt"
	"log"
	"market/models"

	"github.com/google/uuid"
)

// DeleteClient "dasdasdas"
func DeleteClient(id string) (err error) {

	row := pgPool.QueryRow(`deleteClient`, id)
	if err != nil {
		log.Println("DeleteClient()", err)
		return err
	}

	fmt.Println(row)
	return nil
}

// GetClients "Функция для передачи клиентов"
func GetClients() (clients []models.Client, err error) {
	rows, err := pgPool.Query("getClients")
	if err != nil {
		log.Println("GetClients", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		client := models.Client{}
		err = rows.Scan(
			&client.ID,
			&client.FirstName,
			&client.SurName,
			&client.Address,
			&client.Email,
			&client.PhoneNumbers,
		)
		if err != nil {
			log.Println(err)
			return
		}
		clients = append(clients, client)
	}
	return
}

// AddClient "Сканирует и возвращает значение client"
func AddClient(newclient models.CreateNewClient) (models.CreateNewClient, error) {
	id, _ := uuid.NewUUID()

	tx, _ := pgPool.Begin()

	ct, err := tx.Exec("addClient",
		id,
		newclient.FirstName,
		newclient.SurName,
		newclient.Address,
		newclient.Email,
		newclient.PhoneNumbers,
	)
	// models.CreateNewClient "dasddasdasdasd"
	client := models.CreateNewClient{
		newclient.FirstName,
		newclient.SurName,
		newclient.Address,
		newclient.Email,
		newclient.PhoneNumbers,
	}

	/////Я вообще хз что ето
	if err != nil {
		log.Println("AddClient()", client, err)
		tx.Rollback()
		//	return nil, nil
	}

	if n := ct.RowsAffected(); n < 1 {
		log.Println("AddClient()", client, err)
	}
	tx.Commit()

	return client, nil
}
