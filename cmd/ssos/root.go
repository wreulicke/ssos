package main

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssos",
		Short: "ssos",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	var profile string
	var region string
	cmd.PersistentFlags().StringVarP(&profile, "profile", "p", "", "aws profile")
	cmd.PersistentFlags().StringVarP(&region, "region", "r", "", "aws region")
	cmd.AddCommand(NewCreateUserCommand(&profile, &region), NewAddKeyCommand(&profile, &region))
	return cmd
}
