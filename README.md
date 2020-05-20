
# ssos ( ssh-setup-over-ssm )

You can setup your own user and ssh keys in EC2 instance over AWS Systems Manager, even if there is no ssh private key for EC2 instances.
You don't need private key for EC2 instance forever.

This is inspired by [ssh-over-ssm](https://github.com/elpy1/ssh-over-ssm), Thanks @elpy1.

## Usage

```bash
$ ssos
ssos

Usage:
  ssos [flags]
  ssos [command]

Available Commands:
  add-ssh-key add-ssh-key
  create-user create-user
  help        Help about any command

Flags:
  -h, --help             help for ssos
  -p, --profile string   aws profile
  -r, --region string    aws region

Use "ssos [command] --help" for more information about a command.
```

## Disclaimer

I have just checked this only in Amazon Linux2.

## TODO

* Release to GitHub Release
* Wait termination of command executed via SSM
* Allow the user created to use sudo