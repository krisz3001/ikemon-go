package main

type Card struct {
	CardId      string
	Name        string
	Type        string
	Hp          int
	Attack      int
	Defense     int
	Price       float64
	Description string
	Image       string
	OwnerId     string
	IsLocked    bool
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
