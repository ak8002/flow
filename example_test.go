package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test_wf(t *testing.T) {
	workflowFolder = "testdata"
	type args struct {
		name  string
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "hello world",
			args: args{
				name: "helloworld",
			},
			want:    `{"result": "Hello World!"}`,
			wantErr: false,
		},
		{
			name: "greeting",
			args: args{
				name:  "greeting",
				input: `{"person": {"name": "John"}}`,
			},
			want:    `{"greeting": "Welcome to Serverless Workflow, John!"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input any
			if tt.args.input != "" {
				err := json.Unmarshal([]byte(tt.args.input), &input)
				if err != nil {
					t.Errorf("json.Unmarshal() error = %v", err)
					return
				}
			}

			got, err := invokeWf(tt.args.name, input)
			if (err != nil) != tt.wantErr {
				t.Errorf("startWf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotStr, err := json.Marshal(got)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
				return
			}

			equal, err := equalJSON(string(gotStr), tt.want)
			if err != nil {
				t.Errorf("equalJSON() error = %v", err)
				return
			}

			if !equal {
				t.Errorf("startWf() = %v, want %v", string(gotStr), tt.want)
			}
		})
	}
}

func equalJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
