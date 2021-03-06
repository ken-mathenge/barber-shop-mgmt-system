// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Customer struct {
	ID      string    `json:"id"`
	TimeIn  time.Time `json:"timeIn"`
	TimeOut time.Time `json:"timeOut"`
	Service []string  `json:"service"`
	Pay     int       `json:"pay"`
}

type Employee struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Gender   string   `json:"gender"`
	IDNumber int      `json:"idNumber"`
	Position []string `json:"position"`
}

type Service struct {
	Service []string `json:"service"`
}
