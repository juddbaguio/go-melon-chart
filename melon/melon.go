package melon

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
)

type Client struct {
	c *colly.Collector
}

func New() *Client {
	return &Client{
		c: colly.NewCollector(),
	}
}

func (melon *Client) Top100() *TopSongsSummary {
	c := melon.c
	var songDataList SongDataList
	year, month, day := time.Now().Date()
	var hour int = time.Now().Hour()

	c.OnHTML(`tr[class="lst50"]`, func(e *colly.HTMLElement) {
		songId, _ := strconv.Atoi(e.Attr("data-song-no"))
		rank, _ := strconv.Atoi(e.ChildText(`span[class="rank "]`))
		albumImg := e.ChildAttr(`img[width="60"]`, "src")
		songTitle := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank01"] a`, "title")
		songArtist := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank02"] a`, "title")
		album := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank03"] a`, "title")

		songDataList = append(songDataList, SongData{
			Id:       songId,
			Rank:     rank,
			AlbumImg: albumImg,
			Title:    songTitle,
			Artist:   songArtist,
			Album:    album,
		})
	})

	c.OnHTML(`tr[class="lst100"]`, func(e *colly.HTMLElement) {
		songId, _ := strconv.Atoi(e.Attr("data-song-no"))
		rank, _ := strconv.Atoi(e.ChildText(`span[class="rank "]`))
		albumImg := e.ChildAttr(`img[width="60"]`, "src")
		songTitle := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank01"] a`, "title")
		songArtist := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank02"] a`, "title")
		album := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank03"] a`, "title")

		songDataList = append(songDataList, SongData{
			Id:       songId,
			Rank:     rank,
			AlbumImg: albumImg,
			Title:    songTitle,
			Artist:   songArtist,
			Album:    album,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.melon.com/chart/")

	return &TopSongsSummary{
		Summary: TOP_100,
		Date:    fmt.Sprintf("%v-%v-%v", year, month, day),
		Time:    fmt.Sprintf("%v:00", hour),
		Records: songDataList,
	}
}

func (melon *Client) TopToday() *TopSongsSummary {
	c := melon.c
	var songDataList SongDataList
	year, month, day := time.Now().Add(-24 * time.Hour).Date()
	var hour int = time.Now().Hour()

	c.OnHTML(`tr[class="lst50"]`, func(e *colly.HTMLElement) {
		songId, _ := strconv.Atoi(e.Attr("data-song-no"))
		rank, _ := strconv.Atoi(e.ChildText(`span[class="rank "]`))
		albumImg := e.ChildAttr(`img[width="60"]`, "src")
		songTitle := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank01"] a`, "title")
		songArtist := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank02"] a`, "title")
		album := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank03"] a`, "title")

		songDataList = append(songDataList, SongData{
			Id:       songId,
			Rank:     rank,
			AlbumImg: albumImg,
			Title:    songTitle,
			Artist:   songArtist,
			Album:    album,
		})
	})

	c.OnHTML(`tr[class="lst100"]`, func(e *colly.HTMLElement) {
		songId, _ := strconv.Atoi(e.Attr("data-song-no"))
		rank, _ := strconv.Atoi(e.ChildText(`span[class="rank "]`))
		albumImg := e.ChildAttr(`img[width="60"]`, "src")
		songTitle := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank01"] a`, "title")
		songArtist := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank02"] a`, "title")
		album := e.ChildAttr(`div[class="wrap_song_info"] div[class="ellipsis rank03"] a`, "title")

		songDataList = append(songDataList, SongData{
			Id:       songId,
			Rank:     rank,
			AlbumImg: albumImg,
			Title:    songTitle,
			Artist:   songArtist,
			Album:    album,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.melon.com/chart/day/index.htm")

	return &TopSongsSummary{
		Summary: DAILY,
		Date:    fmt.Sprintf("%v-%v-%v", year, month, day),
		Time:    fmt.Sprintf("%v:00", hour),
		Records: songDataList,
	}
}
