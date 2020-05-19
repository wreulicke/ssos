package main

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/spf13/cobra"
	"github.com/wreulicke/ssos"
	"github.com/wreulicke/ssos/utils"
)

func NewCreateUserCommand(profile, region *string) *cobra.Command {
	var instanceIds []string
	var username string
	cmd := &cobra.Command{
		Use:   "create-user",
		Short: "create-user",
		RunE: func(cmd *cobra.Command, args []string) error {
			if username == "" {
				return errors.New("username is required")
			}
			if len(instanceIds) == 0 {
				return errors.New("instance-id is not specified")
			}
			c, err := utils.ReadAwsConfig(*profile, *region)
			if err != nil {
				return err
			}
			return ssos.CreateUser(ssm.New(c), instanceIds, username)
		},
	}
	cmd.Flags().StringSliceVarP(&instanceIds, "instance-ids", "i", []string{}, "instance ids")
	cmd.Flags().StringVarP(&username, "username", "u", "", "username")
	return cmd
}
