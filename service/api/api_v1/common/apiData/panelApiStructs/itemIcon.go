package panelApiStructs

import (
	"Mi-Panel/api/api_v1/common/apiData/commonApiStructs"
	"Mi-Panel/models"
)

type ItemIconEditRequest struct {
	models.ItemIcon
	IconJson string
}

type ItemIconSaveSortRequest struct {
	SortItems       []commonApiStructs.SortRequestItem `json:"sortItems"`
	ItemIconGroupId uint                               `json:"itemIconGroupId"`
}
