package models

//Book struct
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Category    string `json:"category"`
	OwnerID     string `json:"owner_id"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	NoOfPages   int64  `json:"no_of_pages"`
	DueDate     int64  `json:"due_date"`
}
