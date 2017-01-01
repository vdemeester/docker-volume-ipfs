package version

import "github.com/spf13/cobra"

var (
	// Cmd can be added to other commands to provide a version subcommand with
	// the correct version of swarm.
	Cmd = &cobra.Command{
		Use:   "version",
		Short: "Print version number of docker-volume-ipfs",
		Run: func(cmd *cobra.Command, args []string) {
			PrintVersion()
		},
	}
)
