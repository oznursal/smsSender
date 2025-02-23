package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"smsSender/model"
)

func connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbInfo, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer dbInfo.Close()

	err = dbInfo.Ping()
	if err != nil {
		return nil, err
	}
	return dbInfo, nil
}

func GetUnsentMessages() ([]model.MessageDto, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbInfo, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer dbInfo.Close()
	count := 2
	rows, err := dbInfo.Query("SELECT * FROM insider.messages WHERE sending_status=false AND LENGTH(replace(content,' ','')) BETWEEN 1 AND 160 ORDER BY ID ASC LIMIT $1", count)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var messages []model.MessageDto

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var message model.MessageDto
		if err := rows.Scan(&message.ID, &message.Content, &message.PhoneNumber,
			&message.SendingStatus); err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return messages, err
	}
	return messages, nil
}

func GetSentMessages() ([]model.Message, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbInfo, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer dbInfo.Close()
	rows, err := dbInfo.Query("SELECT * FROM insider.messages WHERE sending_status=true AND LENGTH(replace(content,' ','')) BETWEEN 1 AND 160 ORDER BY ID ASC")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var messages []model.Message

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var message model.Message
		if err := rows.Scan(&message.ID, &message.Content, &message.PhoneNumber,
			&message.SendingStatus); err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return messages, err
	}
	return messages, nil
}

func UpdateSentMessages(dto model.MessageDto) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbInfo, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer dbInfo.Close()

	_, err = dbInfo.Exec("UPDATE insider.messages SET sending_status=true WHERE id=$1", dto.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
