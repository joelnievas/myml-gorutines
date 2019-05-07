package myml

import (
	"io/ioutil"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"fmt"
	"encoding/json"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	CurrencyID         string `json:"currency_id"`
	GeoInformation     struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`

}


const urlCounty = "https://api.mercadolibre.com/countries/"

func (country *Country) Get() *apierrors.ApiError {
	if country.ID == "" {
		return &apierrors.ApiError{
			Message: "countryID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%s", urlCounty, country.ID)
	response, err := http.Get(final)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(data), &country); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
