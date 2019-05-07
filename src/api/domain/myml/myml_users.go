package myml

import (
    "io/ioutil"
    "github.com/mercadolibre/myml/src/api/utils/apierrors"
    "net/http"
    "fmt"
    "encoding/json"
)

type User struct {
    ID               int64  `json:"id"`
    Nickname         string `json:"nickname"`
    RegistrationDate string `json:"registration_date"`
    CountryID        string `json:"country_id"`
    Email            string `json:"email"`
    SiteID           string `json:"site_id"`
}


const urlUsers = "https://api.mercadolibre.com/users/"

func (user *User) Get() *apierrors.ApiError {
    if user.ID == 0 {
        return &apierrors.ApiError{
            Message: "userID is empty",
            Status:  http.StatusBadRequest,
        }
    }

    final := fmt.Sprintf("%s%d", urlUsers, user.ID)
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

    if err := json.Unmarshal([]byte(data), &user); err != nil {
        return &apierrors.ApiError{
            Message: err.Error(),
            Status:  http.StatusInternalServerError,
        }
    }

    return nil
}
