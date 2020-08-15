package constants

var (
	DocumentTitle   string = "%s | bozdag.dev"
	UrlRegex        string = `(http[s]?:\/\/)?([^\/\s]+\/)(.*)`
	HtmlRegex       string = `([a-zA-Z0-9\s_\\.\-\(\):])+(.html)$`
	MediumCdnRegex  string = `(https?:\/\/(.+?\.)?cdn-images-1.medium\.com(\/[A-Za-z0-9\-\._~:\/\?#\[\]@!$&'\(\)\*\+,;\=]*)?)`
	MediumHrefRegex string = `>(https?:\/\/(.+?\.)?medium\.com(\/[A-Za-z0-9\-\._~:\/\?#\[\]@!$&'\(\)\*\+,;\=]*)href?)`
)

var (
	PostsCacheKey string = "PostsCacheKey"
)
