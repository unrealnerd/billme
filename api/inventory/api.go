package inventory

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//API ...
type API struct {
	service Service
}

//ProvideInventoryAPI ...
func ProvideInventoryAPI(s Service) API {
	return API{service: s}
}

//Add ...
func (a *API) Find(c *gin.Context) {

	ID := c.Param("id")
	items := a.service.Find(ID)

	c.JSON(http.StatusOK, items)
}

//Ping ...
func (a *API) Ping(c *gin.Context) {

	a.service.Ping()

	c.JSON(http.StatusOK, nil)
}

//Add ...
func (a *API) Add(c *gin.Context) {

	var item Item
	c.BindJSON(&item)
	a.service.Add(item)

	c.JSON(http.StatusOK, nil)
}
