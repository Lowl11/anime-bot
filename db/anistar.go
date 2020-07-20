package db

import (
	"anime-bot/parser"

	// Library for using PostgreSQL
	_ "github.com/lib/pq"
)

// AnistarTableName - name of table
const AnistarTableName = "Anistar"

// CompareAnimeList - comparing anime lists from site and database
func CompareAnimeList(animeList []parser.Anime) []parser.Anime {
	defer Database.Close()
	animeListFromDb := ParseRowsToAnimeList(AnistarTableName)

	comparedAnimeMap := make(map[string]*parser.Anime)
	var comparedAnimeList []parser.Anime

	// anime list in past
	for _, anime := range animeListFromDb {
		comparedAnimeMap[anime.Title] = &anime
	}

	for _, anime := range animeList {
		if _, ok := comparedAnimeMap[anime.Title]; !ok {
			comparedAnimeList = append(comparedAnimeList, anime)
		}
	}

	if len(comparedAnimeList) > 0 {
		DeleteAll(AnistarTableName)
		InsertAnimeList(animeList, AnistarTableName)
	}

	return comparedAnimeList
}
