package service

import (
	"context"
	"reflect"
	"testing"

	"modul"
	"modul/dto"

	repository "modul/repository/gorm"
)

func Test_service_Create(t *testing.T) {
	type fields struct {
		modulRepository modul.Repository
	}
	type args struct {
		ctx  context.Context
		args dto.CreateRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         *dto.ExecResponse
		wantErr      bool
		registerMock func(*repository.Mock, fields, args)
	}{
		{
			name:   "should success",
			fields: fields{},
			args: args{
				args: dto.CreateRequest{
					Name: "Lorem",
				},
			},
			want: &dto.ExecResponse{
				ID:           "1",
				RowsAffected: 1,
			},
			wantErr: false,
			registerMock: func(mock *repository.Mock, fields fields, args args) {
				entity := &modul.Entity{Name: args.args.Name}
				mock.On("Create", args.ctx, entity).Return(&dto.ExecResponse{ID: "1", RowsAffected: 1}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := new(repository.Mock)
			tt.registerMock(mock, tt.fields, tt.args)

			s := service{
				modulRepository: mock,
			}
			got, err := s.Create(tt.args.ctx, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
