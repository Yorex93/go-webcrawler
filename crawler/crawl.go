package crawler

// A single crawl task
func crawl(site string) []string {
	resp, err := fetchHtml(site)
	if err != nil {
		return []string{}
	}
	siteUrls := parseHtml(resp)
	return siteUrls
}
