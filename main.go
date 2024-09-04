package main

import (
	"fmt"
	"go-web-scrap/model"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	url := 	 "https://www.logammulia.com/id/harga-emas-hari-ini";
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })
    collector.OnResponse(func(r *colly.Response) {
        fmt.Println("Success Visited: Behold, The results: ....", r.Request.URL)
    })
    collector.OnError(func(r *colly.Response, e error) {
        fmt.Println("Error During Site Visit:", e)
    })

	collector.OnHTML("div.grid-child h2.ngc-title:first-of-type", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	collector.OnHTML("div.grid-child table:first-of-type", func(e *colly.HTMLElement) {
		
		gold := []model.Gold{};

		breakLoop := false;

		now := time.Now();

		today := now.Format("2006-01-02")

		e.ForEach("tr", func(i int, h *colly.HTMLElement) {

			
			if breakLoop {
				return;
			}

			firstChild := h.DOM.Children().First()
			
			if goquery.NodeName(firstChild) == "td" {
				row := model.Gold{};
				row.Date = today
				counter := 0;
				h.ForEach("td", func(j int, x *colly.HTMLElement) {
					if counter == 0 {
						row.Weight = x.Text
						counter++;
					} else if counter == 1 {
					row.Price = x.Text;
					counter++
					} else {
						gold = append(gold, row);
					}
				})
			} else {
				thValue := h.DOM.Find("th").First()
				if thValue.Text() == "Emas Batangan Gift Series" {
					breakLoop = true;
				}
			}
		})
		for _, item := range gold {
			fmt.Printf("Weight: %s, Price: %s, Date: %s\n", item.Weight, item.Price, item.Date)
		}
	})

    collector.Visit(url)

}