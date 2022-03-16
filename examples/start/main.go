// @@@SNIPSTART hello-world-project-template-go-start-workflow
package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	"hello-world-project-template-go/app"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{
		HostPort: "temporal-app.localhost:7233",
		ConnectionOptions: client.ConnectionOptions{
			DisableHealthCheck: true,
		},
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	fmt.Println("Created client for start")
	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: app.GreetingTaskQueue,
		// WorkflowTaskTimeout: 90,
	}
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	fmt.Println("Execute workflow")
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	fmt.Println("Got workflow result")
	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}

// @@@SNIPEND
