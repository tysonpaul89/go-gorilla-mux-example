package models

// Book Model
type Book struct {
	ID     string  `json:"id"` // <-- Struct tags. Its represent what to display when its encoded as json. {"id": "as-12"}
	Title  string  `json:"title"`
	Author *Author `json:"author"`
	Price  float64 `json:"price"`
}

// Author model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
