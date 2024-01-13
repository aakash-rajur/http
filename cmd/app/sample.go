package main

var books = []Book{
	{
		Id:          1,
		Name:        "The Alchemist",
		Description: "lorem ipsum",
	},
	{
		Id:          2,
		Name:        "The Monk Who Sold His Ferrari",
		Description: "lorem ipsum",
	},
	{
		Id:          3,
		Name:        "The Subtle Art of Not Giving a F*ck",
		Description: "lorem ipsum",
	},
	{
		Id:          4,
		Name:        "The 5 AM Club",
		Description: "lorem ipsum",
	},
	{
		Id:          5,
		Name:        "The Power of Now",
		Description: "lorem ipsum",
	},
	{
		Id:          6,
		Name:        "The Secret",
		Description: "lorem ipsum",
	},
	{
		Id:          7,
		Name:        "The 7 Habits of Highly Effective People",
		Description: "lorem ipsum",
	},
	{
		Id:          8,
		Name:        "The 4-Hour Workweek",
		Description: "lorem ipsum",
	},
	{
		Id:          9,
		Name:        "The 48 Laws of Power",
		Description: "lorem ipsum",
	},
	{
		Id:          10,
		Name:        "The 10X Rule",
		Description: "lorem ipsum",
	},
}

var users = []User{
	{
		Id:    1,
		Name:  "Aakash Rajur",
		Email: "aakashrajur@example.com",
	},
	{
		Id:    2,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	},
	{
		Id:    3,
		Name:  "Jane Doe",
		Email: "janedoe@example.com",
	},
}

type Book struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type User struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
