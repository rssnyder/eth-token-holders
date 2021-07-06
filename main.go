package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

const (
	ChainURL = "https://%s/token/%s"
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

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.GET("/:chain/:contract", func(c *gin.Context) {
		contract := c.Param("contract")
		var chain string
		switch c.Param("chain") {
		case "ethereum":
			chain = "etherscan.io"
		case "binance-smart-chain":
			chain = "bscscan.com"
		default:
			c.String(http.StatusBadRequest, "%d", 0)
			return
		}
		tokenHolders := holders(chain, contract)
		if tokenHolders != "" {
			c.String(http.StatusOK, "%s", tokenHolders)
			return
		}
		c.String(http.StatusBadRequest, "%d", 0)
	})

	router.Run(*addr)
}

func holders(chain, contract string) string {
	var holders string

	reqURL := fmt.Sprintf(ChainURL, chain, contract)

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
