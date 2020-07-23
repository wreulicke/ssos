package ssos

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var createUserTmpl = `sudo useradd %s || exit 1
sudo usermod -aG adm,wheel %s
echo "%s ALL=(ALL) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/%s
`

func CreateUser(ssmCli *ssm.Client, instanceIds []string, user string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	req := ssmCli.SendCommandRequest(&ssm.SendCommandInput{
		InstanceIds:  instanceIds,
		DocumentName: aws.String("AWS-RunShellScript"),
		Comment:      aws.String("create user"),
		Parameters: map[string][]string{
			"commands": {fmt.Sprintf(createUserTmpl, user, user, user, user)},
		},
	})
	res, err := req.Send(ctx)
	if err != nil {
		return err
	}
	commandID := res.Command.CommandId
	err = WaitCommandInvocations(ctx, ssmCli, commandID)
	if err != nil {
		return err
	}
	return PrintCommandOutput(ctx, ssmCli, commandID)
}
