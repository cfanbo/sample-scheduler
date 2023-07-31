package main

import (
	"os"

	"github.com/cfanbo/sample/pkg/plugin"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(plugin.Name, plugin.New),
	)

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
