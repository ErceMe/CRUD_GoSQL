package entities

type Product struct {
	ID_Product int
	Name       string
	Created_at string
	Updated_at string
}

type Variant struct {
	ID           int
	Variant_name string
	Quantity     int
	Product_id   int
	Created_at   string
	Updated_at   string
}
