package main

import "testing"

func TestDo(t *testing.T) {
	type args struct {
		s string
		i int
		b bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"0,0,true,invalid", args{"0", 0, false}, "", true},
		{"1,1,true,invalid", args{"1", 1, false}, "", true},
		{"d,8,false,valid", args{"d", 8, false}, "[8d]", false},
		{"z,13,true,valid", args{"z", 13, true}, "", true},
		{"c,34,true,valid", args{"c", 34, true}, "C", false},
		{"a,1,false,valid", args{"a", 1, false}, "[1a]", false},
		{"d,8,true,valid", args{"d", 8, true}, "[8D]", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Do(tt.args.s, tt.args.i, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
