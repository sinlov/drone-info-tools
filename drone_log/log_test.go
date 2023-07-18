package drone_log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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