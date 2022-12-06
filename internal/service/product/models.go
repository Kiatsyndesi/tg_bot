package product

var allProducts = []Product{
	{Title: "iPhone"},
	{Title: "Google Pixel"},
	{Title: "Macbook"},
	{Title: "Chromebook"},
	{Title: "Banana"},
}

type Product struct {
	Title string
}
