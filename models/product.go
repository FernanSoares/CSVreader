package models

type Product struct {
	ID           string
	Name         string
	Description  string
	Brand        string
	Category     string
	Price        float64
	Currency     string
	Stock        int
	EAN          string
	Color        string
	Size         string
	Availability string
	InternalID   string
}
