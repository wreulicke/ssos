package utils

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)

type Logger struct {
}

func (*Logger) Log(args ...interface{}) {
	log.Println(args...)
}

func LoggerConfig() aws.Config {
	return aws.Config{
		Logger: &Logger{},
	}
}

func ReadAwsConfig(profile, region string) (aws.Config, error) {
	if profile != "" {
		return external.LoadDefaultAWSConfig(LoggerConfig(), external.WithSharedConfigProfile(profile), external.WithRegion(region))
	}
	return external.LoadDefaultAWSConfig(LoggerConfig(), external.WithRegion(region))
}
