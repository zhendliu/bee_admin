package models

import (
	"bee_admin/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BingImg struct {
	Images  []Img `json:"images"`

}
type Img struct {
	StartDate string `json:"startdate"`
	FullStartDate string `json:"fullstartdate"`
	EndDate string `json:"enddate"`
	Url string `json:"url"`
	UrlBase string `json:"urlbase"`
}


// ResourceTreeGridByUserId 根据用户获取有权管理的资源列表，并整理成teegrid格式
func GetBingImg() (string,error) {
	urlHeader := "https://cn.bing.com"
	cacheKey := time.Now().Format("20060102")
	var bingImg BingImg
	err := utils.GetCache(cacheKey, &bingImg)
	if  err == nil {
		fmt.Println("命中缓存")
		return urlHeader+bingImg.Images[0].Url,nil
	}else{
		fmt.Println("遇到的错误是:",err)
	}
	fmt.Println("未命中缓存")
	url := "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN"
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("search query failed: %s", resp.Status)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	err  =json.Unmarshal(data,&bingImg)
	if err !=nil{
		return "",err
	}
	utils.SetCache(cacheKey, &bingImg, 84600)
	fmt.Printf("保存到缓存:cacheKey:%s,bingImg :%v",cacheKey,bingImg)
	return urlHeader+bingImg.Images[0].Url ,nil
}