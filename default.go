package gim


func NewDefaultGim() (*Gim, error){
	router := NewGinRouter()
	builder := NewParcelBuilder()
	printer := NewSpinnerPrinter()

	config := NewDefaultConfig()

	workspace := NewWorkspace(".", config)

	return NewGim(workspace, router, builder, printer)
}