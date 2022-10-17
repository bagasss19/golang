package feature

import (
	"context"
	"errors"
	"fico_ar/config"
	"fico_ar/domain/giro/model"
	"fico_ar/domain/giro/repository"
	giroMock "fico_ar/domain/giro/repository/mock"
	"fico_ar/domain/shared/response"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_giroFeature_GetAllData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := giroMock.NewMockGiroRepository(mockCtrl)
	type fields struct {
		config         config.EnvironmentConfig
		giroRepository repository.GiroRepository
	}
	type args struct {
		ctx     context.Context
		payload *model.GetGiroListPayload
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp response.Data
		wantErr  bool
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.GetGiroListPayload{},
			},
			wantResp: response.Data{
				Items: []model.Giro{{}},
				Pagination: response.Pagination{
					LimitPerPage: 5,
					CurrentPage:  1,
					TotalPage:    1,
					TotalRows:    1,
					TotalItems:   1,
					First:        true,
					Last:         true,
				},
			},
			wantErr: false,
			mockFunc: func() {
				mockRepo.EXPECT().GetAllData(gomock.Any(), gomock.Any()).Return(model.GetAllGiroResponse{
					Data:      []model.Giro{{}},
					TotalPage: 1,
					TotalItem: 1,
					Page:      1,
					Size:      5,
					First:     false,
					Last:      true,
				}, nil)
			},
		},
		{
			name: "error case",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.GetGiroListPayload{},
			},
			wantResp: response.Data{},
			wantErr:  true,
			mockFunc: func() {
				mockRepo.EXPECT().GetAllData(gomock.Any(), gomock.Any()).Return(model.GetAllGiroResponse{
					Data:      []model.Giro{{}},
					TotalPage: 1,
					TotalItem: 1,
					Page:      1,
					Size:      5,
					First:     false,
					Last:      true,
				}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := giroFeature{
				config:         tt.fields.config,
				giroRepository: tt.fields.giroRepository,
			}
			tt.mockFunc()
			gotResp, err := g.GetAllData(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("giroFeature.GetAllData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotResp, tt.wantResp, cmpopts.EquateEmpty()) {
				t.Errorf("giroFeature.GetAllData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_giroFeature_GetOneData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := giroMock.NewMockGiroRepository(mockCtrl)
	type fields struct {
		config         config.EnvironmentConfig
		giroRepository repository.GiroRepository
	}
	type args struct {
		ctx    context.Context
		giroID int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp model.Giro
		wantErr  bool
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:    context.Background(),
				giroID: 1,
			},
			wantResp: model.Giro{},
			wantErr:  false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{}, nil)
			},
		},
		{
			name: "error case",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:    context.Background(),
				giroID: 1,
			},
			wantResp: model.Giro{},
			wantErr:  true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := giroFeature{
				config:         tt.fields.config,
				giroRepository: tt.fields.giroRepository,
			}
			tt.mockFunc()
			gotResp, err := g.GetOneData(tt.args.ctx, tt.args.giroID)
			if (err != nil) != tt.wantErr {
				t.Errorf("giroFeature.GetOneData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotResp, tt.wantResp, cmpopts.EquateEmpty()) {
				t.Errorf("arFeature.GetOneData() got = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_giroFeature_CreateData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := giroMock.NewMockGiroRepository(mockCtrl)
	type fields struct {
		config         config.EnvironmentConfig
		giroRepository repository.GiroRepository
	}
	type args struct {
		ctx     context.Context
		request model.GiroRequest
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantGiroID int64
		wantErr    bool
		mockFunc   func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				request: model.GiroRequest{
					CompanyID:    "COMP0001",
					AccountID:    "string",
					AccountName:  "string",
					BankName:     "string",
					CreatedBy:    "string",
					CreatedTime:  "2020-12-19",
					DueDate:      "2020-12-19",
					GiroAmount:   0,
					GiroDate:     "2020-12-19",
					GiroNum:      0,
					LastUpdate:   "2020-12-19",
					ProfitCenter: "string",
					Status:       0,
					Type:         "string",
					UpdatedBy:    "string",
				},
			},
			wantGiroID: 1,
			wantErr:    false,
			mockFunc: func() {
				mockRepo.EXPECT().CreateData(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "error case",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				request: model.GiroRequest{
					CompanyID:    "COMP0001",
					AccountID:    "string",
					AccountName:  "string",
					BankName:     "string",
					CreatedBy:    "string",
					CreatedTime:  "2020-12-19",
					DueDate:      "2020-12-19",
					GiroAmount:   0,
					GiroDate:     "2020-12-19",
					GiroNum:      0,
					LastUpdate:   "2020-12-19",
					ProfitCenter: "string",
					Status:       0,
					Type:         "string",
					UpdatedBy:    "string",
				},
			},
			wantGiroID: 1,
			wantErr:    true,
			mockFunc: func() {
				mockRepo.EXPECT().CreateData(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := giroFeature{
				config:         tt.fields.config,
				giroRepository: tt.fields.giroRepository,
			}
			tt.mockFunc()
			err := g.CreateData(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("giroFeature.CreateData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_giroFeature_DeleteData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := giroMock.NewMockGiroRepository(mockCtrl)
	type fields struct {
		config         config.EnvironmentConfig
		giroRepository repository.GiroRepository
	}
	type args struct {
		ctx    context.Context
		giroID int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:    context.Background(),
				giroID: 1,
			},
			wantErr: false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{}, nil)
				mockRepo.EXPECT().DeleteData(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "error case - data not found",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:    context.Background(),
				giroID: 1,
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{}, errors.New("Error"))
			},
		},
		{
			name: "error case - error on delete data",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:    context.Background(),
				giroID: 1,
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{}, nil)
				mockRepo.EXPECT().DeleteData(gomock.Any(), gomock.Any()).Return(errors.New("Error"))
			},
		},
		{
			name: "error case - status is not 0",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:    context.Background(),
				giroID: 1,
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{
					Status: 1,
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := giroFeature{
				config:         tt.fields.config,
				giroRepository: tt.fields.giroRepository,
			}
			tt.mockFunc()
			if err := g.DeleteData(tt.args.ctx, tt.args.giroID); (err != nil) != tt.wantErr {
				t.Errorf("giroFeature.DeleteData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_giroFeature_UpdateData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := giroMock.NewMockGiroRepository(mockCtrl)
	type fields struct {
		config         config.EnvironmentConfig
		giroRepository repository.GiroRepository
	}
	type args struct {
		ctx     context.Context
		request model.GiroUpdatePayload
		giroID  int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp bool
		wantErr  bool
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:    context.Background(),
				giroID: 1,
				request: model.GiroUpdatePayload{
					Data: map[string]interface{}{
						"giro_amount":     float64(10500.00),
						"giro_num":        int64(123),
						"status":          1,
						"total_payment":   20000,
						"disc_payment":    10000,
						"cash_payment":    10000,
						"transfer_amount": 5000,
						"transfer_num":    5000,
						"cndn_amount":     2000,
						"cndn_num":        3000,
						"return_amount":   2000,
						"return_num":      1000,
					},
				},
			},
			wantErr:  false,
			wantResp: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{Status: 0}, nil)
				mockRepo.EXPECT().UpdateData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			},
		},
		// {
		// 	name: "error case - error on get one data",
		// 	fields: fields{
		// 		giroRepository: mockRepo,
		// 	},
		// 	args: args{
		// 		ctx:    context.Background(),
		// 		giroID: 1,
		// 		request: model.GiroUpdatePayload{
		// 			Data: map[string]interface{}{
		// 				"giro_amount": float64(10500.00),
		// 				"giro_num":    int64(123),
		// 			},
		// 		},
		// 	},
		// 	wantErr:  true,
		// 	wantResp: false,
		// 	mockFunc: func() {
		// 		mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{}, errors.New("error"))
		// 	},
		// },
		// {
		// 	name: "error case - error on update data",
		// 	fields: fields{
		// 		giroRepository: mockRepo,
		// 	},
		// 	args: args{
		// 		ctx:    context.Background(),
		// 		giroID: 1,
		// 		request: model.GiroUpdatePayload{
		// 			Data: map[string]interface{}{
		// 				"giro_amount": float64(10500.00),
		// 				"giro_num":    int64(123),
		// 			},
		// 		},
		// 	},
		// 	wantErr:  true,
		// 	wantResp: false,
		// 	mockFunc: func() {
		// 		mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{Status: 0}, nil)
		// 		mockRepo.EXPECT().UpdateData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(false, errors.New("error"))
		// 	},
		// },
		{
			name: "error case - error on data status is not 0",
			fields: fields{
				giroRepository: mockRepo,
			},
			args: args{
				ctx:    context.Background(),
				giroID: 1,
				request: model.GiroUpdatePayload{
					Data: map[string]interface{}{
						"giro_amount": float64(10500.00),
						"giro_num":    int64(123),
					},
				},
			},
			wantErr:  true,
			wantResp: false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.Giro{Status: 1}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := giroFeature{
				config:         tt.fields.config,
				giroRepository: tt.fields.giroRepository,
			}
			tt.mockFunc()
			gotResp, err := g.UpdateData(tt.args.ctx, tt.args.request, tt.args.giroID)
			if (err != nil) != tt.wantErr {
				t.Errorf("giroFeature.UpdateData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp != tt.wantResp {
				t.Errorf("giroFeature.UpdateData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
