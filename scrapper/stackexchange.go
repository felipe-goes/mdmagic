package scrapper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func formatStackExchangeHtml(html string) string {
	formattedHtml := html

	// remove <p>
	formattedHtml = strings.Replace(formattedHtml, "<p>", "", -1)
	formattedHtml = strings.Replace(formattedHtml, "</p>", "", -1)

	// replace <pre><code> with ```
	formattedHtml = strings.Replace(formattedHtml, "<pre><code>", "```\n", -1)
	formattedHtml = strings.Replace(formattedHtml, "</code></pre>", "\n```", -1)

	// replace <code> with `
	formattedHtml = strings.Replace(formattedHtml, "<code>", "`", -1)
	formattedHtml = strings.Replace(formattedHtml, "</code>", "`", -1)

	// replace special characters
	formattedHtml = strings.Replace(formattedHtml, "&#39;", "'", -1)
	formattedHtml = strings.Replace(formattedHtml, "&#34;", "\"", -1)

	return formattedHtml
}

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
				fmt.Printf("%s\n", formatStackExchangeHtml(html))
				countAnswers++
			} else {
				html, _ := h.DOM.Html()
				fmt.Printf("# Answer %d\n", countAnswers)
				fmt.Printf("%s\n", formatStackExchangeHtml(html))
				countAnswers++
			}
		})

	// start scraping
	c.Visit(url)
}
