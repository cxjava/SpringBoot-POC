package cmd

import "testing"

func Test_concatUrl(t *testing.T) {
	type args struct {
		base string
		ref  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "get value",
			args: args{
				base: "http://www.baidu.com/",
				ref:  "/abcd",
			},
			want:    "http://www.baidu.com/abcd",
			wantErr: false,
		}, {
			name: "get value 2",
			args: args{
				base: "http://www.baidu.com",
				ref:  "/345/def/?88=99&name=x",
			},
			want:    "http://www.baidu.com/345/def/?88=99&name=x",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := concatUrl(tt.args.base, tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("concatUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("concatUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
