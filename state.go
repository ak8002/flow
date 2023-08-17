package main

import (
	"fmt"

	"github.com/serverlessworkflow/sdk-go/v2/model"
)

func executeState(state *model.State) (map[string]model.Object, error) {
	// Execute the state.
	switch state.Type {
	// case model.StateTypeOperation:
	// 	return executeOperation(state)
	// case model.StateTypeEvent:
	// 	return executeEvent(state)
	// case model.StateTypeSubFlow:
	// 	return executeSubFlow(state)
	// case model.StateTypeSwitch:
	// 	return executeSwitch(state)
	// case model.StateTypeForEach:
	// 	return executeForEach(state)
	case model.StateTypeInject:
		return executeInject(state)
	// case model.StateTypeCallback:
	// 	return executeCallback(state)
	// case model.StateTypeSleep:
	// 	return executeSleep(state)
	// case model.StateTypeEventBasedSwitch:
	// 	return executeEventBasedSwitch(state)
	// case model.StateTypeDatabasedSwitch:
	// 	return executeDatabasedSwitch(state)
	// case model.StateTypeEventTimeout:
	// 	return executeEventTimeout(state)
	// case model.StateTypeDefault:
	// 	return execute
	default:
		return nil, fmt.Errorf("unknown state type %s", state.Type)
	}
}

func executeInject(state *model.State) (map[string]model.Object, error) {
	injectState := (*model.InjectState)(state.InjectState)
	return injectState.Data, nil
}
