package models

type AddCockroachData struct {
	Amount uint32 `json:"amount"`
}

type IdCockroachData struct {
	// Id uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	Id uint32 `json:"id"`
}
