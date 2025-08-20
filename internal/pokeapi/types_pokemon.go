package pokeapi

type Pokemon struct {
	BaseExperience         int    `json:"base_experience"`
	ID                     int    `json:"id"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Name                   string `json:"name"`
	Height                 int    `json:"height"`
	Weight                 int    `json:"weight"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
