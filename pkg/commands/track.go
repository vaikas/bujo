package commands

import (
	"context"
	"errors"
	"github.com/n3wscott/bujo/pkg/commands/options"
	"github.com/n3wscott/bujo/pkg/runner/track"
	"github.com/n3wscott/bujo/pkg/store"
	"github.com/spf13/cobra"
	"strings"
)

func addTrack(topLevel *cobra.Command) {
	co := &options.CollectionOptions{}

	cmd := &cobra.Command{
		Use:   "track",
		Short: "track something",
		Example: `
bujo track <thing>
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a collection")
			}
			co.Collection = strings.Join(args, " ")

			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			p, err := store.Load(nil)
			if err != nil {
				return err
			}
			s := track.Track{
				Collection:  co.Collection,
				Persistence: p,
			}
			err = s.Do(context.Background())
			return output.HandleError(err)
		},
	}

	topLevel.AddCommand(cmd)
}
