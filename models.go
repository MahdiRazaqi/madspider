package main

type auth struct {
	Widgets []struct {
		Request struct {
			KeywordType string `json:"keywordType"`
		} `json:"request"`
		ID    string `json:"id"`
		Token string `json:"token"`
	} `json:"widgets"`
}

type trendsResp struct {
	Default struct {
		RankedList []struct {
			RankedKeyword []struct {
				Topic struct {
					Mid   string `json:"mid"`
					Title string `json:"title"`
					Type  string `json:"type"`
				} `json:"topic"`
				Value          int    `json:"value"`
				FormattedValue string `json:"formattedValue"`
				HasData        bool   `json:"hasData"`
				Link           string `json:"link"`
				Query          string `json:"query"`
			} `json:"rankedKeyword"`
		} `json:"rankedList"`
	} `json:"default"`
}

type countryTrends struct {
	Query struct {
		Top    []string `json:"top"`
		Rising []string `json:"rising"`
	} `json:"query"`
	Entity struct {
		Top    []string `json:"top"`
		Rising []string `json:"rising"`
	} `json:"entity"`
}
