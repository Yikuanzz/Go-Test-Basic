package model

type Item struct {
	ID          int
	Name        string `gorm:"type:varchar(100);unique_index"`
	Description string
}
