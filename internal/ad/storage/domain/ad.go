package domain

type Ad struct {
	Id          string   `json:"-"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Price       int      `json:"price"`
	DateAdded   string   `json:"-" db:"date_added"`
	MainPhoto   string   `json:"main_photo" db:"main_photo"`
	Photos      []string `json:"photos,omitempty"`
}
