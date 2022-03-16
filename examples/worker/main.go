// @@@SNIPSTART hello-world-project-template-go-worker
package main

import (
	"fmt"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

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
	fmt.Println("Created client for worker.")
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, app.GreetingTaskQueue, worker.Options{
		// WorkerStopTimeout: 90,
	})
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreeting)

	fmt.Println("Start listening to Task Queue...")

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

// @@@SNIPEND
