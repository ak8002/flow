package main

import (
	"fmt"

	"github.com/serverlessworkflow/sdk-go/v2/model"
	"github.com/serverlessworkflow/sdk-go/v2/parser"
)

type workflow model.Workflow

func parseWf(filePath string) (*workflow, error) {
	wf, err := parser.FromFile(filePath)
	if err != nil {
		return nil, err
	}

	return (*workflow)(wf), nil
}

func (wf workflow) findState(name string) (*model.State, error) {
	// Find the state.
	for _, state := range wf.States {
		if state.Name == name {
			return &state, nil
		}
	}

	// If we didn't find a start state, return nil.
	// Parse should have caught this.
	return nil, fmt.Errorf("state %s not found", name)
}

func executeWf(wf *workflow, input any) (any, error) {
	// Start at the start state.
	state, err := wf.findState(wf.Start.StateName)
	if err != nil {
		return nil, err
	}

	// Execute each state until we reach an end state.
	for {
		// Execute the state.
		out, err := executeState(state, input)
		if err != nil {
			return nil, err
		}

		// If we reached an end state, return the output.
		if state.End != nil {
			return out, nil
		}

		// Otherwise, find the next state.
		nextState, err := wf.findState(state.Transition.NextState)
		if err != nil {
			return nil, err
		}

		// And execute it.
		state = nextState
	}
}

var workflowFolder = "workflows"

func invokeWf(name string, input any) (any, error) {
	// Parse workflow.
	workflow, err := parseWf(fmt.Sprintf("%s/%s.sw.json", workflowFolder, name))
	if err != nil {
		return nil, fmt.Errorf("failed to parse workflow: %w", err)
	}

	// Execute workflow.
	out, err := executeWf(workflow, input)
	if err != nil {
		return nil, fmt.Errorf("failed to execute workflow: %w", err)
	}
	return out, nil
}
