package main

import (
	"fmt"
	"math"
)

type Card struct {
	CardId      string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"typeid"`
	Hp          int     `json:"hp"`
	Attack      int     `json:"attack"`
	Defense     int     `json:"defense"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	OwnerId     string  `json:"ownerid"`
	IsLocked    bool    `json:"islocked"`
}

func GetAllCards() ([]Card, error) {
	rows, err := db.Query("SELECT * FROM `CARDS`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]Card, 0)

	var card Card
	for rows.Next() {
		rows.Scan(&card.CardId, &card.Name, &card.Type, &card.Hp, &card.Attack, &card.Defense, &card.Price, &card.Description, &card.Image, &card.OwnerId, &card.IsLocked)
		result = append(result, card)
	}
	return result, nil
}

func GetCardById(id string) *Card {
	row := db.QueryRow("SELECT * FROM `CARDS` WHERE CardId = ?", id)

	var card Card
	err := row.Scan(&card.CardId, &card.Name, &card.Type, &card.Hp, &card.Attack, &card.Defense, &card.Price, &card.Description, &card.Image, &card.OwnerId, &card.IsLocked)
	if err != nil {
		return nil
	}
	return &card
}

func GetNthNineByType(n int, t int) ([]Card, error) {
	var typeQuery string
	if t == 0 {
		typeQuery = "%"
	} else {
		typeQuery = fmt.Sprint(t)
	}
	rows, err := db.Query("SELECT * FROM `CARDS` WHERE TypeId LIKE ? LIMIT ?, 9", typeQuery, n*9)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]Card, 0)

	var card Card
	for rows.Next() {
		rows.Scan(&card.CardId, &card.Name, &card.Type, &card.Hp, &card.Attack, &card.Defense, &card.Price, &card.Description, &card.Image, &card.OwnerId, &card.IsLocked)
		result = append(result, card)
	}
	return result, nil
}

func CountPageOfType(t int) (int, error) {
	var typeQuery string
	if t == 0 {
		typeQuery = "%"
	} else {
		typeQuery = fmt.Sprint(t)
	}
	res := db.QueryRow("SELECT COUNT(*) FROM `CARDS` WHERE TypeId LIKE ?", typeQuery)

	var count float64

	err := res.Scan(&count)
	if err != nil {
		return 0, err
	}

	pages := math.Ceil(count / 9)

	return int(pages), nil
}
