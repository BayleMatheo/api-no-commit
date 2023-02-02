package Groupie

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}

func Compile() []Artist {
	allArtists := []Artist{}
	rawArtists := getArtist()
	locations := getLocations()
	dates := getDates()
	relations := getRelations()

	for i := 0; i < len(rawArtists); i++ {
		allArtists = append(allArtists, Artist{
			Id:           rawArtists[i].Id,
			Image:        rawArtists[i].Image,
			Name:         rawArtists[i].Name,
			Members:      rawArtists[i].Members,
			CreationDate: rawArtists[i].CreationDate,
			FirstAlbum:   Date(rawArtists[i].FirstAlbum),
			Locations:    locations.Index[i].Locations,
			ConcertDates: dates.Index[i].Dates,
			Relations:    relations.Index[i].DatesLocations,
		})
	}
	return allArtists
}
