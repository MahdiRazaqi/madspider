package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type queryParams map[string]string

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
