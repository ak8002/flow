package main

import (
	"fmt"

	"github.com/serverlessworkflow/sdk-go/v2/model"
	"github.com/serverlessworkflow/sdk-go/v2/parser"
)

type workflow model.Workflow

func (wf workflow) findState(name string) (*model.State, error) {
	// Find the state.
	for _, state := range wf.States {
		if state.Name == name {
			return &state, nil
		}
	}

	// If we didn't find a start state, return nil.
	return nil, fmt.Errorf("state %s not found", name)
}

func parseWf(filePath string) (*workflow, error) {
	wf, err := parser.FromFile(filePath)
	if err != nil {
		return nil, err
	}

	return (*workflow)(wf), nil
}

func executeWf(wf *workflow) (any, error) {
	// Start at a start state.
	state, err := wf.findState(wf.Start.StateName)
	if err != nil {
		return nil, err
	}

	// Execute each state until we reach an end state.
	for {
		// Execute the state.
		out, err := executeState(state)
		if err != nil {
			return nil, err
		}

		// If we reached an end state, return the output.
		if out != nil {
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

func startWf(name string) (any, error) {
	// Parse workflow.
	workflow, err := parseWf(fmt.Sprintf("workflows/%s.sw.json", name))
	if err != nil {
		return nil, err
	}

	// Execute workflow.
	return executeWf(workflow)
}
