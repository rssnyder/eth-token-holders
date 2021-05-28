package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

const (
	EthURL = "https://etherscan.io/token/%s"
)

var (
	addr *string
)

func init() {
	addr = flag.String("addr", "localhost:8080", "address to bind http server to.")
	flag.Parse()
}

func main() {
	router := gin.Default()

	router.GET("/:contract", func(c *gin.Context) {
		contract := c.Param("contract")
		tokenHolders := holders(contract)
		if tokenHolders != "" {
			c.String(http.StatusOK, "%s", tokenHolders)
			return
		}
		c.String(http.StatusBadRequest, "%d", 0)
	})

	router.Run(*addr)
}

func holders(contract string) string {
	var holders string

	reqURL := fmt.Sprintf(EthURL, contract)

	response, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("div").Each(func(index int, element *goquery.Selection) {
		exists := element.HasClass("mr-3")
		if exists {
			holders = strings.TrimSpace(element.Text())
		}
	})

	return holders
}
