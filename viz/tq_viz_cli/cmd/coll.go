package cmd

import (
	//"fmt"
	. "github.com/newlee/tequila/viz"
	"github.com/spf13/cobra"
)

var collCmd *cobra.Command = &cobra.Command{
	Use:   "coll",
	Short: "full collaboration grpahh",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		result := ParseColl(cmd.Flag("source").Value.String())

		if cmd.Flag("mergePackage").Value.String() == "true" {
			result = result.MergeHeaderFile(MergePackageFunc)
		}
		result.ToDot(cmd.Flag("output").Value.String(), ".")
	},
}

func init() {
	rootCmd.AddCommand(collCmd)

	collCmd.Flags().StringP("source", "s", "", "source code directory")
	collCmd.Flags().StringP("output", "o", "dep.dot", "output dot file name")
	collCmd.Flags().BoolP("mergePackage", "P", false, "merge package/folder for include dependencies")
}
