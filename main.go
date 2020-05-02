package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	phoneinfoga "gopkg.in/sundowndev/phoneinfoga.v2/pkg/scanners"
)

func processDorks(ctx context.Context, title string, links []*phoneinfoga.GoogleSearchDork) error {
	var results string

	color.New(color.FgWhite).Printf("[-] %s\n", title)

	for _, link := range links {
		err := chromedp.Run(ctx,
			chromedp.Navigate(link.URL),
			// chromedp.WaitVisible("*"),
			chromedp.OuterHTML("body", &results),
		)
		if err != nil {
			return err
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(results))
		if err != nil {
			return err
		}

		if _, exist := doc.Find("div#recaptcha").Attr("class"); exist {
			log.Fatal("Chromium is facing Google catpcha.")
		}

		doc.Find("div#search div.g").Each(func(_ int, el *goquery.Selection) {
			str, _ := el.Find("a").Attr("href")
			color.New(color.FgGreen).Printf("[+] %s\n", str)
		})
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		color.New(color.FgRed).Printf("[!] %s\n", "Please specify a phone number to scan.")
		os.Exit(1)
	}

	input := os.Args[1]
	number, err := phoneinfoga.LocalScan(input)

	if err != nil {
		log.Fatal(err)
	}

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	links := phoneinfoga.GoogleSearchScan(number)

	err = processDorks(ctx, "General sources", links.General)
	if err != nil {
		log.Fatal(err)
	}

	err = processDorks(ctx, "Individual sources", links.Individuals)
	if err != nil {
		log.Fatal(err)
	}

	err = processDorks(ctx, "Reputation footprints", links.Reputation)
	if err != nil {
		log.Fatal(err)
	}

	err = processDorks(ctx, "Social medias", links.SocialMedia)
	if err != nil {
		log.Fatal(err)
	}

	err = processDorks(ctx, "Disposable number providers", links.DisposableProviders)
	if err != nil {
		log.Fatal(err)
	}
}
