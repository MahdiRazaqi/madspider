package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

var countryList = map[string]string{
	"AF": "Afghanistan",
	"AX": "Åland Islands",
	"AL": "Albania",
	"DZ": "Algeria",
	"AS": "American Samoa",
	"AD": "Andorra",
	"AO": "Angola",
	"AI": "Anguilla",
	"AQ": "Antarctica",
	"AG": "Antigua & Barbuda",
	"AR": "Argentina",
	"AM": "Armenia",
	"AW": "Aruba",
	"AU": "Australia",
	"AT": "Austria",
	"AZ": "Azerbaijan",
	"BS": "Bahamas",
	"BH": "Bahrain",
	"BD": "Bangladesh",
	"BB": "Barbados",
	"BY": "Belarus",
	"BE": "Belgium",
	"BZ": "Belize",
	"BJ": "Benin",
	"BM": "Bermuda",
	"BT": "Bhutan",
	"BO": "Bolivia",
	"BA": "Bosnia & Herzegovina",
	"BW": "Botswana",
	"BV": "Bouvet Island",
	"BR": "Brazil",
	"IO": "British Indian Ocean Territory",
	"VG": "British Virgin Islands",
	"BN": "Brunei",
	"BG": "Bulgaria",
	"BF": "Burkina Faso",
	"BI": "Burundi",
	"KH": "Cambodia",
	"CM": "Cameroon",
	"CA": "Canada",
	"CV": "Cape Verde",
	"BQ": "Caribbean Netherlands",
	"KY": "Cayman Islands",
	"CF": "Central African Republic",
	"TD": "Chad",
	"CL": "Chile",
	"CN": "China",
	"CX": "Christmas Island",
	"CC": "Cocos (Keeling) Islands",
	"CO": "Colombia",
	"KM": "Comoros",
	"CG": "Congo - Brazzaville",
	"CD": "Congo - Kinshasa",
	"CK": "Cook Islands",
	"CR": "Costa Rica",
	"CI": "Côte d’Ivoire",
	"HR": "Croatia",
	"CU": "Cuba",
	"CW": "Curaçao",
	"CY": "Cyprus",
	"CZ": "Czechia",
	"DK": "Denmark",
	"DJ": "Djibouti",
	"DM": "Dominica",
	"DO": "Dominican Republic",
	"EC": "Ecuador",
	"EG": "Egypt",
	"SV": "El Salvador",
	"GQ": "Equatorial Guinea",
	"ER": "Eritrea",
	"EE": "Estonia",
	"SZ": "Eswatini",
	"ET": "Ethiopia",
	"FK": "Falkland Islands (Islas Malvinas)",
	"FO": "Faroe Islands",
	"FJ": "Fiji",
	"FI": "Finland",
	"FR": "France",
	"GF": "French Guiana",
	"PF": "French Polynesia",
	"TF": "French Southern Territories",
	"GA": "Gabon",
	"GM": "Gambia",
	"GE": "Georgia",
	"DE": "Germany",
	"GH": "Ghana",
	"GI": "Gibraltar",
	"GR": "Greece",
	"GL": "Greenland",
	"GD": "Grenada",
	"GP": "Guadeloupe",
	"GU": "Guam",
	"GT": "Guatemala",
	"GG": "Guernsey",
	"GN": "Guinea",
	"GW": "Guinea-Bissau",
	"GY": "Guyana",
	"HT": "Haiti",
	"HM": "Heard & McDonald Islands",
	"HN": "Honduras",
	"HK": "Hong Kong",
	"HU": "Hungary",
	"IS": "Iceland",
	"IN": "India",
	"ID": "Indonesia",
	"IR": "Iran",
	"IQ": "Iraq",
	"IE": "Ireland",
	"IM": "Isle of Man",
	"IL": "Israel",
	"IT": "Italy",
	"JM": "Jamaica",
	"JP": "Japan",
	"JE": "Jersey",
	"JO": "Jordan",
	"KZ": "Kazakhstan",
	"KE": "Kenya",
	"KI": "Kiribati",
	"XK": "Kosovo",
	"KW": "Kuwait",
	"KG": "Kyrgyzstan",
	"LA": "Laos",
	"LV": "Latvia",
	"LB": "Lebanon",
	"LS": "Lesotho",
	"LR": "Liberia",
	"LY": "Libya",
	"LI": "Liechtenstein",
	"LT": "Lithuania",
	"LU": "Luxembourg",
	"MO": "Macao",
	"MG": "Madagascar",
	"MW": "Malawi",
	"MY": "Malaysia",
	"MV": "Maldives",
	"ML": "Mali",
	"MT": "Malta",
	"MH": "Marshall Islands",
	"MQ": "Martinique",
	"MR": "Mauritania",
	"MU": "Mauritius",
	"YT": "Mayotte",
	"MX": "Mexico",
	"FM": "Micronesia",
	"MD": "Moldova",
	"MC": "Monaco",
	"MN": "Mongolia",
	"ME": "Montenegro",
	"MS": "Montserrat",
	"MA": "Morocco",
	"MZ": "Mozambique",
	"MM": "Myanmar (Burma)",
	"NA": "Namibia",
	"NR": "Nauru",
	"NP": "Nepal",
	"NL": "Netherlands",
	"NC": "New Caledonia",
	"NZ": "New Zealand",
	"NI": "Nicaragua",
	"NE": "Niger",
	"NG": "Nigeria",
	"NU": "Niue",
	"NF": "Norfolk Island",
	"KP": "North Korea",
	"MK": "North Macedonia",
	"MP": "Northern Mariana Islands",
	"NO": "Norway",
	"OM": "Oman",
	"PK": "Pakistan",
	"PW": "Palau",
	"PS": "Palestine",
	"PA": "Panama",
	"PG": "Papua New Guinea",
	"PY": "Paraguay",
	"PE": "Peru",
	"PH": "Philippines",
	"PN": "Pitcairn Islands",
	"PL": "Poland",
	"PT": "Portugal",
	"PR": "Puerto Rico",
	"QA": "Qatar",
	"RE": "Réunion",
	"RO": "Romania",
	"RU": "Russia",
	"RW": "Rwanda",
	"WS": "Samoa",
	"SM": "San Marino",
	"ST": "São Tomé & Príncipe",
	"SA": "Saudi Arabia",
	"SN": "Senegal",
	"RS": "Serbia",
	"SC": "Seychelles",
	"SL": "Sierra Leone",
	"SG": "Singapore",
	"SX": "Sint Maarten",
	"SK": "Slovakia",
	"SI": "Slovenia",
	"SB": "Solomon Islands",
	"SO": "Somalia",
	"ZA": "South Africa",
	"GS": "South Georgia & South Sandwich Islands",
	"KR": "South Korea",
	"SS": "South Sudan",
	"ES": "Spain",
	"LK": "Sri Lanka",
	"BL": "St. Barthélemy",
	"SH": "St. Helena",
	"KN": "St. Kitts & Nevis",
	"LC": "St. Lucia",
	"MF": "St. Martin",
	"PM": "St. Pierre & Miquelon",
	"VC": "St. Vincent & Grenadines",
	"SD": "Sudan",
	"SR": "Suriname",
	"SJ": "Svalbard & Jan Mayen",
	"SE": "Sweden",
	"CH": "Switzerland",
	"SY": "Syria",
	"TW": "Taiwan",
	"TJ": "Tajikistan",
	"TZ": "Tanzania",
	"TH": "Thailand",
	"TL": "Timor-Leste",
	"TG": "Togo",
	"TK": "Tokelau",
	"TO": "Tonga",
	"TT": "Trinidad & Tobago",
	"TN": "Tunisia",
	"TR": "Turkey",
	"TM": "Turkmenistan",
	"TC": "Turks & Caicos Islands",
	"TV": "Tuvalu",
	"UM": "U.S. Outlying Islands",
	"VI": "U.S. Virgin Islands",
	"UG": "Uganda",
	"UA": "Ukraine",
	"AE": "United Arab Emirates",
	"GB": "United Kingdom",
	"US": "United States",
	"UY": "Uruguay",
	"UZ": "Uzbekistan",
	"VU": "Vanuatu",
	"VA": "Vatican City",
	"VE": "Venezuela",
	"VN": "Vietnam",
	"WF": "Wallis & Futuna",
	"EH": "Western Sahara",
	"YE": "Yemen",
	"ZM": "Zambia",
	"ZW": "Zimbabwe",
}

