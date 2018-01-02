package cmd

import (
	"fmt"
	. "github.com/newlee/tequila/viz"
	"github.com/spf13/cobra"
)

var includeCmd = &cobra.Command{
	Use:   "include",
	Short: "include dependencies of source code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		result := ParseInclude(cmd.Flag("source").Value.String())

		if cmd.Flag("findCrossRefs").Value.String() == "true" {
			crossRefs := result.FindCrossRef(MergeHeaderFunc)
			for _, cf := range crossRefs {
				fmt.Println(cf)
			}
			return
		}

		if cmd.Flag("entryPoints").Value.String() == "true" {
			entryPoints := result.EntryPoints(MergeHeaderFunc)
			for _, cf := range entryPoints {
				fmt.Println(cf)
			}
			return
		}

		if cmd.Flag("mergeHeader").Value.String() == "true" {
			result = result.MergeHeaderFile(MergeHeaderFunc)
		}

		if cmd.Flag("mergePackage").Value.String() == "true" {
			result = result.MergeHeaderFile(MergePackageFunc)
		}
		var filter = func(key string) bool {
			return key == "main.cpp" || key == "main"
		}
		result.ToDot(cmd.Flag("output").Value.String(), "/", filter)
	},
}

func init() {
	rootCmd.AddCommand(includeCmd)

	includeCmd.Flags().StringP("source", "s", "", "source code directory")
	includeCmd.Flags().StringP("output", "o", "dep.dot", "output dot file name")
	includeCmd.Flags().BoolP("findCrossRefs", "F", false, "find cross references")
	includeCmd.Flags().BoolP("entryPoints", "E", false, "list entry points")
	includeCmd.Flags().BoolP("mergeHeader", "H", false, "merge header file to same cpp file")
	includeCmd.Flags().BoolP("mergePackage", "P", false, "merge package/folder for include dependencies")
}
