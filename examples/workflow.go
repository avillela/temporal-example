// @@@SNIPSTART hello-world-project-template-go-workflow
package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		// See https://docs.temporal.io/blog/activity-timeouts/#lifecycle-of-an-activity
		// See https://docs.temporal.io/blog/activity-timeouts/#putting-it-all-together---a-recruiting-example
		StartToCloseTimeout: time.Minute * 45,
		HeartbeatTimeout:    time.Minute * 10,
		WaitForCancellation: false,
	}
	fmt.Println("**** Set the workflow options ***")
	ctx = workflow.WithActivityOptions(ctx, options)
	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	return result, err
}

// @@@SNIPEND
