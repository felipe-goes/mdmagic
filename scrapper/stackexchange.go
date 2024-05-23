package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
)

func StackExchange(url string) {
	// instantiate default collector
	c := colly.NewCollector()

	// find question title
	foundQT := false
	c.OnHTML("a[class=question-hyperlink]",
		func(h *colly.HTMLElement) {
			if !foundQT {
				fmt.Printf("# %s", h.Text)
			}
			foundQT = true
		})

	countAnswers := 0
	// get question and answers
	c.OnHTML("div[class=\"s-prose js-post-body\"][itemprop=text]",
		func(h *colly.HTMLElement) {
			if countAnswers == 0 {
				html, _ := h.DOM.Html()
				fmt.Printf("%s\n", FormatStackExchange(html))
				countAnswers++
			} else {
				html, _ := h.DOM.Html()
				fmt.Printf("# Answer %d\n", countAnswers)
				fmt.Printf("%s\n", FormatStackExchange(html))
				countAnswers++
			}
		})

	// start scraping
	c.Visit(url)
}
