package drone_log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShowLogLineNo(t *testing.T) {
	// mock ShowLogLineNo
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name     string
		openLine bool
		args     args
	}{
		{
			name:     "test verbose open line",
			openLine: true,
			args: args{
				format: "test verbose",
				v:      nil,
			},
		},
		{
			name:     "test verbose close line",
			openLine: false,
			args: args{
				format: "test verbose",
				v:      nil,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do ShowLogLineNo
			ShowLogLineNo(tc.openLine)

			// verify ShowLogLineNo
			Verbosef(tc.args.format, tc.args.v...)
		})
	}
}

func TestDebugf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test debug",
			args: args{
				format: "test debug",
				v:      nil,
			},
		},
	}
	OpenDebug()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugf(tt.args.format, tt.args.v...)
		})
	}
}

func TestDebugJson(t *testing.T) {
	type args struct {
		d any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test debug json",
			args: args{
				d: struct {
					Name string
					Age  int
				}{
					Name: "foo",
					Age:  18,
				},
			},
		},
	}
	OpenDebug()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DebugJson(tt.args.d)
		})
	}
}

func TestDebugJsonf(t *testing.T) {
	type args struct {
		d      any
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test debug jsonf",
			args: args{
				d: struct {
					Name string
					Age  int
				}{
					Name: "foo",
					Age:  18,
				},
				format: "test debug jsonf",
				v:      nil,
			},
		},
	}
	OpenDebug()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DebugJsonf(tt.args.d, tt.args.format, tt.args.v...)
		})
	}
}

func TestVerbosef(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test verbose",
			args: args{
				format: "test verbose",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Verbosef(tt.args.format, tt.args.v...)
		})
	}
}

func TestVerboseJson(t *testing.T) {
	type args struct {
		d any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test verbose json",
			args: args{
				d: struct {
					Name string
					Age  int
				}{
					Name: "foo",
					Age:  18,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseJson(tt.args.d)
		})
	}
}

func TestVerboseJsonf(t *testing.T) {
	type args struct {
		d      any
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test verbose jsonf",
			args: args{
				d: struct {
					Name string
					Age  int
				}{
					Name: "foo",
					Age:  18,
				},
				format: "test verbose jsonf",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerboseJsonf(tt.args.d, tt.args.format, tt.args.v...)
		})
	}
}

func TestInfof(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test info",
			args: args{
				format: "test info",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.format, tt.args.v...)
		})
	}
}

func TestWarnf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test warn",
			args: args{
				format: "test warn",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnf(tt.args.format, tt.args.v...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test error",
			args: args{
				err: fmt.Errorf("new test error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.err)
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		err    error
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test error",
			args: args{
				err:    fmt.Errorf("new test error"),
				format: "test error",
				v:      nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.err, tt.args.format, tt.args.v...)
		})
	}
}

func TestPanicf(t *testing.T) {
	// mock TestPanicName

	if !assert.Panics(t, func() {
		// do TestPanicName
		Panicf("this is Panicf: %s", "panic desc")
	}) {
		// verify TestPanicName
		t.Fatalf("TestPanicName should panic")
	}
}
