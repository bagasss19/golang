package feature

import (
	"context"
	"errors"
	"fico_ar/config"
	"fico_ar/domain/ar/model"
	"fico_ar/domain/ar/repository"
	arMock "fico_ar/domain/ar/repository/mock"
	"fico_ar/domain/shared/response"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_arFeature_GetAllData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := arMock.NewMockArRepository(mockCtrl)

	type fields struct {
		config       config.EnvironmentConfig
		arRepository repository.ArRepository
	}
	type args struct {
		ctx     context.Context
		payload *model.ARFilterList
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
				arRepository: mockRepo,
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.ARFilterList{},
			},
			wantResp: response.Data{
				Items: []model.AR{{}},
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
				mockRepo.EXPECT().GetAllData(gomock.Any(), gomock.Any()).Return(model.GetAllARResponse{
					Data:      []model.AR{{}},
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
				arRepository: mockRepo,
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.ARFilterList{},
			},
			wantResp: response.Data{},
			wantErr:  true,
			mockFunc: func() {
				mockRepo.EXPECT().GetAllData(gomock.Any(), gomock.Any()).Return(model.GetAllARResponse{
					Data:      []model.AR{{}},
					TotalPage: 1,
					TotalItem: 1,
				}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := arFeature{
				config:       tt.fields.config,
				arRepository: tt.fields.arRepository,
			}

			tt.mockFunc()
			gotResp, err := ar.GetAllData(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("arFeature.GetAllData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotResp, tt.wantResp, cmpopts.EquateEmpty()) {
				t.Errorf("arFeature.GetAllData() got = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_arFeature_GetOneData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := arMock.NewMockArRepository(mockCtrl)
	type fields struct {
		config       config.EnvironmentConfig
		arRepository repository.ArRepository
	}
	type args struct {
		ctx  context.Context
		arID int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp model.AR
		wantErr  bool
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
			},
			wantResp: model.AR{},
			wantErr:  false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{}, nil)
			},
		},
		{
			name: "error case",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
			},
			wantResp: model.AR{},
			wantErr:  true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := arFeature{
				config:       tt.fields.config,
				arRepository: tt.fields.arRepository,
			}

			tt.mockFunc()
			gotResp, err := ar.GetOneData(tt.args.ctx, tt.args.arID)
			if (err != nil) != tt.wantErr {
				t.Errorf("arFeature.GetOneData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotResp, tt.wantResp, cmpopts.EquateEmpty()) {
				t.Errorf("arFeature.GetOneData() got = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_arFeature_GetAllCompanyCode(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := arMock.NewMockArRepository(mockCtrl)

	type fields struct {
		config       config.EnvironmentConfig
		arRepository repository.ArRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp []model.ARSales
		wantErr  bool
		mockFunc func()
	}{
		// TODO: Add test cases.
		{
			name: "positive case",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			wantResp: []model.ARSales{},
			wantErr:  false,
			mockFunc: func() {
				mockRepo.EXPECT().GetAllCompanyCode(gomock.Any()).Return([]model.ARSales{}, nil)
			},
		},
		{
			name: "error case",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx: context.Background(),
			},
			wantResp: []model.ARSales{},
			wantErr:  true,
			mockFunc: func() {
				mockRepo.EXPECT().GetAllCompanyCode(gomock.Any()).Return([]model.ARSales{}, errors.New("error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := arFeature{
				config:       tt.fields.config,
				arRepository: tt.fields.arRepository,
			}
			tt.mockFunc()
			gotResp, err := ar.GetAllCompanyCode(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("arFeature.GetAllCompanyCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotResp, tt.wantResp, cmpopts.EquateEmpty()) {
				t.Errorf("arFeature.GetAllCompanyCode() got = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_arFeature_DeleteData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := arMock.NewMockArRepository(mockCtrl)

	type fields struct {
		config       config.EnvironmentConfig
		arRepository repository.ArRepository
	}
	type args struct {
		ctx  context.Context
		arID int64
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
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
			},
			wantErr: false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{}, nil)
				mockRepo.EXPECT().DeleteData(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "error case - data not found",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{}, errors.New("Error"))
			},
		},
		{
			name: "error case - error on delete data",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{}, nil)
				mockRepo.EXPECT().DeleteData(gomock.Any(), gomock.Any()).Return(errors.New("Error"))
			},
		},
		{
			name: "error case - status is not 0",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
			},
			wantErr: true,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{
					Status: 1,
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := arFeature{
				config:       tt.fields.config,
				arRepository: tt.fields.arRepository,
			}
			tt.mockFunc()
			if err := ar.DeleteData(tt.args.ctx, tt.args.arID); (err != nil) != tt.wantErr {
				t.Errorf("arFeature.DeleteData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_arFeature_UpdateData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := arMock.NewMockArRepository(mockCtrl)

	type fields struct {
		config       config.EnvironmentConfig
		arRepository repository.ArRepository
	}
	type args struct {
		ctx     context.Context
		request model.ARUpdatePayload
		arID    int64
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
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
				request: model.ARUpdatePayload{
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
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{Status: 0}, nil)
				mockRepo.EXPECT().UpdateData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)
			},
		},
		{
			name: "error case - error on get one data",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
				request: model.ARUpdatePayload{
					Data: map[string]interface{}{
						"giro_amount": float64(10500.00),
						"giro_num":    int64(123),
					},
				},
			},
			wantErr:  true,
			wantResp: false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{}, errors.New("error"))
			},
		},
		{
			name: "error case - error on update data",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
				request: model.ARUpdatePayload{
					Data: map[string]interface{}{
						"giro_amount": float64(10500.00),
						"giro_num":    int64(123),
					},
				},
			},
			wantErr:  true,
			wantResp: false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{Status: 0}, nil)
				mockRepo.EXPECT().UpdateData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(false, errors.New("error"))
			},
		},
		{
			name: "error case - error on data status is not 0",
			fields: fields{
				arRepository: mockRepo,
			},
			args: args{
				ctx:  context.Background(),
				arID: 1,
				request: model.ARUpdatePayload{
					Data: map[string]interface{}{
						"giro_amount": float64(10500.00),
						"giro_num":    int64(123),
					},
				},
			},
			wantErr:  true,
			wantResp: false,
			mockFunc: func() {
				mockRepo.EXPECT().GetOneData(gomock.Any(), gomock.Any()).Return(model.AR{Status: 1}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := arFeature{
				config:       tt.fields.config,
				arRepository: tt.fields.arRepository,
			}
			tt.mockFunc()

			gotResp, err := ar.UpdateData(tt.args.ctx, tt.args.request, tt.args.arID)
			if (err != nil) != tt.wantErr {
				t.Errorf("arFeature.UpdateData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp != tt.wantResp {
				t.Errorf("arFeature.UpdateData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
