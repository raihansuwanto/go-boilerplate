package dto

type ProductLoaderRequest struct {
	ID string `json:"id"`
}

type ProductLoaderResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	CategoryName string  `json:"category_name"`
	CategoryID   string  `json:"category_id"`
}
