//@program: work
//@description: 
//@author: edte
//@create: 2020-05-30 18:41
package app

import (
	"regexp"
	"work/database/dmysql"
)

const (
	gameCnName = "class=\"l\">(.*?)<\\/a>\\s(<s|<\\/h3>)"
	gameJnName = "<small class=\"grey\">(.*?)</small>"
	gameInfo   = "<p class=\"info tip\">\\s((.*?) / (.*?) / (.*?))<\\/p>"
)

func SaveNames(m int) {
	games, err := GetGames(m)
	if err != nil {
		panic(err)
	}

	for _, v := range games {
		dmysql.G_db.Create(&v)
	}
}

func GetAllNames() {
	for i := 0; i < 1374; i++ {
		//go SaveNames(i)
		SaveNames(i)
	}
}

func GetGames(m int) (games []dmysql.Game, err error) {
	body, err := NewUrl(m).GetBody()
	if err != nil {
		return nil, err
	}

	cnName := regexp.MustCompile(gameCnName)
	cnNames := cnName.FindAllStringSubmatch(body, -1)
	for _, i := range cnNames {
		games = append(games, dmysql.Game{CnName: i[1]})
	}

	jnName := regexp.MustCompile(gameJnName)
	jnNames := jnName.FindAllStringSubmatch(body, -1)
	for k, v := range jnNames {
		games[k].JnName = v[1]
	}

	info := regexp.MustCompile(gameInfo)
	infos := info.FindAllStringSubmatch(body, -1)
	for k, v := range infos {
		games[k].Info = v[1]
	}
	return games, nil
}
