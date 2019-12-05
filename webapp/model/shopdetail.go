package model

type ShopDetail struct {
	ImageURL         string
	URL              string
	Name             string
	DescriptionShort string
	PricePerLiter    float32
}

var shopDetails []ShopDetail = []ShopDetail{
	{
		ImageURL:         "lemon.png",
		URL:              "shop_detail.html",
		Name:             "Lemon Juice",
		DescriptionShort: "Made from fresh, organic California lemons.",
		PricePerLiter:    1.09,
	}, {
		ImageURL:         "apple.png",
		URL:              "..",
		Name:             "Apple Juice",
		DescriptionShort: "The perfect blend of Washington apples",
		PricePerLiter:    0.89,
	},
	{
		ImageURL:         "watermelon.png",
		URL:              "..",
		Name:             "Watermelon Juice",
		DescriptionShort: "From sun-drenched fields in Florida.",
		PricePerLiter:    3.99,
	},
	{
		ImageURL:         "kiwi.png",
		URL:              "..",
		Name:             "Kiwi Juice",
		DescriptionShort: "California sunshine and rain distilled into sweet goodness.",
		PricePerLiter:    5.99,
	}, {
		ImageURL:         "mangosteen.png",
		URL:              "..",
		Name:             "Mangosteen Juice",
		DescriptionShort: "Our latest taste sensation, imported directly from Hawaii",
		PricePerLiter:    6.87,
	}, {
		ImageURL:         "orange.png",
		URL:              "..",
		Name:             "Orange Juice",
		DescriptionShort: "Fresh squeezed from Florida's best oranges.",
		PricePerLiter:    1.67,
	}, {
		ImageURL:         "pineapple.png",
		URL:              "..",
		Name:             "Pineapple Juice",
		DescriptionShort: "An exotic and refreshing offering. Straight from Hawaii.",
		PricePerLiter:    2.48,
	}, {
		ImageURL:         "strawberry.png",
		URL:              "..",
		Name:             "Strawberry Juice",
		DescriptionShort: "The perfect balance of sweet and tart.",
		PricePerLiter:    4.36,
	},
}
