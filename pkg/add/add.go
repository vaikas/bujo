package add

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"

	"github.com/n3wscott/bujo/pkg/entry"
	"github.com/n3wscott/bujo/pkg/glyph"
)

type Add struct {
	Entry entry.Entry

	Bullet        glyph.Bullet
	Collection    string
	Message       string
	Priority      bool
	Inspiration   bool
	Investigation bool
}

const (
	layoutISO = "2020-01-01"
	layoutUS  = "January 1, 2020"
)

func (n *Add) Do(ctx context.Context) error {
	if n.Collection == "today" {
		n.Collection = time.Now().Format(layoutUS)
	}

	e := entry.New(n.Collection, n.Bullet, n.Message)

	switch {
	case n.Priority:
		e.Signifier = glyph.Priority
	case n.Inspiration:
		e.Signifier = glyph.Inspiration
	case n.Investigation:
		e.Signifier = glyph.Investigation
	}

	fmt.Println(glyph.Underline(glyph.Bold(e.Title())))

	tbl := uitable.New()
	tbl.Separator = " "
	tbl.AddRow(e.Row())
	_, _ = fmt.Fprintln(color.Output, tbl)

	return nil
}
