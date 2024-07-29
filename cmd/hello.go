package cmd

import (
	"github.com/gkwa/birdantler/core"
	"github.com/spf13/cobra"
)

var (
	dirToWatch string
	patterns   []string
	filterType string
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Watch a directory for changes",
	Long:  `Watch a specified directory for changes and print the absolute path of changed files.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		err := core.RunWatcher(ctx, dirToWatch, patterns, filterType)
		if err != nil {
			LoggerFrom(ctx).Error(err, "Error watching directory")
		}
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringVarP(&dirToWatch, "dir", "d", ".", "Directory to watch for changes")
	helloCmd.Flags().StringSliceVarP(&patterns, "pattern", "p", []string{}, "File patterns to watch (e.g. '*.txt', '*.md')")
	helloCmd.Flags().StringVarP(&filterType, "type", "t", "write", "Filter type (create, write, remove, rename, chmod)")
}
