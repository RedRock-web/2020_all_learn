//@program: work
//@description: 
//@author: edte
//@create: 2020-05-30 19:49
package app

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	baseUrl = "https://bgm.tv/game/browser?sort=rank&page="
)

type Url struct {
	url string
}

func (u *Url) Url() string {
	return u.url
}

func (u *Url) SetUrl(url string) {
	u.url = url
}

func NewUrl(i int) *Url {
	return &Url{url: baseUrl + strconv.Itoa(i)}
}

func (u *Url) GetBody() (string, error) {
	resp, err := http.Get(u.Url())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
