package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)

	writer.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"})

	for _, page := range pages {
		var row []string = make([]string, 5)

		row[0] = page.URL
		row[1] = page.H1
		row[2] = page.FirstParagraph
		row[3] = strings.Join(page.OutgoingLinks, ";")
		row[4] = strings.Join(page.ImageURLs, ";")

		writer.Write(row)
	}
	return nil

}
