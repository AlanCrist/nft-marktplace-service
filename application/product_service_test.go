package application_test

import (
	"testing"

	"github.com/alancrist/nft-marktplace-service/application"
	mock_application "github.com/alancrist/nft-marktplace-service/application/mocks"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

func TestProductServiceGetByID(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("12345")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
