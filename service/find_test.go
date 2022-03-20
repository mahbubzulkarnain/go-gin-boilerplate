package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/gomodul/abstraction"
	"modul"
	"modul/dto"

	repository "modul/repository/gorm"
)

func Test_service_Find(t *testing.T) {
	data := make([]*modul.Entity, 11)
	p := abstraction.NewPagination()

	type fields struct {
		modulRepository modul.Repository
	}
	type args struct {
		ctx    context.Context
		filter *dto.Filter
		p      *abstraction.Pagination
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         []*modul.Entity
		want1        *abstraction.PaginationInfo
		wantErr      bool
		registerMock func(*repository.Mock, fields, args)
	}{
		{
			name:    "should success",
			fields:  fields{},
			args:    args{},
			want:    data,
			want1:   nil,
			wantErr: false,
			registerMock: func(mock *repository.Mock, fields fields, args args) {
				mock.On("Find", args.ctx, args.filter, args.p).Return(data, nil)
			},
		},
		{
			name:   "should success with pagination",
			fields: fields{},
			args: args{
				p: p,
			},
			want: data[:*p.PageSize],
			want1: &abstraction.PaginationInfo{
				Pagination: p,
				More:       true,
			},
			wantErr: false,
			registerMock: func(mock *repository.Mock, fields fields, args args) {
				mock.On("Find", args.ctx, args.filter, args.p).Return(data, nil)
			},
		},
		{
			name:    "should error",
			fields:  fields{},
			args:    args{},
			want:    nil,
			want1:   nil,
			wantErr: true,
			registerMock: func(mock *repository.Mock, fields fields, args args) {
				mock.On("Find", args.ctx, args.filter, args.p).
					Return([]*modul.Entity{}, errors.New("something wrong"))
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
			got, got1, err := s.Find(tt.args.ctx, tt.args.filter, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
