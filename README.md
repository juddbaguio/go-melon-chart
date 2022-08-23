# Melon Chart API - Go
## I love K-Pop, do I need to say more?

Test project for using [go-colly](https://github.com/gocolly/colly) web-scraping framework.

# Installation
```bash
go get -u github.com/juddbaguio/go-melon-chart
```

# Usage
```go
func main() {
    melon := melon.New()
    top_100 := melon.Top100()
}
```

## Sample Result
```json
{
	"summary": "TOP 100",
	"date": "2022-August-23",
	"time": "23:00",
	"records": [
		{
			"id": 35454425,
			"rank": 1,
			"albumImg": "https://cdnimg.melon.co.kr/cm2/album/images/110/11/565/11011565_20220801102637_500.jpg/melon/resize/120/quality/80/optimize",
			"title": "Attention 재생",
			"artist": "NewJeans - 페이지 이동",
			"album": "NewJeans 1st EP 'New Jeans' - 페이지 이동"
		},
        // rest of the songs here..
	]
}
```