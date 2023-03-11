package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Beef struct {
	Beef map[string]int `json:"beef"`
}

func main() {
	r := gin.Default()

	r.GET("/beef/summary", func(c *gin.Context) {
		var beef Beef
		beef.Beef = findAllMeats(getMeat())
		c.JSON(http.StatusOK, beef)
	})

	r.Run()
}

func getMeat() []string {
	url := "https://baconipsum.com/api/?type=all-meat&format=text"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	res := string(body)
	allmeats := strings.Fields(res)

	tmp := make(map[string]bool)
	result := []string{}

	for _, m := range allmeats {
		meat := strings.ToLower(strings.Trim(m, ",."))
		if !tmp[meat] {
			tmp[meat] = true
			result = append(result, meat)
		}
	}

	return result
}

func findAllMeats(allmeats []string) map[string]int {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	res := string(body)
	alltext := strings.Fields(res)

	countMeats := make(map[string]int)

	// iterate over the array and add each element to the map
	for _, m := range allmeats {
		countMeats[m] = 0
	}
	for _, m := range alltext {
		meat := strings.ToLower(strings.Trim(m, ",."))
		if _, ok := countMeats[meat]; ok {
			countMeats[meat]++
		}
	}

	return countMeats
}
