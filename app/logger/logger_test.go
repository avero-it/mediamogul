package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGroup(t *testing.T) {
	type args struct {
		group string
	}
	tests := []struct {
		name string
		args args
		want *logrus.Entry
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Group(tt.args.group)
		})
	}
}
