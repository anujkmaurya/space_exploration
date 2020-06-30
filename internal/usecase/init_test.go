package usecase

import (
	"testing"

	"github.com/personal-work/space_exploration/internal/models"
)

func Test_initResources(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "testcase 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initResources()
			if models.DNSMap == nil {
				t.Errorf("DNSMap not init")
			}
			if models.SectorsMap == nil {
				t.Errorf("SectorsMap not init")
			}
			if models.DronesMap == nil {
				t.Errorf("DronesMap not init")
			}
		})
	}
}
