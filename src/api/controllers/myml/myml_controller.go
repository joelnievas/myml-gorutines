package myml

import (
	"github.com/gin-gonic/gin"
	myml2 "github.com/mercadolibre/myml-gorutines/src/api/domain/myml"
	"github.com/mercadolibre/myml-gorutines/src/api/services/myml"
	"github.com/mercadolibre/myml-gorutines/src/api/utils/apierrors"
	"net/http"
	"strconv"
	"sync"
)

const (
	paramUserID = "userID"
)

func GetUser(c *gin.Context) {
	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}
	user, apiErr := myml.GetUserFromAPI(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}

	cadena := make(chan myml2.Myml, 3)

	cadena <- myml2.Myml{
		User: *user,
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go myml.GetSiteFromAPI(user.SiteID, cadena, &wg)
	go myml.GetCountryFromAPI(user.CountryID, cadena, &wg)
	wg.Wait()
	var objetos myml2.Myml

	aux := <-cadena
	objetos.User = aux.User
	aux = <-cadena
	objetos.Site = aux.Site
	objetos.Country = aux.Country
	aux = <-cadena
	if objetos.Country.ID == "" {
		objetos.Country = aux.Country
	} else {
		objetos.Site = aux.Site
	}
	c.JSON(http.StatusOK, objetos)
	return
}
