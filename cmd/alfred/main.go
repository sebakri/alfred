package main

import "github.com/sebakri/alfred/internal"

type Alfred struct {
	// define the fields of the Alfred struct here
}

func (a *Alfred) NewTask(name, description string, callback func(cmd *CommandBuilder, inputs internal.Inputs) *Command) *internal.Task {
	// implement the NewTask method here
	return &internal.Task{}
}

type ExecutionContext struct {
	Inputs  *internal.Inputs  // Add Inputs field
	Outputs *internal.Outputs // Add Outputs field
}

func (ctx *ExecutionContext) WithContainer(image string, args ...string) *internal.Task {
	return &internal.Task{}
}

type CommandBuilder struct {
}

func (c *CommandBuilder) WithContainer(image string, args ...string) *CommandBuilder {
	return c
}

func (c *CommandBuilder) OnSuccess(callback func(outputs *internal.Outputs)) *CommandBuilder {
	return c
}

func (c *CommandBuilder) OnFailure(callback func(outputs *internal.Outputs)) *CommandBuilder {
	return c
}

func (c *CommandBuilder) Build() *Command {
	return &Command{}
}

type Command struct {
}

func main() {
	alfred := new(Alfred)

	bufLint := alfred.NewTask("buf-lint", "Lint proto files", func(cmd *CommandBuilder, inputs internal.Inputs) *Command {
		return cmd.WithContainer("bufbuild/buf", "lint", "--path", inputs.Get("protos_path")).OnSuccess(func(outputs *internal.Outputs) {
			outputs.Set("lint", "success")
		}).OnFailure(func(outputs *internal.Outputs) {
			outputs.Set("lint", "failure")
		}).Build()
	})

	bufLint.Inputs.Set("protos_path", "protos")
}
