package util

import "testing"

func TestDeeSeeChiffre(t *testing.T) {
	type args struct {
		input string
		key   int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Clark",
			args: args{
				input: "clark",
				key:   5,
			},
			want:    "hqfwp",
			wantErr: false,
		},
		{
			name: "cherry blossom",
			args: args{
				input: "cherry blossom",
				key:   3,
			},
			want:    "fkhuub eorvvrp",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeeSeeChiffre(tt.args.input, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeeSeeChiffre() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeeSeeChiffre() got = %v, want %v", got, tt.want)
			}
		})
	}
}
