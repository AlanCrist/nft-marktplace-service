package cli_test

import (
	"fmt"
	"testing"

	"github.com/alancrist/nft-marktplace-service/adapters/cli"
	"github.com/alancrist/nft-marktplace-service/application"
	mock_application "github.com/alancrist/nft-marktplace-service/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product :=
		application.Product{
			ID:     "abc",
			Name:   "Product Test",
			Price:  25.99,
			Status: "enabled",
		}

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetID().Return(product.ID).AnyTimes()
	productMock.EXPECT().GetName().Return(product.Name).AnyTimes()
	productMock.EXPECT().GetStatus().Return(product.Status).AnyTimes()
	productMock.EXPECT().GetPrice().Return(product.Price).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)

	service.EXPECT().Create(product.Name, product.Price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(product.ID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s)",
		product.ID, product.Name, product.Price, product.Status)

	result, err := cli.Run(service, "create", "", product.Name, product.Price)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s ha been enabled.)",
		product.Name)

	result, err = cli.Run(service, "enable", product.ID, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s ha been disabled.)",
		product.Name)

	result, err = cli.Run(service, "disable", product.ID, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		product.ID, product.Name, product.Price, product.Status)

	result, err = cli.Run(service, "get", product.ID, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
