package db

import (
	"anime-bot/parser"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// DeleteAll - delete all records from table
func DeleteAll(tableName string) {
	_, err := Database.Exec(fmt.Sprintf("delete from %s", tableName))
	if err != nil {
		log.Panic(err)
	}
	time.Sleep(500 * time.Millisecond) // half second delay
}

// InsertAnimeList - insertion of anime list to table
func InsertAnimeList(animeList []parser.Anime, tableName string) error {
	insertQuery := fmt.Sprintf("insert into %s (title, url, episodes) values ", tableName)
	for _, anime := range animeList {
		insertQuery += fmt.Sprintf("('%s', '%s', '%s'),", anime.Title, anime.Url, anime.Episode)
	}
	_, err := Database.Exec(insertQuery[:len(insertQuery)-1])
	if err != nil {
		return err
	}
	return nil
}

// ParseRowsToAnimeList - parsing DB rows into Anime objects
func ParseRowsToAnimeList(tableName string) []parser.Anime {
	rows := getAllRows(tableName)
	defer rows.Close()

	var animeList []parser.Anime

	for rows.Next() {
		anime := parser.Anime{}
		err := rows.Scan(&anime.Id, &anime.Title, &anime.Url, &anime.Episode)
		if err != nil {
			log.Panic(err)
			continue
		}
		animeList = append(animeList, anime)
	}
	return animeList
}

func getAllRows(tableName string) *sql.Rows {
	rows, err := Database.Query(fmt.Sprintf("select * from %s", tableName))
	if err != nil {
		log.Panic(err)
	}

	return rows
}
