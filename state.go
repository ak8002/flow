package main

import (
	"encoding/json"
	"fmt"

	"github.com/serverlessworkflow/sdk-go/v2/model"
)

func executeState(state *model.State, input any) (any, error) {
	// Execute the state.
	switch state.Type {
	case model.StateTypeOperation:
		return executeOperation(state, input)
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

func executeInject(state *model.State) (any, error) {
	injectState := (*model.InjectState)(state.InjectState)
	return injectState.Data, nil
}

func executeOperation(state *model.State, input any) (any, error) {
	operationState := (*model.OperationState)(state.OperationState)

	// Return error for unimplemented parallel mode
	if operationState.ActionMode == model.ActionModeParallel {
		return nil, fmt.Errorf("parallel mode not implemented")
	}

	// Execute actions in sequence
	for _, action := range operationState.Actions {
		// Execute action
		output, err := executeAction(&action, input)
		if err != nil {
			return nil, err
		}

		// Update input for next action
		input = output
	}

	return input, nil
}

func executeAction(action *model.Action, input any) (any, error) {
	var out any
	var err error
	// Execute function defined in functionRef
	if action.FunctionRef != nil {
		out, err = executeFunction(action.FunctionRef, input)
		if err != nil {
			return nil, err
		}
	}

	// Filter the action results to select only the result data that should
	// be added/merged back into the state data using its results property. Select the part of state data which
	// the action data results should be added/merged to using the toStateData property.
	if action.ActionDataFilter != (model.ActionDataFilter{}) && action.ActionDataFilter.Results != "" {
		expression := action.ActionDataFilter.Results[2 : len(action.ActionDataFilter.Results)-1]
		out, err = applyExpression(expression, out)
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

func executeFunction(functionRef *model.FunctionRef, input any) (any, error) {
	arguments := functionRef.Arguments

	eArgs := make(map[string]model.Object)
	for key, argument := range arguments {
		// if argument is an expression, evaluate it
		if argument.Type == model.String && argument.StrVal[0] == '$' {
			// strip prefix and suffix
			expression := argument.StrVal[2 : len(argument.StrVal)-1]
			v, err := applyExpression(expression, input)
			if err != nil {
				return nil, err
			}

			argument = model.FromRaw(v)
		}

		eArgs[key] = argument
	}

	// execute referenced function
	return executeFunctionByName(functionRef.RefName, eArgs)
}

func executeFunctionByName(name string, args map[string]model.Object) (any, error) {
	argsStr, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	var argsMap map[string]any
	err = json.Unmarshal(argsStr, &argsMap)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"greeting": fmt.Sprintf("Welcome to Serverless Workflow, %v!", argsMap["name"]),
	}, nil
}
