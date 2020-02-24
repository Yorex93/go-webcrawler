package crawler

type SiteMap struct {
	baseUrl string
	siteUrls []string
}

func BuildSiteMap(siteUrl string) SiteMap {
	var result SiteMap

	links := crawl(siteUrl)
	result.siteUrls = links
	result.baseUrl = siteUrl

	return result
}
