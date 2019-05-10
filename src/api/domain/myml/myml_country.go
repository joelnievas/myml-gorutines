package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml-gorutines/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

type Country struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CurrencyID string `json:"currency_id"`
}

const urlCounty = "http://localhost:8081/country/"

func (country *Country) Get() *apierrors.ApiError {
	if country.ID == "" {
		return &apierrors.ApiError{
			Message: "countryID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%s", urlCounty, country.ID)
	response, err := http.Get(final)
	defer response.Body.Close()
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
