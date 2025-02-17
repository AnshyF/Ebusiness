package model

// Product 商品模型
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	// 可以根据需要添加更多字段
}
