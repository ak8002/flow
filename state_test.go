package main

import (
	"reflect"
	"testing"

	"github.com/serverlessworkflow/sdk-go/v2/model"
)

func Test_executeInject(t *testing.T) {
	type args struct {
		state *model.State
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]model.Object
		wantErr bool
	}{
		{
			name: "returns static data",
			args: args{
				state: &model.State{
					InjectState: &model.InjectState{
						Data: map[string]model.Object{"message": {Type: model.String, StrVal: "Hello, World!"}},
					},
				},
			},
			want:    map[string]model.Object{"message": {Type: model.String, StrVal: "Hello, World!"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executeInject(tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeInject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeInject() = %v, want %v", got, tt.want)
			}
		})
	}
}
