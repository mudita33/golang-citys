package city

import (
	"encoding/json"
	"fmt"
	"sort"
)

type CityResponse struct {
	Citys []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"citys"`
}

type DistrictResponse struct {
	Districts []struct {
		ID   string `json:"id"`
		IlID string `json:"il_id"`
		Name string `json:"name"`
	} `json:"districts"`
}

type Districts struct {
	ID   string `json:"id"`
	IlID string `json:"il_id"`
	Name string `json:"name"`
}

func loadCity() CityResponse {

	rawIn := json.RawMessage(cityData)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}

	var cityResponse CityResponse

	json.Unmarshal(bytes, &cityResponse)

	sort.SliceStable(cityResponse.Citys, func(i, j int) bool {
		return cityResponse.Citys[i].Name < cityResponse.Citys[j].Name
	})

	cityResponse = CityResponse{Citys: cityResponse.Citys}

	return cityResponse

}

func loadDistrict() map[string][]Districts {
	rawIn := json.RawMessage(districtData)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}

	var districts DistrictResponse

	mapDistrict := make(map[string][]Districts)

	json.Unmarshal(bytes, &districts)

	sort.SliceStable(districts.Districts, func(i, j int) bool {
		return districts.Districts[i].ID < districts.Districts[j].ID
	})

	for _, v := range districts.Districts {

		if _, ok := mapDistrict[v.IlID]; !ok {
			mapDistrict[v.IlID] = []Districts{}
		}

		mapDistrict[v.IlID] = append(mapDistrict[v.IlID], v)

	}

	return mapDistrict

}

type city struct {
	citys     CityResponse
	districts map[string][]Districts
}

// NewApp creates a new instance of App
func NewCityService() *city {

	a := loadCity()
	b := loadDistrict()

	return &city{a, b}
}

func (r city) GetCities() CityResponse {

	return r.citys
}

func (r city) GetDistrict() map[string][]Districts {

	return r.districts
}
