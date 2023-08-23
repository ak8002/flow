package main

import (
	"encoding/json"
	"fmt"

	"github.com/itchyny/gojq"
)

func applyExpression(expression string, input any) (any, error) {
	query, err := gojq.Parse(expression)
	if err != nil {
		return nil, fmt.Errorf("error parsing jq expression %s: %w", expression, err)
	}

	//query.Run(v) doesn't take input of type map[string]string
	//so we need to convert it to map[string]interface{}
	input, err = removeTypes(input)
	if err != nil {
		return nil, err
	}

	iter := query.Run(input)
	var v any
	for {
		next, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := next.(error); ok {
			return nil, fmt.Errorf("error running jq expression %s: %w", expression, err)
		}
		v = next
	}

	return v, nil
}

func removeTypes(v any) (any, error) {
	str, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var clean map[string]interface{}
	err = json.Unmarshal(str, &clean)
	if err != nil {
		return nil, err
	}

	return clean, nil
}
