package filelog

import (
	"testing"
	"time"
)

func TestFileController_timeTag(t *testing.T) {
	type fields struct {
		expire time.Duration
		dir    string
		file   string
		period LogPeriod
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "base",
			fields: fields{
				expire: 0,
				dir:    "worker",
				file:   "error",
				period: PeriodHour,
			},
			args: args{
				t: time.Now(),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &FileController{
				expire: tt.fields.expire,
				dir:    tt.fields.dir,
				file:   tt.fields.file,
				period: tt.fields.period,
			}
			if got := w.timeTag(tt.args.t); got != tt.want {
				t.Errorf("timeTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
