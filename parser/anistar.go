package parser

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

const (
	AnistarURL     = "https://anistar.org/anime/"
	AnistarTestURL = "https://localhost:44381/elastic/anistar"
)

var animeList []Anime
var episodes []string

func ParseAnistar() []Anime {
	animeList = findAllAnime(AnistarURL)[2:]
	newAnimeList := make([]Anime, 0, len(animeList))
	counter := 0
	for _, anime := range animeList {
		var episode string
		if counter >= len(episodes) {
			episode = "Нет"
		} else {
			episode = episodes[counter]
		}
		newAnimeList = append(newAnimeList, Anime{
			Title:   anime.Title,
			Episode: episode,
			Url:     anime.Url,
		})
		counter++
	}
	defer cleanUp()
	return newAnimeList
}

func findAllAnime(url string) []Anime {
	c := colly.NewCollector()

	c.OnHTML("div.title_left > a", func(e *colly.HTMLElement) {
		title := e.Text
		url := e.Attr("href")
		if len(title) > 0 {
			anime := Anime{Title: title, Url: url}
			animeList = append(animeList, anime)
		}
	})

	c.OnHTML("p.reason", func(e *colly.HTMLElement) {
		episode := e.Text
		preparedEpisode := parseEpisode(episode)
		if len(preparedEpisode) > 0 {
			episodes = append(episodes, preparedEpisode)
		}
	})

	c.Visit(url)
	return animeList
}

func parseEpisode(episode string) string {
	regexp, err := regexp.Compile("Добавлен.*?\\s+(.*?)\\s+с\\sрусской\\s")
	if err != nil {
		log.Panic(err)
	}
	matches := regexp.FindAllStringSubmatch(episode, -1)
	if len(matches) > 0 {
		return matches[0][1]
	}
	return ""
}

func cleanUp() {
	animeList = make([]Anime, 0, 10)
	episodes = make([]string, 0, 10)
}
