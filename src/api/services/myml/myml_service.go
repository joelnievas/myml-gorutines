package myml

import (
    "fmt"
    "github.com/mercadolibre/myml/src/api/domain/myml"
    "github.com/mercadolibre/myml/src/api/utils/apierrors"
    "sync"
)


func GetUserFromAPI(userID int64) (*myml.User, *apierrors.ApiError) {
    user := &myml.User{
        ID: userID,
    }
    if apiErr := user.Get(); apiErr != nil {
        return nil, apiErr
    }
    return user, nil
}

//GetSiteFromAPI
func GetSiteFromAPI(siteID string, ch chan myml.Myml, wb *sync.WaitGroup) {
    defer wb.Done()
    site := &myml.Site{
        ID: siteID,
    }
    if apiErr := site.Get(); apiErr != nil {
    }
    ch <- myml.Myml{
        Site:    *site,
    }
    fmt.Println(site)

}

func GetCategoryFromAPI(categoryID string) (*myml.Category, *apierrors.ApiError) {
    category := &myml.Category{
        ID: categoryID,
    }
    if apiErr := category.Get(); apiErr != nil {
        return nil, apiErr
    }
    return category, nil
}


func GetCountryFromAPI(countryID string, ch chan myml.Myml, wb *sync.WaitGroup)  {
    defer wb.Done()
    country := &myml.Country{
        ID: countryID,
    }
    if apiErr := country.Get(); apiErr != nil {
    }

    ch <- myml.Myml{
        Country: *country,
    }


}


