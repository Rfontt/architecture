package application_test

import (
	mock_application "github.com/Rfontt/go-hexagonal/application/mocks"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)


func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistense.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistense: persistence
	}

	result, err := service.Get("abc")

	require.Nil(t, err)
	require.Equal(t, product, result)
}