package db

import "gitlab.com/idoko/bucketeer/models"

func (db Database) GetAllItems() (*models.ItemList, error) {
	list := &models.ItemList{}

	rows, err := db.Conn.Query("SELECT * FROM items ORDER BY ID DESC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

func (db Database) AddItem(item models.Item) error {
	var id int
	query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id`
	err := db.Conn.QueryRow(query, item.Name, item.Description).Scan(&id)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	item.ID = id
	return nil
}