package dto

type VariantRequest struct {
	Name string `json:"name" binding:"required"`
}

type OptionRequest struct {
	VariantIndex int    `json:"variant_index" binding:"required"` // index of parent variant
	Value        string `json:"value" binding:"required"`
}

type SubVariantRequest struct {
	OptionIndices []int  `json:"option_indices" binding:"required"` // indexes of options
	SKU           string `json:"sku" binding:"required"`
	Stock         string `json:"stock"`
}

type CreateProductRequest struct {
	ProductID    int64              `json:"product_id" binding:"required"`
	ProductCode  string             `json:"product_code" binding:"required"`
	ProductName  string             `json:"product_name" binding:"required"`
	ProductImage string             `json:"product_image"`
	CreatedUser  string             `json:"created_user"`
	HSNCode      string             `json:"hsn_code"`

	Variants    []VariantRequest    `json:"variants"`
	Options     []OptionRequest     `json:"options"`
	SubVariants []SubVariantRequest `json:"sub_variants"`
}
