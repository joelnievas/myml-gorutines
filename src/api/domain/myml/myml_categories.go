package myml

import (
	"io/ioutil"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"fmt"
	"encoding/json"
)

type Category struct {
	ID               string  `json:"id"`
	Name			 string `json:"name"`
}


const urlCategory = "https://api.mercadolibre.com/sites/"

func (category *Category) Get() *apierrors.ApiError {
	if category.ID == "" {
		return &apierrors.ApiError{
			Message: "categoryID is empty",
			Status:  http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%d", urlCategory, category.ID)
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

	if err := json.Unmarshal([]byte(data), &category); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
