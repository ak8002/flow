package main

import (
	"reflect"
	"testing"

	"github.com/serverlessworkflow/sdk-go/v2/model"
)

func Test_wf(t *testing.T) {
	workflowFolder = "testdata"
	type args struct {
		name  string
		input model.Object
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "hello world",
			args: args{
				name: "helloworld",
			},
			want:    map[string]model.Object{"result": {Type: model.String, StrVal: "Hello World!"}},
			wantErr: false,
		},
		{
			name: "greeting",
			args: args{
				name: "greeting",
				input: model.Object{
					Type:     model.Raw,
					RawValue: []byte(`{"person": {"name": "John"}}`),
				},
			},
			want:    map[string]model.Object{"greeting": {Type: model.String, StrVal: "Welcome to Serverless Workflow, John!"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := invokeWf(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("startWf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("startWf() = %v, want %v", got, tt.want)
			}
		})
	}
}
