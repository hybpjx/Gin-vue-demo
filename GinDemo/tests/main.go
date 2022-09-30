package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.ahhuoshan.gov.cn/site/label/8888?IsAjax=1&dataType=html&_=0.3327293013106256&labelName=publicInfoList&siteId=6786551&pageSize=15&pageIndex=2&action=list&isDate=true&dateFormat=yyyy-MM-dd&length=50&organId=6596221&type=4&catId=7065111&cId=&result=%E6%9A%82%E6%97%A0%E7%9B%B8%E5%85%B3%E4%BF%A1%E6%81%AF&file=%2Fc3%2Fhs%2FpublicInfoList_newest", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Ls-Language", "zh")
	req.Header.Set("Referer", "https://www.ahhuoshan.gov.cn/public/column/6596221?type=4&action=list&nav=3&catId=7065111")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Cookie", "__jsluid_s=57156a4cce4bc7a76ab632065364b484; UM_distinctid=18358747c061f9-013532c507f75a-26021c51-1fa400-18358747c07b40; luan_govc_SHIROJSESSIONID=1cd716e6-a8c1-4207-b94b-c651504b5098; CNZZDATA1279628965=496900267-1663633549-null%7C1664416610")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