type queryParams map[string]string

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

const (
	authURL = "https://trends.google.com/trends/api/explore"
	dataURL = "https://trends.google.com/trends/api/widgetdata/relatedsearches"
)

var authData auth

func main() {
	for countryCode := range countryList {
		getToken(countryCode)
		getCountryTrends(countryCode)
	}
}

func urlEncoding(country, keywordType string) string {
	return "{%22restriction%22:{%22geo%22:{%22country%22:%22" + country + "%22},%22time%22:%222019-07-10%202020-07-10%22,%22originalTimeRangeForExploreUrl%22:%22today%2012-m%22},%22keywordType%22:%22" + keywordType + "%22,%22metric%22:[%22TOP%22,%22RISING%22],%22trendinessSettings%22:{%22compareTime%22:%222018-07-08%202019-07-09%22},%22requestOptions%22:{%22property%22:%22%22,%22backend%22:%22IZG%22,%22category%22:0},%22language%22:%22en%22,%22userCountryCode%22:%22US%22}"
}

func request(url string, params queryParams) (*http.Response, error) {
	q := "?"
	for k, v := range params {
		if q != "?" {
			q += "&"
		}
		q += k + "=" + v
	}

	logrus.Info(url + q)
	req, err := http.NewRequest(http.MethodGet, url+q, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:79.0) Gecko/20100101 Firefox/79.0")
	req.Header.Set("Cookie", "__utma=10102256.368765410.1593583769.1593591249.1593667128.3; __utmz=10102256.1593583769.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmb=10102256.2.10.1593667128; __utmc=10102256; __utmt=1; 1P_JAR=2020-7-2-5; NID=204=rrmu9xNHM2XfD6oNkJV4EeN_U68bNj2wFZyxDOBMPXJA1eFULG808qnKW7q8Uw0OX461o8_sZmtdXQk7xceRLqSzrqK3prcKY9clylMa-oDmWXB0hPDE2-XYDOjxlXaX1GCoIGl_LqmkVtxTkkBamCJP6q9YGwaL3QIduTyv4xs; ANID=AHWqTUk69XZN9RFdqwZysGAqg8LEIZBribJzxCE1h04EpwNxWd2wq0Efp6QXOHGV; CONSENT=WP.28860e.288703")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

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
