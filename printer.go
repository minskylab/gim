package gim

type Printer interface {
	ShowLoading(title string)
	HideLoading()
}
