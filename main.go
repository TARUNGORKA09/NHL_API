package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/TARUNGORKA09/NHL_API/nhlapi"
)

func main() {
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening the file rosters.txt: %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlapi.GetAllTeams()
	if err != nil {
		log.Fatalf("error while getting all teams: %v", err)
	}

	for _, team := range teams {
		log.Println("---------------------")
		log.Printf("Name: %s", team.Name)
		log.Println("---------------------")
	}

	log.Printf("took %v", time.Now().Sub(now).String())

}