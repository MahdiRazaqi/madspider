package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/sirupsen/logrus"
)

var authData auth

const (
	authURL = "https://trends.google.com/trends/api/explore"
	dataURL = "https://trends.google.com/trends/api/widgetdata/relatedsearches"
)

func getToken(country string) {
	resp, err := request(authURL, queryParams{
		"hl":  "en-US",
		"tz":  "-270",
		"req": url.PathEscape(fmt.Sprintf(`{"comparisonItem":[{"geo":"%v","time":"today %v"}],"category":0,"property":""}`, country, "12-m")),
	})
	if err != nil {
		logrus.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	if err := json.Unmarshal(body[4:], &authData); err != nil {
		logrus.Error(err)
	}
}

func getCountryTrends(country string) {
	response := map[string]trendsResp{}
	for _, r := range authData.Widgets {
		resp, err := request(dataURL, queryParams{
			"hl":    "en-US",
			"tz":    "-270",
			"req":   urlEncoding(country, r.Request.KeywordType),
			"token": r.Token,
		})
		if err != nil {
			logrus.Error(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
		}

		var temp trendsResp
		if err := json.Unmarshal(body[5:], &temp); err != nil {
			logrus.Error(err)
		}

		response[r.Request.KeywordType] = temp
	}

	var cTrends countryTrends
	for k, v := range response {
		tc := &cTrends.Query
		if k == "ENTITY" {
			tc = &cTrends.Entity
		}

		for i, r := range v.Default.RankedList {
			tr := &tc.Top
			if i == 1 {
				tr = &tc.Rising
			}

			for _, trend := range r.RankedKeyword {
				if trend.Query == "" {
					*tr = append(*tr, trend.Topic.Title+" - "+trend.Topic.Type)
				} else {
					*tr = append(*tr, trend.Query)
				}
			}
		}
	}

	jsonData, err := json.Marshal(cTrends)
	if err != nil {
		logrus.Error(err)
	}

	if err := ioutil.WriteFile("exports/"+countryList[country]+".json", jsonData, 0644); err != nil {
		logrus.Error(err)
	}
}

func urlEncoding(country, keywordType string) string {
	return "{%22restriction%22:{%22geo%22:{%22country%22:%22" + country + "%22},%22time%22:%222019-07-10%202020-07-10%22,%22originalTimeRangeForExploreUrl%22:%22today%2012-m%22},%22keywordType%22:%22" + keywordType + "%22,%22metric%22:[%22TOP%22,%22RISING%22],%22trendinessSettings%22:{%22compareTime%22:%222018-07-08%202019-07-09%22},%22requestOptions%22:{%22property%22:%22%22,%22backend%22:%22IZG%22,%22category%22:0},%22language%22:%22en%22,%22userCountryCode%22:%22US%22}"
}
