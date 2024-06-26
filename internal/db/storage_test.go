package db

import (
	"Avito-Project/internal/db/mocks"
	"Avito-Project/internal/models"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataBase_GetUserByID(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    *models.User
		wantErr bool
	}{
		{
			name:    "User exists",
			id:      1,
			want:    &models.User{Id: 1, Name: "John Doe"},
			wantErr: false,
		},
		{
			name:    "User does not exist",
			id:      2,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(mocks.Storage)

			if tt.wantErr {
				mockStorage.On("GetUserByID", tt.id).
					Return(nil, assert.AnError)
			} else {
				mockStorage.On("GetUserByID", tt.id).
					Return(tt.want, nil)
			}

			got, err := mockStorage.GetUserByID(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
			mockStorage.AssertExpectations(t)
		})

	}
}

func TestDataBase_GetUserByToken(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		want    *models.User
		wantErr bool
	}{
		{
			name:    "Valid-token",
			token:   "Valid-token",
			want:    &models.User{Id: 1, Name: "Oleg", Token: "Valid-token"},
			wantErr: false,
		},
		{
			name:    "Invalid-token",
			token:   "Invalid-token",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(mocks.Storage)

			if tt.wantErr {
				mockStorage.On("GetUserByToken", tt.token).
					Return(nil, errors.New("user not found"))
			} else {
				mockStorage.On("GetUserByToken", tt.token).
					Return(tt.want, nil)
			}

			got, err := mockStorage.GetUserByToken(tt.token)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)

			mockStorage.AssertExpectations(t)
		})
	}
}

func TestDataBase_GetBanner(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    *models.Banner
		wantErr bool
	}{
		{
			name:    "Banner exist",
			id:      1,
			want:    &models.Banner{Id: 1, Title: "Car", Text: "Sale Car"},
			wantErr: false,
		},
		{
			name:    "Banner does not exist",
			id:      2,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(mocks.Storage)

			if tt.wantErr {
				mockStorage.On("GetBanner", tt.id).
					Return(nil, assert.AnError)
			} else {
				mockStorage.On("GetBanner", tt.id).
					Return(tt.want, nil)
			}

			got, err := mockStorage.GetBanner(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)

			mockStorage.AssertExpectations(t)
		})
	}
}

func TestDataBase_GetUsersPaginated(t *testing.T) {
	tests := []struct {
		name    string
		page    int
		size    int
		want    []models.User
		wantErr bool
	}{
		{
			name: "Valid page and size",
			page: 1,
			size: 2,
			want: []models.User{
				{Id: 1, Name: "Name1", AccessLevels: 1, Token: "token1"},
				{Id: 2, Name: "Name2", AccessLevels: 2, Token: "token2"},
			},
			wantErr: false,
		},
		{
			name:    "No users on page",
			page:    100,
			size:    2,
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Database error",
			page:    1,
			size:    2,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(mocks.Storage)

			if tt.wantErr {
				mockStorage.On("GetUsersPaginated", tt.page, tt.size).
					Return(nil, assert.AnError)
			} else {
				mockStorage.On("GetUsersPaginated", tt.page, tt.size).
					Return(tt.want, nil)
			}

			got, err := mockStorage.GetUsersPaginated(tt.page, tt.size)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)

			mockStorage.AssertExpectations(t)
		})
	}
}

func TestDataBase_GetBannersPaginated(t *testing.T) {
	tests := []struct {
		name    string
		page    int
		size    int
		want    []models.Banner
		wantErr bool
	}{
		{
			name: "Valid page and size",
			page: 1,
			size: 2,
			want: []models.Banner{
				{Id: 1, Title: "Title1", Text: "Text1", OwnerId: 1},
				{Id: 2, Title: "Title2", Text: "Text2", OwnerId: 2},
			},
			wantErr: false,
		},
		{
			name:    "No banners on page",
			page:    100,
			size:    2,
			want:    nil,
			wantErr: false,
		},
		{
			name:    "Database error",
			page:    1,
			size:    2,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(mocks.Storage)

			if tt.wantErr {
				mockStorage.On("GetBannersPaginated", tt.page, tt.size).
					Return(nil, assert.AnError)
			} else {
				mockStorage.On("GetBannersPaginated", tt.page, tt.size).
					Return(tt.want, nil)
			}

			got, err := mockStorage.GetBannersPaginated(tt.page, tt.size)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)

			mockStorage.AssertExpectations(t)
		})
	}
}
