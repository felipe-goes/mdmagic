package scrapper

import (
	"regexp"
	"strings"
)

func FormatCppByExample(html string) string {
	var regex *regexp.Regexp
	formattedHtml := html

	// remove <p>
	formattedHtml = strings.Replace(formattedHtml, "<p>", "", -1)
	formattedHtml = strings.Replace(formattedHtml, "</p>", "\n", -1)

	// edit <h1>
	formattedHtml = strings.Replace(formattedHtml, "<h1>", "# ", -1)
	formattedHtml = strings.Replace(formattedHtml, "</h1>", "\n", -1)

	// add \n to <h4>
	formattedHtml = strings.Replace(formattedHtml, "</h4>", "</h4>\n", -1)

	// replace <span>
	regex = regexp.MustCompile(`<span.+?>`)
	formattedHtml = regex.ReplaceAllString(formattedHtml, "")
	formattedHtml = strings.Replace(formattedHtml, "</span>", "", -1)

	// replace <pre> with ```
	formattedHtml = strings.Replace(formattedHtml, "<pre>", "```\n", -1)
	formattedHtml = strings.Replace(formattedHtml, "</pre>", "\n```", -1)

	// replace <code> with `
	formattedHtml = strings.Replace(formattedHtml, "<code>", "`", -1)
	formattedHtml = strings.Replace(formattedHtml, "</code>", "`", -1)

	// replace special characters
	formattedHtml = strings.Replace(formattedHtml, "&#39;", "'", -1)
	formattedHtml = strings.Replace(formattedHtml, "&#34;", "\"", -1)
	formattedHtml = strings.Replace(formattedHtml, "&lt;", "<", -1)
	formattedHtml = strings.Replace(formattedHtml, "&gt;", ">", -1)

	return formattedHtml
}

func FormatStackExchange(html string) string {
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
	formattedHtml = strings.Replace(formattedHtml, "&lt;", "<", -1)
	formattedHtml = strings.Replace(formattedHtml, "&gt;", ">", -1)

	return formattedHtml
}
