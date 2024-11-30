package postgres

func (c *ConnectDB) GetIdGroup(name string) (string, error) {
	query := `
		with s as (
			SELECT id FROM groups
			WHERE name = $1 
		), i as (
			INSERT INTO groups (name)
			SELECT $1
			WHERE NOT EXISTS (SELECT 1 FROM s)
			RETURNING id
		)
		SELECT id
		FROM i
		UNION ALL
		SELECT id
		FROM s
	`

	var id string
	err := c.db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *ConnectDB) GetGroup(id string) (string, error) {
	query := `SELECT name FROM groups WHERE id = $1;`

	var name string

	err := c.db.QueryRow(query, id).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}
