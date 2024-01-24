package main

type Type struct {
	TypeId   int    `json:"typeid"`
	TypeName string `json:"name"`
}

func GetTypes() ([]Type, error) {
	rows, err := db.Query("SELECT * FROM `TYPES`")
	if err != nil {
		return nil, err
	}
	data := &Type{}
	res := make([]Type, 0)
	for rows.Next() {
		err = rows.Scan(&data.TypeId, &data.TypeName)
		if err != nil {
			return nil, err
		}
		res = append(res, *data)
	}
	return res, nil
}
