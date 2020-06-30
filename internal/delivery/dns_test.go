package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/personal-work/space_exploration/internal/models"
	"github.com/personal-work/space_exploration/internal/usecase"
)

func TestDelivery_GetDroneLocation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsecaseWrapper := usecase.NewMockUsecaseInterface(mockCtrl)

	d := &Delivery{
		UsecaseLayer: mockUsecaseWrapper,
	}
	type args struct {
		Orgurl, url, method string
		dnsID, droneID      uint64
		isCustom            bool
		routeVariable       string
		payload             string
		locReq              *models.LocationReq
	}

	r := mux.NewRouter()

	tests := []struct {
		name                     string
		args                     args
		wantErr                  bool
		wantHTTPResponse         int
		want                     map[string]string
		mockGetDroneLocationFunc func(locReq *models.LocationReq)
	}{
		{
			name: "Test case 1: Positive case - normal case",
			args: args{
				url:      "/dns/%d/drones/%d/location?is_custom=%s",
				Orgurl:   "/dns/{dnsID}/drones/{id}/location",
				method:   "POST",
				dnsID:    1,
				droneID:  1,
				isCustom: false,
				payload: `{
					"x": 123.12,
					"y": 456.56,
					"z": 789.89,
					"vel": 20.0
					}`,
				locReq: &models.LocationReq{
					X:        123.12,
					Y:        456.56,
					Z:        789.89,
					Vel:      20.0,
					IsCustom: false,
					DroneID:  1,
					DnsID:    1,
				},
			},
			mockGetDroneLocationFunc: func(locReq *models.LocationReq) {
				mockUsecaseWrapper.EXPECT().GetDroneLocation(gomock.Any()).Return(
					models.Location{
						Loc: 1389.00,
					}, nil)
			},
			wantHTTPResponse: http.StatusOK,
			wantErr:          false,
			want: map[string]string{
				"loc": "1389",
			},
		},
		{
			name: "Test case 2: Positive case- friend company case",
			args: args{
				url:      "/dns/%d/drones/%d/location?is_custom=%s",
				Orgurl:   "/dns/{dnsID}/drones/{id}/location",
				method:   "POST",
				dnsID:    1,
				droneID:  1,
				isCustom: true,
				payload: `{
					"x": 123.12,
					"y": 456.56,
					"z": 789.89,
					"vel": 20.0
					}`,
				locReq: &models.LocationReq{
					X:        123.12,
					Y:        456.56,
					Z:        789.89,
					Vel:      20.0,
					IsCustom: true,
					DroneID:  1,
					DnsID:    1,
				},
			},
			mockGetDroneLocationFunc: func(locReq *models.LocationReq) {
				mockUsecaseWrapper.EXPECT().GetDroneLocation(gomock.Any()).Return(
					models.LocationCustom{
						Loc: 1389.00,
					}, nil)
			},
			wantHTTPResponse: http.StatusOK,
			wantErr:          false,
			want: map[string]string{
				"location": "1389",
			},
		},
		{
			name: "Test case 3: Negative case - parse drone id failure",
			args: args{
				url:      "/dns/%d/drones/$%d/location?is_custom=%s",
				Orgurl:   "/dns/{dnsID}/drones/{id}/location",
				method:   "POST",
				dnsID:    1,
				droneID:  1,
				isCustom: true,
				payload: `{
					"x": 123.12,
					"y": 456.56,
					"z": 789.89,
					"vel": 20.0
					}`,
				locReq: &models.LocationReq{},
			},
			mockGetDroneLocationFunc: func(locReq *models.LocationReq) {

			},
			wantHTTPResponse: http.StatusBadRequest,
			wantErr:          true,
			want:             nil,
		},
		{
			name: "Test case 3: Negative case - parse dnsID failure",
			args: args{
				url:      "/dns/$%d/drones/%d/location?is_custom=%s",
				Orgurl:   "/dns/{dnsID}/drones/{id}/location",
				method:   "POST",
				dnsID:    1,
				droneID:  1,
				isCustom: true,
				payload: `{
					"x": 123.12,
					"y": 456.56,
					"z": 789.89,
					"vel": 20.0
					}`,
				locReq: &models.LocationReq{},
			},
			mockGetDroneLocationFunc: func(locReq *models.LocationReq) {

			},
			wantHTTPResponse: http.StatusBadRequest,
			wantErr:          true,
			want:             nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlPath := fmt.Sprintf(tt.args.url, tt.args.dnsID, tt.args.droneID, tt.args.isCustom)
			req, err := http.NewRequest(tt.args.method, urlPath, strings.NewReader(tt.args.payload))
			if err != nil {
				t.Errorf("Error in creating  NewRequest error = %v\n", err)
				return
			}

			tt.mockGetDroneLocationFunc(tt.args.locReq)

			r.Handle(tt.args.Orgurl, HandlerFunc(d.GetDroneLocation))
			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantHTTPResponse {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantHTTPResponse)
			}
			if rr.Code == http.StatusOK {

				response := &models.Response{}
				json.NewDecoder(rr.Body).Decode(&response)

				if got := fmt.Sprintf("%v", response.Data); got != fmt.Sprintf("%v", tt.want) {
					t.Errorf("handler returned wrong response, got %v want %v",
						got, fmt.Sprintf("%v", tt.want))
				}
			}

		})
	}
}
