package postgres

import "fmt"

type SongPostgres struct {
	Id       string
	SongName string
	GroupId  string
	Text     string
	Link     string
	Date     string
}

type FilterSongList struct {
	SongName  string
	GroupId   string
	GroupName string
	Text      string
	Link      string
	DateStart string
	DateEnd   string
}

func (c *ConnectDB) AddSong(song, group, text, link, releaseDate string) (string, error) {
	groupId, err := c.GetIdGroup(group)
	if err != nil {
		return "", err
	}

	query := `INSERT INTO songs (song, group_id, text, link, release_date) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var id string

	err = c.db.QueryRow(query, song, groupId, text, link, releaseDate).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *ConnectDB) DeleteSong(id string) (*SongPostgres, error) {
	query := `DELETE FROM songs WHERE id = $1 RETURNING *;`

	res := c.db.QueryRow(query, id)

	var song SongPostgres
	if err := res.Scan(song.Id, song.GroupId, song.Text, song.Link, song.Date); err != nil {
		return nil, err
	}

	return &song, nil
}

func (c *ConnectDB) UpdateSong(id, songName, group, text, link, releaseDate string) (*SongPostgres, error) {
	query := `UPDATE songs SET song_name = $1, group_id = $2, text = $3, link = $4, release_date = $5 WHERE id = $6 RETURNING *;`

	res := c.db.QueryRow(query, songName, group, text, link, releaseDate, id)

	var song SongPostgres
	if err := res.Scan(song.Id, song.GroupId, song.Text, song.Link, song.Date); err != nil {
		return nil, err
	}

	return &song, nil
}

func (c *ConnectDB) GetSong(id string) (*SongPostgres, error) {
	query := `SELECT id, song_name, group_id, text, link, release_date FROM songs WHERE id = $1;`

	song := SongPostgres{}

	err := c.db.QueryRow(query, id).Scan(&song.Id, &song.SongName, &song.GroupId, &song.Text, &song.Link, &song.Date)
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (c *ConnectDB) GetSongs(page, volume int, filter FilterSongList) ([]SongPostgres, error) {
	paginationMark, err := c.valueForPaginationSong(page, volume)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, song_name, group_id, text, link, release_date FROM songs 
		WHERE created_at > $1 
		`

	numFilter := 2
	values := make([]interface{}, 0)
	values = append(values, paginationMark)

	if filter.SongName != "" {
		query += fmt.Sprint(`AND song_name ilike '%' || $`, numFilter, ` || '%' `)
		values = append(values, filter.SongName)
		numFilter++
	}
	if filter.GroupId != "" {
		query += fmt.Sprintf(`AND group_id = $%d `, numFilter)
		values = append(values, filter.GroupId)
		numFilter++
	}
	if filter.GroupName != "" {
		query += fmt.Sprint(`AND group_name ilike '%' || $`, numFilter, ` || '%' `)
		values = append(values, filter.GroupName)
		numFilter++
	}
	if filter.Text != "" {
		query += fmt.Sprint(`AND text ilike '%' || $`, numFilter, ` || '%' `)
		values = append(values, filter.Text)
		numFilter++
	}
	if filter.Link != "" {
		query += fmt.Sprintf(`AND link = $%d `, numFilter)
		values = append(values, filter.Link)
		numFilter++
	}
	if filter.DateStart != "" {
		query += fmt.Sprintf(`AND release_date >= $%d `, numFilter)
		values = append(values, filter.DateStart)
		numFilter++
	}
	if filter.DateEnd != "" {
		query += fmt.Sprintf(`AND release_date <= $%d `, numFilter)
		values = append(values, filter.DateEnd)
		numFilter++
	}

	query += fmt.Sprintf(`ORDER BY created_at DESC LIMIT $%d;`, numFilter)
	values = append(values, fmt.Sprint(volume))

	rows, err := c.db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	var songs []SongPostgres
	for rows.Next() {
		var song SongPostgres
		if err := rows.Scan(&song.Id, &song.SongName, &song.GroupId, &song.Text, &song.Link, &song.Date); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (c *ConnectDB) valueForPaginationSong(page, volume int) (string, error) {
	query := `SELECT created_at FROM songs ORDER BY created_at DESC LIMIT $1;`

	rows, err := c.db.Query(query, (volume*(page-1))+1)
	if err != nil {
		return "", err
	}

	for i := 0; i < volume*(page-1); i++ {
		rows.Next()
	}

	for rows.Next() {
		var date string
		if err := rows.Scan(&date); err != nil {
			return "", err
		}
		return date, nil
	}

	return "", nil
}
