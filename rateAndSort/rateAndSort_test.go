package rateAndSort

import (
	"testing"
)

func Test_writeEvaluationFile(t *testing.T) {
	t.Skip() // fix!!

	environment = "test"
	envFileName = "test.env"

	type args struct {
		data []stock
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "write sample file",
			args: args{data: []stock{
				{Symbol: "BAYN.de", Value: -5},
				{Symbol: "DTE.DE", Value: 5},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeEvaluationFile(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("writeEvaluationFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
