/*
 * Онлайн библиотека песен
 *
 * тестовое задание
 *
 * API version: 0.0.1
 * Contact: apa.birds@mail.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package server

type Song struct {
	Id       string `json:"id,omitempty"`
	SongName string `json:"song_name,omitempty"`
	GroupId  string `json:"group_id,omitempty"`
	Text     string `json:"text,omitempty"`
	Link     string `json:"link,omitempty"`
	Date     string `json:"date,omitempty"`
}

type InfoAddSong struct {
	Group string `json:"group,omitempty"`
	Song  string `json:"song,omitempty"`
}

type InfoSong struct {
	ReleaseDate string `json:"releaseDate,omitempty"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty"`
}

type VerseSong struct {
	Id     string `json:"id,omitempty"`
	Verse  string `json:"text,omitempty"`
	Verses int    `json:"verses,omitempty"`
}
