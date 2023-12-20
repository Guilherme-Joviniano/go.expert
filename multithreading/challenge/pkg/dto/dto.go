package dto

type ZipCodeDetails struct {
	ZipCode      string `json:"zip_code"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	State        string `json:"state"`
}

func NewZipCodeDetails(zipCode, city, neighborhood, street, state string) *ZipCodeDetails {
	return &ZipCodeDetails{
		ZipCode:      zipCode,
		City:         city,
		Neighborhood: neighborhood,
		Street:       street,
		State:        state,
	}
}
