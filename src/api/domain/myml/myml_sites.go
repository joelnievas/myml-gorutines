package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml-gorutines/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

type (
	Site struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)

const urlSite = "http://localhost:8081/site/"

func (site *Site) Get() *apierrors.ApiError {
	if site.ID == "" {
		return &apierrors.ApiError{
			Message: "userID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%s", urlSite, site.ID)

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

	if err := json.Unmarshal([]byte(data), &site); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
