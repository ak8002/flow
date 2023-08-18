package main

import (
	"reflect"
	"testing"

	"github.com/serverlessworkflow/sdk-go/v2/model"
)

func Test_invokeWf(t *testing.T) {
	workflowFolder = "testdata"
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "returns error if workflow cannot be parsed",
			args: args{
				name: "notfound",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns error if workflow is invalid",
			args: args{
				name: "invalid",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns successful result",
			args: args{
				name: "helloworld",
			},
			want:    map[string]model.Object{"result": {Type: model.String, StrVal: "Hello World!"}},
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
