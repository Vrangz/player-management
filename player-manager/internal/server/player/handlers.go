package player

import (
	"net/http"
	"player-manager/internal/server/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:route GET /players/{username} player playerInfo
//
// Gets player information.
//
//		Responses:
//		  200: PlayerResponse
//	      400: CommonError
func (ctrl *Controller) GetPlayer(c *gin.Context) {
	var username = c.Param("username")

	player, err := ctrl.playerRepository.GetPlayer(c, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, "failed to get player", err))
		return
	}

	c.JSON(http.StatusOK, ToPlayerResponse(player))
}

// swagger:route GET /players/{username}/items player playerItems
//
// List all items of a player.
//
//		Responses:
//		  200: PlayerItemsResponse
//	      400: CommonError
func (ctrl *Controller) ListItems(c *gin.Context) {
	var username = c.Param("username")

	items, err := ctrl.playerRepository.ListItems(c, username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, "failed to get items", err))
		return
	}

	c.JSON(http.StatusOK, ToItemsResponse(items))
}

// swagger:route PUT /players/{username}/items/{item} player playerAddItem
//
// Adds defined item to the player.
//
//		Responses:
//		  204: NoContentResponse
//	      400: CommonError
func (ctrl *Controller) AddItem(c *gin.Context) {
	var (
		username = c.Param("username")
		item     = c.Param("item")
		quantity int
		err      error
	)

	if quantity, err = ctrl.parseQuantity(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, "invalid quantity", err))
		return
	}

	if err = ctrl.playerRepository.AddItem(c, username, item, quantity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, "failed to add item", err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// swagger:route DELETE /players/{username}/items/{item} player playerDeleteItem
//
// Deletes defined item from the player.
//
//		Responses:
//		  204: NoContentResponse
//	      400: CommonError
func (ctrl *Controller) DeleteItem(c *gin.Context) {
	var (
		username = c.Param("username")
		item     = c.Param("item")
		quantity int
		err      error
	)

	if quantity, err = ctrl.parseQuantity(c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, "invalid quantity", err))
		return
	}

	if err = ctrl.playerRepository.DeleteItem(c, username, item, quantity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, "failed to delete item", err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// swagger:route POST /players/{username}/action/build player build
//
// If the player has enough of resources, he can perform build action.
//
//		Responses:
//		  204: NoContentResponse
//	      400: CommonError
func (ctrl *Controller) Build(c *gin.Context) {
	var username = c.Param("username")

	if err := ctrl.playerRepository.Build(c, username); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.NewError(http.StatusBadRequest, "failed to build", err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (ctrl *Controller) parseQuantity(c *gin.Context) (quantity int, err error) {
	var (
		strQuantity = c.Query("quantity")
	)

	if strQuantity == "" {
		return 1, nil
	}

	if quantity, err = strconv.Atoi(strQuantity); err != nil {
		return quantity, err
	}

	return quantity, err
}
