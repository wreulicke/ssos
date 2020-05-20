package ssos

import (
	"bytes"
	"context"
	"io/ioutil"
	"text/template"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var AddKeyCommandTmpl = `u=$(getent passwd {{.user}}) && x=$(echo $u |cut -d: -f6) || exit 1
install -d -m700 -o{{.user}} ${x}/.ssh; grep '{{.publicKey}}' ${x}/.ssh/authorized_keys > /dev/null 2>/dev/null && exit 0
echo '{{.publicKey}}'| tee -a ${x}/.ssh/authorized_keys
`

func AddKey(ssmCli *ssm.Client, instanceIds []string, user string, publicKeyPath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	bs, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return err
	}
	publicKey := string(bs)
	t, err := template.New("test").Parse(AddKeyCommandTmpl)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = t.Execute(&b, map[string]string{
		"user":      user,
		"publicKey": publicKey,
	})
	if err != nil {
		return err
	}

	req := ssmCli.SendCommandRequest(&ssm.SendCommandInput{
		InstanceIds:  instanceIds,
		DocumentName: aws.String("AWS-RunShellScript"),
		Comment:      aws.String("create user"),
		Parameters: map[string][]string{
			"commands": []string{b.String()},
		},
	})
	res, err := req.Send(ctx)
	commandID := res.Command.CommandId
	err = WaitCommandInvocations(ctx, ssmCli, commandID)
	if err != nil {
		return err
	}
	return PrintCommandOutput(ctx, ssmCli, commandID)
}
