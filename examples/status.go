package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/c2technology/pubg-go"
	"os"
)

func main(){

	apiKey := os.Getenv("PUBGAPIKEY")
	api := pubg.Api{
		&http.Client{},
		apiKey,
		"https://api.playbattlegrounds.com",
	}
	status, err := api.Status()
	if err != nil {
		fmt.Println(err)
	}
	out, _ := json.Marshal(status)
	fmt.Println(fmt.Sprintf("Status: %v", string(out)))

	//players, err := api.GetPlayers(pubg.PCNorthAmerica, []string{}, []string{"JimmyTehBanana"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(fmt.Sprintf("Players: %v", players))

	matches, err := api.GetMatches(pubg.PCNorthAmerica, "fb835210-1c54-491f-aea6-e6d584290975")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Matches: %v", matches))
}