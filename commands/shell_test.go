package commands

import (
	"testing"
)


func TestAWSCommand(t *testing.T) {
	RunShellCommand("aws eks --region us-east-1 update-kubeconfig --name xxxxxxxxxxx")
}