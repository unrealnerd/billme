package inventory

import ()

//ToItem ... from DTO to native type
func ToItem(itemDTO DTO) Item {
	return Item{
		Code:         itemDTO.Code,
		SellingPrice: itemDTO.SellingPrice,
		Description:  itemDTO.Description,
	}
}

//ToItemDTO ...
func ToItemDTO(item Item) DTO {
	return DTO{
		Code:         item.Code,
		Description:  item.Description,
		SellingPrice: item.SellingPrice,
	}
}

//ToItemDTOs ...
func ToItemDTOs(items []Item) []DTO {
	itemDTOs := make([]DTO, len(items))

	for i, itm := range items {
		itemDTOs[i] = ToItemDTO(itm)
	}

	return itemDTOs
}
