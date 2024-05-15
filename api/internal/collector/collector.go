package collector

import (
	"log"

	"github.com/mmcdole/gofeed"
)

func CollectContent(sites []string) []string {
	var content []string
	fp := gofeed.NewParser()

	for _, site := range sites {
		feed, err := fp.ParseURL(site)
		if err != nil {
			log.Printf("Failed to parse feed: %v", err)
			continue
		}

		for _, item := range feed.Items {
			content = append(content, item.Title+": "+item.Link)
		}
	}

	return content
}
