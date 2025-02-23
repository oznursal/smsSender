package model

type MessageDto struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	Content       string `json:"content"`
	PhoneNumber   string `json:"phoneNumber"`
	SendingStatus bool   `json:"sendingStatus"`
}

type Message struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	Content       string `json:"content"`
	PhoneNumber   string `json:"phoneNumber"`
	SendingStatus bool   `json:"sendingStatus"`
}

/*
CREATE TABLE IF NOT EXISTS messages (
  id integer NOT NULL DEFAULT 1,
  content varchar(160) NOT NULL,
  phone_number varchar(12) NOT NULL,
  sending_status boolean NOT NULL DEFAULT false,
  PRIMARY KEY (id)
)
*/
