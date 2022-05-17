package usecase

import (
	"context"
	"golang_api/model"
	"reflect"
	"testing"
)

func Test_userUsecase_GetUserByID(t *testing.T) {
	type args struct {
		ctx    context.Context
		UserID uint32
	}
	tests := []struct {
		name     string
		u        *userUsecase
		args     args
		wantData *model.User
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := tt.u.GetUserByID(tt.args.ctx, tt.args.UserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("userUsecase.GetUserByID() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
