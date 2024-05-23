package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
)

func CppByExample(url string) {
	// instantiate default collector
	c := colly.NewCollector()

	// find question title
	c.OnHTML("article",
		func(h *colly.HTMLElement) {
			html, _ := h.DOM.Html()
			fmt.Printf("%s\n", FormatCppByExample(html))
		})

	// start scraping
	c.Visit(url)
}
