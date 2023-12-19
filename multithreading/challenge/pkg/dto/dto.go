package dto

type ZipCodeDetails struct {
	Address string `json:"address"`
}

func NewZipCodeDetails(
	address string,
) *ZipCodeDetails {
	return &ZipCodeDetails{
		Address: address,
	}
}


