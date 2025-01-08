package models

type Donation struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	DonorName   string `json:"donorName"`
	Location    string `json:"location"`
	Contact     string `json:"contact"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
}
