package gim

import (
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

type SpinnerPrinter struct {
	s *spinner.Spinner
}

func NewSpinnerPrinter() *SpinnerPrinter {
	return &SpinnerPrinter{
		s: spinner.New(spinner.CharSets[11], 100*time.Millisecond),
	}
}


func (sp *SpinnerPrinter) ShowLoading(title string) {
	sp.s.Suffix = " " + strings.TrimSpace(title)
	sp.s.Start()
}

func (sp *SpinnerPrinter) HideLoading() {
	sp.s.Stop()
}