package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myblog-gin/model"
)

func InitSearchJson() {
	var articleList []model.Article
	// articleList := []struct {
	// 	ID      uint
	// 	Content string `json:"content"`
	// 	Title   string `json:"title"`
	// }{}
	GetAllArt(&articleList)
	v, err := json.Marshal(articleList)
	if err != nil {
		fmt.Println("marshal failed!", err)
		return
	}
	// fmt.Println(v)
	// fmt.Println(articleList)
	err = ioutil.WriteFile("Search.json", v, 0744)
	if err != nil {
		return
	}
}
