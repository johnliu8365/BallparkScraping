package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	ballClubs := make(map[int]map[string]string)
	var value int

	resp, err := soup.Get("http://mlb.mlb.com/team/index.jsp")
	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)

	alTeam := doc.FindAll("ul", "class", "al team")
	nlteam := doc.FindAll("ul", "class", "nl team")

	for _, Teams := range alTeam {
		d := make(map[string]string)
		d["name"] = Teams.Find("h5").Find("a").Text()
		d["ballpark"] = Teams.FindAll("li")[1].Text()
		d["address"] = Teams.FindAll("li")[2].Text()
		d["location"] = Teams.FindAll("li")[3].Text()
		ballClubs[value] = d
		value++
	}

	for _, Teams := range nlteam {
		d := make(map[string]string)
		d["name"] = Teams.Find("h5").Find("a").Text()
		d["ballpark"] = Teams.FindAll("li")[1].Text()
		d["address"] = Teams.FindAll("li")[2].Text()
		d["location"] = Teams.FindAll("li")[3].Text()
		ballClubs[value] = d
		value++
	}

	printMap(ballClubs)
	outputJSON(ballClubs)
}

func printMap(c map[int]map[string]string) {
	for _, d := range c {
		fmt.Println("-----------")
		for Key, Data := range d {
			fmt.Println(Key + ": " + Data)
		}
	}
}

func outputJSON(c map[int]map[string]string) {
	b, _ := json.Marshal(c)
	err := ioutil.WriteFile("data.json", []byte(b), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
