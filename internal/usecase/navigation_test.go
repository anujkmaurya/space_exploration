package usecase

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/personal-work/space_exploration/internal/models"
)

func TestUsecase_findLocationCoreLogic(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecaseWrapper := NewMockUsecaseInterface(mockCtrl)

	u := &Usecase{
		UsecaseInterface: mockUsecaseWrapper,
	}

	type args struct {
		drone *models.Drone
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test case 1: Positive case ",
			args: args{
				drone: &models.Drone{
					ID:       1,
					Name:     "drone1",
					Type:     "v1",
					SectorID: 1,
					Loc: models.Coordinates{
						X: 1,
						Y: 1,
						Z: 1,
					},
					Vel: 5.0,
				},
			},
			want: 8.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := u.findLocationCoreLogic(tt.args.drone); got != tt.want {
				t.Errorf("Usecase.findLocationCoreLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_findLocationAdvancedCoreLogic(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecaseWrapper := NewMockUsecaseInterface(mockCtrl)

	u := &Usecase{
		UsecaseInterface: mockUsecaseWrapper,
	}

	type args struct {
		drone *models.Drone
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test case 1: Positive case ",
			args: args{
				drone: &models.Drone{
					ID:       2,
					Name:     "drone2",
					Type:     "v2",
					SectorID: 1,
					Loc: models.Coordinates{
						X: 1,
						Y: 1,
						Z: 1,
					},
					Vel: 5.0,
				},
			},
			want: 108.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := u.findLocationAdvancedCoreLogic(tt.args.drone); got != tt.want {
				t.Errorf("Usecase.findLocationAdvancedCoreLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}
