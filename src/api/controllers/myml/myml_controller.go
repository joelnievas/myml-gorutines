package myml

import (
    "fmt"
    myml2 "github.com/mercadolibre/myml/src/api/domain/myml"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/mercadolibre/myml/src/api/services/myml"
    "net/http"
    "github.com/mercadolibre/myml/src/api/utils/apierrors"
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

    cadena := make(chan myml2.Myml,3)

    cadena <- myml2.Myml{
        User: *user,
    }

    var wg sync.WaitGroup
    wg.Add(2)
    fmt.Print("antes de entrar ")
    go myml.GetSiteFromAPI(user.SiteID, cadena, &wg)
    go myml.GetCountryFromAPI(user.CountryID, cadena, &wg)
    wg.Wait()
    fmt.Print("-------- salio")

    var objetos myml2.Myml

    aux := <- cadena
    objetos.User = aux.User
    //fmt.Print(objetos)
    aux = <- cadena
    objetos.Site = aux.Site
    objetos.Country = aux.Country
    //fmt.Print(objetos)
    aux = <- cadena
    if objetos.Country.ID== ""{
        objetos.Country = aux.Country
    }else{
        objetos.Site = aux.Site
    }
    c.JSON(http.StatusOK, objetos)
    return
}

