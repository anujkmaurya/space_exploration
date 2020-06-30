package delivery

import (
	"testing"

	"github.com/personal-work/space_exploration/internal/models"

	"github.com/golang/mock/gomock"
	"github.com/personal-work/space_exploration/internal/usecase"
)

func TestDelivery_Init(t *testing.T) {
	type args struct {
		iUsecase usecase.UsecaseInterface
	}
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := usecase.NewMockUsecaseInterface(ctrl)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Testcase 1",
			args: args{
				iUsecase: m,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.iUsecase)
		})
	}
}

func TestDelivery_initResources(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Testcase 1",
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
