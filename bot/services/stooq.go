package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (b *Bot) Process(data []byte, stockCode string) (string, error) {
	reader := csv.NewReader(strings.NewReader(string(data)))
	var price string

	// Read every row and process
	rows, err := reader.ReadAll()
	if err != nil {
		return "", err
	}
	for _, row := range rows {
		if len(row) >= 8 {
			price = row[6]
			if price == "N/D" {
				return "N/D", nil
			}
		}
	}
	return fmt.Sprintf("%s quote is $%v per share", strings.ToUpper(stockCode), price), nil
}

func (b *Bot) FetchFile(stockCode string) ([]byte, error) {
	url := fmt.Sprintf("%s%s&f=sd2t2ohlcv&h&e=csv", b.Env.StooqUrl, stockCode)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in HTTP request: %s", response.Status)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
