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

// //Find ...
// func (a *API) Find(c *gin.Context) {

// 	quoteID := c.Param("id")
// 	affinityName := c.Param("affinityname")

// 	qhs := a.service.Ping()

// 	c.JSON(http.StatusOK, gin.H{"qhs": ToItemDTOs(qhs)})
// }

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
