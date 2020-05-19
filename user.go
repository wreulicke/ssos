package ssos

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var createUserTmpl = `sudo useradd %s
sudo usermod -aG adm,wheel %s
`

func CreateUser(ssmCli *ssm.Client, instanceIds []string, user string) error {
	req := ssmCli.SendCommandRequest(&ssm.SendCommandInput{
		InstanceIds:  instanceIds,
		DocumentName: aws.String("AWS-RunShellScript"),
		Comment:      aws.String("create user"),
		Parameters: map[string][]string{
			"commands": []string{fmt.Sprintf(createUserTmpl, user, user)},
		},
	})
	res, err := req.Send(context.Background())
	if err != nil {
		return err
	}
	return json.NewEncoder(os.Stdout).Encode(res)
}
