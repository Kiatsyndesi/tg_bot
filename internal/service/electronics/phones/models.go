package phones

var AllPhones = []Phone{
	{ID: 1, Title: "iPhone 14 Pro"},
	{ID: 2, Title: "Google Pixel"},
	{ID: 3, Title: "Asus Republic Of Gamers"},
	{ID: 4, Title: "Samsung Galaxy S22"},
	{ID: 5, Title: "Xiaomi"},
}

type Phone struct {
	ID    uint64
	Title string
}
