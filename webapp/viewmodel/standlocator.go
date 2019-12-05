package viewmodel

type StandLocator struct {
	Title  string
	Active string
}

func NewStandLocator() Base {
	return Base{
		Title:  "Stand Locator",
		Active: "standlocator",
	}
}
