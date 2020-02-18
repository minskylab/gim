package gim

import "errors"

type Gim struct {
	Workspace *Workspace

	Router Router
	Builder Builder
	Printer Printer
}

func NewGim(w *Workspace, router Router, builder Builder, printer Printer) (*Gim, error) {
	if w == nil || w.Config == nil {
		return nil, errors.New("the Workspace and Config must have a correct values")
	}

	return &Gim{
		Workspace: w,
		Router:    router,
		Builder:   builder,
		Printer:   printer,
	}, nil
}