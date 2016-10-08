package ImgScraper

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func getImgURLs(rowURL string) []string {
	urlsPic := getURLsPic(rowURL)
	fmt.Println("全URL取得。フィルターをかけます。")
	var checkedURLsPic []string
	result := make([]bool, 100)
	fmt.Println("URLのフィルターが完了")

	c0 := make(chan []bool)
	c1 := make(chan []bool)
	c2 := make(chan []bool)
	c3 := make(chan []bool)
	c4 := make(chan []bool)
	c5 := make(chan []bool)
	c6 := make(chan []bool)
	c7 := make(chan []bool)
	c8 := make(chan []bool)
	c9 := make(chan []bool)

	go func() {
		result0 := checkPicSize(urlsPic[0:9])
		c0 <- result0
	}()
	go func() {
		result1 := checkPicSize(urlsPic[10:19])
		c1 <- result1
	}()
	go func() {
		result2 := checkPicSize(urlsPic[20:29])
		c2 <- result2
	}()
	go func() {
		result3 := checkPicSize(urlsPic[30:39])
		c3 <- result3
	}()
	go func() {
		result4 := checkPicSize(urlsPic[40:49])
		c4 <- result4
	}()
	go func() {
		result5 := checkPicSize(urlsPic[50:59])
		c5 <- result5
	}()
	go func() {
		result6 := checkPicSize(urlsPic[60:69])
		c6 <- result6
	}()
	go func() {
		result7 := checkPicSize(urlsPic[70:79])
		c7 <- result7
	}()
	go func() {
		result8 := checkPicSize(urlsPic[80:89])
		c8 <- result8
	}()
	go func() {
		result9 := checkPicSize(urlsPic[90:99])
		c9 <- result9
	}()

	result = <-c0
	result = append(result, <-c1...)
	result = append(result, <-c2...)
	result = append(result, <-c3...)
	result = append(result, <-c4...)
	result = append(result, <-c5...)
	result = append(result, <-c6...)
	result = append(result, <-c7...)
	result = append(result, <-c8...)
	result = append(result, <-c9...)

	for i, bool := range result {
		if bool == true {
			checkedURLsPic = append(checkedURLsPic, urlsPic[i])
		}
	}

	return checkedURLsPic
}

func checkPicSize(partOfURLsPic []string) []bool {
	largePic := make([]bool, 10)
	for i, urlPic := range partOfURLsPic {
		resp, err := http.Get(urlPic)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		fileSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
		if err != nil {
			log.Fatal(err)
		}
		if fileSize > 30000 {
			largePic[i] = true
		}
	}
	return largePic
}
