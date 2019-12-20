package service

type fulfillmentStatus struct {
	SKU             string `json:"sku"`
	ShipsWithin     int    `json:"shipsWithin"`
	QuantityInStock int    `json:"quantityInStock"`
}

type catalogItem struct {
	ProductID       int    `json:"productId"`
	Description     string `json:"description"`
	Price           uint32 `json:"price"`
	SKU             string `json:"sku"`
	ShipsWithin     int    `json:"shipsWithin"`
	QuantityInStock int    `json:"quantityInStock"`
}
