package signals

import "testing"

func TestInitSignals(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "all ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitSignals()
		})
	}
}
