package crawler

type SiteMap struct {
	baseUrl string
	siteUrls []string
}

func BuildSiteMap(siteUrl string) SiteMap {
	return SiteMap{}
}
