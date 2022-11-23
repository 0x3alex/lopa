package apis

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func UrbanGetRandom(term string) *Urban {
	//https://api.urbandictionary.com/v0/random
	req := gorequest.New()
	url := "https://api.urbandictionary.com/v0/random"
	if term != "" {
		url = fmt.Sprintf("https://api.urbandictionary.com/v0/define?term=%s",
			term)
	}
	req.Get(url)
	_, body, _ := req.End()

	var result map[string]any
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	items := result["list"].([]interface{})
	var urban Urban
	for _, v := range items {
		info := v.(map[string]any)
		urban.Author = info["author"].(string)
		urban.Word = info["word"].(string)
		urban.Definition = info["definition"].(string)
		urban.ThumbsUp = int(info["thumbs_up"].(float64))
		urban.ThumbsDown = int(info["thumbs_down"].(float64))
		urban.Date = info["written_on"].(string)
		urban.Example = info["example"].(string)
		urban.Link = info["permalink"].(string)
		break
	}
	return &urban
}
