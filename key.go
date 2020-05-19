package ssos

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var AddKeyCommandTmpl = `u=$(getent passwd {{.user}}) && x=$(echo $u |cut -d: -f6) || exit 1
install -d -m700 -o{{.user}} ${x}/.ssh; grep '{{.publicKey}}' ${x}/.ssh/authorized_keys && exit 0
echo '{{.publicKey}}'| tee -a ${x}/.ssh/authorized_keys
`

func AddKey(ssmCli *ssm.Client, instanceIds []string, user string, publicKeyPath string) error {
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
	res, err := req.Send(context.Background())
	if err != nil {
		return err
	}
	return json.NewEncoder(os.Stdout).Encode(res)
}
