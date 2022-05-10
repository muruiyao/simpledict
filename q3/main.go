package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// type DictResponse struct {
// 	TranslateResult [][]struct {
// 		Tgt string `json:"tgt"`
// 		Src string `json:"src"`
// 	} `json:"translateResult"`
// 	ErrorCode   int    `json:"errorCode"`
// 	Type        string `json:"type"`
// 	SmartResult struct {
// 		Entries []string `json:"entries"`
// 		Type    int      `json:"type"`
// 	} `json:"smartResult"`
// }
type DictResponse struct {
	TranslateResult []string `json:"translateResult"`
	ErrorCode       int      `json:"errorCode"`
	Type            string   `json:"type"`
	SmartResult     struct {
		Entries []string `json:"entries"`
		Type    int      `json:"type"`
	} `json:"smartResult"`
}

func query(word string) {
	//数据准备
	t := time.Now().UnixMilli()          //生成13位时间戳
	r := strconv.FormatInt(t, 10)        //转为字符串类型的时间戳
	rand.Seed(time.Now().UnixNano())     //为生成随机数做准备
	sufix := strconv.Itoa(rand.Intn(10)) //生成0-9之间的整数
	i := r + sufix
	var salt = i
	var ts = r
	// var word = "我爱你"
	var chuan = "fanyideskweb" + word + salt + "Ygy_4c=r#e#4EX^NUGUc5"
	var arr = md5.Sum([]byte(chuan))
	sign := fmt.Sprintf("%x", arr)
	client := &http.Client{}
	var data = strings.NewReader(`i=` + word + `&from=AUTO&to=AUTO&smartresult=dict&client=fanyideskweb&salt=` + salt +
		`&sign=` + sign +
		`&lts=` + ts +
		`&bv=e70edeacd2efbca394a58b9e43a6ed2a
	&doctype=json&version=2.1
	&keyfrom=fanyi.web&action=FY_BY_REALTlME`)
	req, err := http.NewRequest("POST", "https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="96", "Google Chrome";v="96"`)
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("Origin", "https://fanyi.youdao.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.youdao.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "OUTFOX_SEARCH_USER_ID=1709629224@10.108.160.100; OUTFOX_SEARCH_USER_ID_NCOO=178109166.30855334; _ntes_nnid=f6382c9da163857c60fd66ebca2e036d,1616070577146; fanyi-ad-id=305838; fanyi-ad-closed=1; ___rl__test__cookies=1652068360929")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("sjsjsjsjjs")
	// fmt.Printf("%s\n", bodyText)
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(word, " english: ")
	for _, item := range dictResponse.TranslateResult {
		fmt.Println(item)
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
	example: simpleDict hello
			`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}
