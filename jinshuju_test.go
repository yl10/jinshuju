package jinshuju

import (
	"testing"
)

func TestJSJ_GetFormInfo(t *testing.T) {
	type args struct {
		formid string
	}
	tests := []struct {
		name    string
		j       JSJ
		args    args
		want    *Form
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "测试获取表单定义",
			j:    JSJ{"xI-7LBVbemj_xK9HNUo9sg", "Jq0HUkF4WgH5VvhLjz_TDw"},
			args: args{"XuHJKg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetFormInfo(tt.args.formid)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSJ.GetFormInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("form:\r\n", *got)
		})
	}
}
