package models

import "time"

type UserModel struct {
	ID             int    `json:"ID" gorm:"unique;primaryKey;autoIncrement"`
	Name           string `json:"Name"`
	Surname        string `json:"Surname"`
	Address        string `json:"Address"`
	PassportNumber string `json:"Passport Number"`
}

type TaskModel struct {
	ID          int           `json:"ID" gorm:"unique;primaryKey;autoIncrement"`
	Task        string        `json:"Task"`
	UserID      int           `json:"User ID"`
	StartedAt   time.Time     `json:"Started At" gorm:"type:timestamp without time zone"`
	FinishedAt  time.Time     `json:"Finished At" gorm:"type:timestamp without time zone"`
	SummaryTime time.Duration `json:"Summary Time"`
}
