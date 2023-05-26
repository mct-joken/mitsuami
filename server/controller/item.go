package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mct-joken/mitsuami/model/api"
	"github.com/mct-joken/mitsuami/repository"
	"github.com/mct-joken/mitsuami/service"
	"net/http"
)

type ItemController struct {
	itemService service.ItemService
}

func NewItemController(logRepository repository.LogRepository, itemRepository repository.ItemRepository) *ItemController {
	return &ItemController{
		itemService: *service.NewItemService(itemRepository, logRepository),
	}
}

func (c *ItemController) CreateItem(e echo.Context) error {
	req := api.CreateItemRequestJSON{}
	if err := e.Bind(&req); err != nil {
		return e.String(http.StatusBadRequest, "invalid json type")
	}

	args := service.CreateItemArgs{
		ID:          req.ID,
		ItemType:    req.Type,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
	}
	if err := c.itemService.Create(args); err != nil {
		return err
	}

	res := api.CreateItemResponseJSON{
		ID:          req.ID,
		Type:        req.Type,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Status: struct {
			Status  int    `json:"status"`
			User    string `json:"user,omitempty"`
			DueDate string `json:"dueDate,omitempty"`
		}(struct {
			Status  int
			User    string
			DueDate string
		}{Status: -1}),
	}

	return e.JSON(http.StatusCreated, res)
}

func (c *ItemController) GetItem(e echo.Context) error {
	r := c.itemService.GetItemList()
	res := make(api.GetItemListResponseJSON, len(r))
	for i, v := range r {
		res[i] = api.GetItemResponseJSON{
			ID:          v.GetID(),
			Type:        v.GetItemType(),
			Code:        v.GetCode(),
			Name:        v.GetName(),
			Description: v.GetDescription(),
			Image:       v.GetImage(),
			Status: struct {
				Status  int    `json:"status"`
				User    string `json:"user,omitempty"`
				DueDate string `json:"dueDate,omitempty"`
			}(struct {
				Status  int
				User    string
				DueDate string
				// ToDo: Itemのステータスを取得する
			}{Status: -1}),
		}
	}

	return e.JSON(http.StatusOK, res)
}

func (c *ItemController) GetItemByID(e echo.Context) error {
	r, err := c.itemService.GetItem(e.Param("id"))
	if err != nil {
		return e.String(http.StatusNotFound, "not found")
	}

	res := api.GetItemResponseJSON{
		ID:          r.GetID(),
		Type:        r.GetItemType(),
		Code:        r.GetCode(),
		Name:        r.GetName(),
		Description: r.GetDescription(),
		Image:       r.GetImage(),
		Status: struct {
			Status  int    `json:"status"`
			User    string `json:"user,omitempty"`
			DueDate string `json:"dueDate,omitempty"`
		}(struct {
			Status  int
			User    string
			DueDate string
			// ToDo: Itemのステータスを取得する
		}{Status: -1}),
	}

	return e.JSON(http.StatusOK, res)
}

func (c *ItemController) StartUsing(e echo.Context) error {
	// ToDo: ユーザーIDを入れる
	if err := c.itemService.StartUsing(e.Param("id"), ""); err != nil {
		return e.JSON(http.StatusBadRequest, api.StartUsingItemResponseJSON{Status: "failed"})
	}
	return e.JSON(http.StatusOK, api.StartUsingItemResponseJSON{Status: "ok"})
}

func (c *ItemController) EndUsing(e echo.Context) error {
	// ToDo: ユーザーIDを入れる
	if err := c.itemService.EndUsing(e.Param("id"), ""); err != nil {
		return e.JSON(http.StatusBadRequest, api.EndUsingItemResponseJSON{Status: "failed"})
	}
	return e.JSON(http.StatusOK, api.EndUsingItemResponseJSON{Status: "ok"})
}
