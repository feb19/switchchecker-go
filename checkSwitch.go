package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "strings"
)

const (
    //Swith url
    URL = "https://www.amazon.co.jp/Nintendo-Switch-Joy-Con-%E3%83%8D%E3%82%AA%E3%83%B3%E3%83%96%E3%83%AB%E3%83%BC-%E3%83%8D%E3%82%AA%E3%83%B3%E3%83%AC%E3%83%83%E3%83%89/dp/B01NCXFWIZ"
    //キーワード
    KEYWORD = "この商品は、Amazon.co.jp が販売、発送します。"
)

//ページのチェック
func CheckPage(url string) {
    doc, _ := goquery.NewDocument(url)
    t := doc.Find("#merchant-info").Text()
    if strings.Contains(t, KEYWORD) {
        fmt.Println("Switch売ってるよ！！！")
        //PostSlack()
    } else {
        fmt.Println("Switch売ってない...")
    }
}

func main() {
    CheckPage(URL)
}
