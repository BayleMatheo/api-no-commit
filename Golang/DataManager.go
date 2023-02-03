package Groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getData(url string, data interface{}) {
	rawData := getRawData(url)
	err := json.Unmarshal(rawData, &data)
	if err != nil {
		log.Panic("Problème dans la fonction getData lors du déclassement des données :", err)
	}
}

func getRawData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Panic("Problème dans la fonction getRawData lors de l'obtention de la réponse :", err)
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic("Problème dans la fonction getRawData lors de la lecture de la réponse :", err)
		return nil
	}
	return responseData
}

func getArtist() []rawArtist {
	artists := []rawArtist{}
	getData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	return artists
}

func getArtists(Id string) rawArtist {
	artist := rawArtist{}
	getData("https://groupietrackers.herokuapp.com/api/artists/"+Id, &artist)
	return artist
}

// func getDates() rawDates {
// 	dates := rawDates{}
// 	getData("https://groupietrackers.herokuapp.com/api/dates", &dates)
// 	return dates
// }

// func getLocations() rawLocations {
// 	locations := rawLocations{}
// 	getData("https://groupietrackers.herokuapp.com/api/locations", &locations)
// 	return locations
// }

func getRelations(Id string) rawRelation {
	// Initalize to store the raw data
	relation := rawRelation{}
	getData("https://groupietrackers.herokuapp.com/api/relation/"+Id, &relation)
	// Convert the names to the correct format
	// EX: "los_angeles-usa" -> "Los Angeles, USA"
	marshalledRelations, err := json.Marshal(relation)
	if err != nil {
		panic(err)
	}
	marshalledRelations = []byte(
		Title(
			strings.ReplaceAll(
				strings.ReplaceAll(
					string(marshalledRelations), "_", " "), "-", ", ")))
	// Reset to avoid duplicates when unmarshalling
	relation = rawRelation{}
	err = json.Unmarshal(marshalledRelations, &relation)
	if err != nil {
		panic(err)
	}

	// Convert the dates to the correct format
	for location, dates := range relation.DatesLocations {
		// Convert the dates to the correct format
		// EX: "2019-01-01" -> "January 1, 2019"
		for dateIndex, date := range dates {
			relation.DatesLocations[location][dateIndex] = Date(date)
		}
	}
	fmt.Println(relation)
	return relation
}
