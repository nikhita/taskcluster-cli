package taskstatus

import (
	"fmt"
	"os"

	"github.com/taskcluster/taskcluster-cli/extpoints"
	"github.com/taskcluster/taskcluster-client-go"
	"github.com/taskcluster/taskcluster-client-go/queue"
)

func init() {
	extpoints.Register("task-status", taskstatus{})
}

type taskstatus struct{}

func (taskstatus) ConfigOptions() map[string]extpoints.ConfigOption {
	return nil
}

func (taskstatus) Summary() string {
	return "Check status of task from taskId."
}

func usage() string {
	return `Usage:
  taskcluster task-status <taskId>

This command returns the task status structure from given taskId.
`
}

func (taskstatus) Usage() string {
	return usage()
}

func (taskstatus) Execute(context extpoints.Context) bool {
	argv := context.Arguments
	taskId := argv["<taskId>"].(string)

	if argv["task-status"].(bool) {
		fmt.Printf("%s\n", taskStatus(taskId))
		return true
	}
	return true
}

func taskStatus(taskId string) string {
	q := queue.New(&tcclient.Credentials{})
	q.Authenticate = false

	resp, err := q.Status(taskId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error obtaining task status: %s\n", err)
	}

	// TODO: Format this to pretty print the status
	foo := resp.Status
	return foo.State
}
