package ssos

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func WaitCommandInvocations(ctx context.Context, ssmCli *ssm.Client, commandID *string) error {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			req := ssmCli.ListCommandsRequest(&ssm.ListCommandsInput{
				CommandId: commandID,
			})
			res, err := req.Send(ctx)
			if err != nil {
				return err
			}
			if res.ListCommandsOutput.Commands[0].Status != "Pending" && res.ListCommandsOutput.Commands[0].Status != "InProgress" {
				return nil
			}
		}
	}

}

func PrintCommandOutput(ctx context.Context, ssmCli *ssm.Client, commandID *string) error {
	listCommandInvocationReq := ssmCli.ListCommandInvocationsRequest(&ssm.ListCommandInvocationsInput{
		CommandId: commandID,
		Details:   aws.Bool(true),
	})
	listCommandInvocationRes, err := listCommandInvocationReq.Send(ctx)
	if err != nil {
		return err
	}
	for _, invocation := range listCommandInvocationRes.CommandInvocations {
		p := invocation.CommandPlugins[0]

		lines := strings.Split(strings.TrimSpace(*p.Output), "\n")
		for _, l := range lines {
			fmt.Printf("%s %s\n", *invocation.InstanceId, l)
		}
	}
	return nil
}
