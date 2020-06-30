package usecase

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/personal-work/space_exploration/internal/models"
)

func TestUsecase_GetDroneLocation(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecaseWrapper := NewMockUsecaseInterface(mockCtrl)

	u := &Usecase{
		UsecaseInterface: mockUsecaseWrapper,
	}

	//init map
	models.DronesMap = make(map[uint64]*models.Drone)
	models.DNSMap = make(map[uint64]*models.DNS)

	//prepare data insert value in map
	models.DNSMap[1] = &models.DNS{
		ID:         1,
		Name:       "1",
		SectorList: []uint64{1},
	}
	models.DronesMap[1] = &models.Drone{
		ID:       1,
		Name:     "d1",
		Type:     "v1",
		SectorID: 1,
	}
	type args struct {
		droneLocReq *models.LocationReq
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test case 1: Positive case - normal case",
			args: args{
				droneLocReq: &models.LocationReq{
					X:        123.12,
					Y:        456.56,
					Z:        789.89,
					Vel:      20.0,
					IsCustom: false,
					DroneID:  1,
					DnsID:    1,
				},
			},
			wantErr: false,
			want: models.Location{
				Loc: 1389.5700000000002,
			},
		},
		{
			name: "Test case 2: Positive case- friend company case",
			args: args{
				droneLocReq: &models.LocationReq{
					X:        123.12,
					Y:        456.56,
					Z:        789.89,
					Vel:      20.0,
					IsCustom: true,
					DroneID:  1,
					DnsID:    1,
				},
			},
			wantErr: false,
			want: models.LocationCustom{
				Loc: 1389.5700000000002,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := u.GetDroneLocation(tt.args.droneLocReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetDroneLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetDroneLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
