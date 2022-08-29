package habr

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	part    = "https://habr.com/ru/post/"
	dataUrl = "https://habr.com/kek/v2/articles/?hub=go&sort=all&fl=ru&hl=ru&page=1"
)

type Ids struct {
	Id []string `json:"articleIds"`
}

func GetPostsIds() (Ids, error) {
	var ids Ids

	request, err := http.Get(dataUrl)
	if err != nil {
		log.Fatalf(err.Error())
		return ids, err
	}

	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalf(err.Error())
		return ids, err
	}

	err = json.Unmarshal(content, &ids)
	if err != nil {
		log.Fatalf(err.Error())
		return ids, err
	}
	return ids, nil
}

func GetNewPostsUrls(lastPost string, newIds Ids) []string {
	var urls []string

	for _, el := range newIds.Id {
		if el == lastPost {
			return urls
		}
		urls = append(urls, part+el)
	}

	return urls
}
