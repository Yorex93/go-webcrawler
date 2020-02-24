package crawler

// A single crawl task
func crawl(site string) []string {
	var links []string
	resp, err := fetchHtml(site)
	if err != nil {
		return links
	}
	links = parseHtml(resp)
	return links
}
