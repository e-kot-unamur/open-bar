package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// This is generic to all received events
type websocketEvent struct {
	Type string `json:"type"`

	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Debt             int     `json:"debt"`
	Price            float64 `json:"price"`
	ResetParticipant bool    `json:"resetParticipant"`
}

// Following structures are the answers to websocket received events
// They keep the same Type, keep fields or adapt them and have the
// purpose to clear unused fields to reduce network payloads
type allDataAnswer struct {
	Type string      `json:"type"`
	Data openBarData `json:"data"`
}

type newUserAnswer struct {
	Type string   `json:"type"`
	User userData `json:"user"`
}

type updatePriceAnswer struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type updateDebtAnswer struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
	Debt int    `json:"debt"`
}

type resetAnswer struct {
	Type             string `json:"type"`
	ResetParticipant bool   `json:"resetParticipant"`
}

////////////////////////////////////////////////////
type userData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Debt int    `json:"debt"`
}

type history struct {
	Date         time.Time `json:"date"`
	TargetID     int       `json:"targetId"`
	NumberOfBars int       `json:"numberOfBars"`
}

type openBarData struct {
	Price   float64    `json:"price"`
	History []history  `json:"history"`
	Users   []userData `json:"users"`
}

func load(path string) (dataJSON openBarData) {
	log.Println("Loading data")
	if _, err := os.Stat(path); err == nil {
		// File does exist
		dataBytes, err := ioutil.ReadFile(path)
		err = json.Unmarshal(dataBytes, &dataJSON)
		if err != nil {
			log.Fatalln("Something went wrong while reading data")
		}
	} else {
		// File doesn't exist
		dataJSON = openBarData{}
		dataJSON.save(historyFile)
	}
	return dataJSON
}

func (data openBarData) save(destination string) {
	formatedData, _ := json.MarshalIndent(data, "", "	")
	if err := ioutil.WriteFile(destination, formatedData, os.ModePerm); err != nil {
		log.Println("Something went wrong")
	} else {
		log.Println("Data successfully saved")
	}
}
