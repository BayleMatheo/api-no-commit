package Groupie

// Créer un structure Global -> Artist, RElations (map)

type rawArtist struct {
	Id             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      string              `json:"locations"`
	ConcertDates   string              `json:"concertDates"`
	Relations      string              `json:"relations"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// type rawLocations struct {
// 	ID        int      `json:"id"`
// 	Locations []string `json:"locations"`
// 	Dates     string   `json:"dates"`
// }

// type rawDates struct {
// 	ID    int      `json:"id"`
// 	Dates []string `json:"dates"`
// }

type rawRelations struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	}
}

type rawRelation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
