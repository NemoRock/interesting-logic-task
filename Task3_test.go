package main

import "testing"

func Test_experiment(t *testing.T) {
	type args struct {
		heightMax    int
		heightBroken int
	}
	tests := []struct {
		name                  string
		args                  args
		wantCounterExperiment int
		wantHeightCast        int
	}{
		{
			name:                  "1",
			args:                  args{heightMax: 5000, heightBroken: 1},
			wantCounterExperiment: 2,
			wantHeightCast:        1,
		},
		{
			name:                  "3",
			args:                  args{heightMax: 5000, heightBroken: 3},
			wantCounterExperiment: 4,
			wantHeightCast:        3,
		},
		{
			name:                  "50",
			args:                  args{heightMax: 5000, heightBroken: 50},
			wantCounterExperiment: 51,
			wantHeightCast:        50,
		},
		{
			name:                  "100",
			args:                  args{heightMax: 5000, heightBroken: 100},
			wantCounterExperiment: 101,
			wantHeightCast:        100,
		},
		{
			name:                  "101",
			args:                  args{heightMax: 5000, heightBroken: 101},
			wantCounterExperiment: 3,
			wantHeightCast:        101,
		},
		{
			name:                  "5000",
			args:                  args{heightMax: 5000, heightBroken: 5000},
			wantCounterExperiment: 96,
			wantHeightCast:        5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCounterExperiment, gotHeightCast := experiment(tt.args.heightMax, tt.args.heightBroken)
			if gotCounterExperiment != tt.wantCounterExperiment {
				t.Errorf("experiment() gotCounterExperiment = %v, want %v", gotCounterExperiment, tt.wantCounterExperiment)
			}
			if gotHeightCast != tt.wantHeightCast {
				t.Errorf("experiment() gotHeightCast = %v, want %v", gotHeightCast, tt.wantHeightCast)
			}
		})
	}
}
