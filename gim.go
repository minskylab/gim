package gim

import "errors"

type Gim struct {
	Workspace *workspace
	Config *config

	Router Router
	Builder Builder
	Printer Printer
}

func NewGim(w *workspace, c *config, router Router, builder Builder, printer Printer) (*Gim, error) {
	if w == nil || c == nil {
		return nil, errors.New("the workspace and config must have a correct values")
	}
	if w.Config != c {
		w.Config = c
	}

	return &Gim{
		Workspace: w,
		Config:    c,
		Router:    router,
		Builder:   builder,
		Printer:   printer,
	}, nil
}