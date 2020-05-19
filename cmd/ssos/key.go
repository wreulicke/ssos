package main

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/spf13/cobra"
	"github.com/wreulicke/ssos"
	"github.com/wreulicke/ssos/utils"
)

func NewAddKeyCommand(profile, region *string) *cobra.Command {
	var instanceIds []string
	var username string
	var publicKeyPath string
	cmd := &cobra.Command{
		Use:   "add-ssh-key",
		Short: "add-ssh-key",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := utils.ReadAwsConfig(*profile, *region)
			if err != nil {
				return err
			}

			return ssos.AddKey(ssm.New(c), instanceIds, username, publicKeyPath)
		},
	}
	cmd.Flags().StringSliceVarP(&instanceIds, "instance-ids", "i", []string{}, "instance ids")
	cmd.Flags().StringVarP(&username, "username", "u", "", "username")
	cmd.Flags().StringVarP(&publicKeyPath, "public-key", "k", "", "")
	return cmd
}
