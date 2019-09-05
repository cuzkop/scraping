package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sclevine/agouti"
)

// BaseURL URL共通部分のベース
const BaseURL string = "https://search.yahoo.co.jp/search?p="

func main(word string) ([]string, []string, []string, []string) {
	readerURL := BaseURL + word
	capabilities := agouti.NewCapabilities()
	capabilities["phantomjs.page.settings.userAgent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36"
	capabilitiesOption := agouti.Desired(capabilities)

	driver := agouti.PhantomJS(capabilitiesOption)
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("phantomjs"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate(readerURL); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	page.Screenshot("./tmp/capture_id_" + strconv.Itoa(id) + ".jpg")

	html, err := page.HTML()
	if err != nil {
		log.Fatalf("Failed to get HTML:%v", err)
	}

	readerHTML := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(readerHTML)
	if err != nil {
		panic(err)
	}

	var title []string
	var url []string
	var visibleURL []string
	var content []string

	doc.Find("a.t").Each(func(_ int, s *goquery.Selection) {
		t, _ := s.Html()
		t = strings.NewReplacer(
			"<em>", "",
			"</em>", "",
			"amp;", "",
		).Replace(t)
		title = append(title, t)
	})

	doc.Find("a.t").Each(func(_ int, s *goquery.Selection) {
		u, _ := s.Attr("href")
		url = append(url, u)
	})

	doc.Find("div.cf > span.u").Each(func(_ int, s *goquery.Selection) {
		u, _ := s.Html()
		u = strings.NewReplacer(
			"<span class=\"ad\">広告</span>", "",
			"<b>", "",
			"</b>", "",
			"amp;", "",
		).Replace(u)
		visibleURL = append(visibleURL, u)
	})

	doc.Find("div.w > p.x").Each(func(_ int, s *goquery.Selection) {
		c, _ := s.Html()
		c = strings.NewReplacer(
			"<em>", "",
			"</em>", "",
			"amp;", "",
		).Replace(c)
		content = append(content, c)
	})

	return title, url, visibleURL, content
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
