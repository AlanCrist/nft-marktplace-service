package application_test

import (
	"testing"

	"github.com/alancrist/nft-marktplace-service/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProductEnable(t *testing.T) {
	product := application.Product{
		Name:   "Avatar",
		Status: application.DISABLED,
		Price:  10,
	}

	err := product.Enable()
	require.Nil(t, err)
}

func TestProductEnableWithError(t *testing.T) {
	product := application.Product{
		Name:   "Avatar",
		Status: application.DISABLED,
		Price:  0,
	}

	err := product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProductDisable(t *testing.T) {
	product := application.Product{
		Name:   "Avatar",
		Status: application.ENABLED,
		Price:  0,
	}

	err := product.Disable()
	require.Nil(t, err)
}

func TestProductDisableWithError(t *testing.T) {
	product := application.Product{
		Name:   "Avatar",
		Status: application.ENABLED,
		Price:  10,
	}

	err := product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProductIsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Avatar",
		Status: application.DISABLED,
		Price:  10,
	}

	_, err := product.IsValid()

	require.Nil(t, err)
}

func TestProductIsInvalidWhenIDIsNotUUIDV4(t *testing.T) {
	product := application.Product{
		ID:     "1h2uj12k22k",
		Name:   "Avatar",
		Status: application.DISABLED,
		Price:  10,
	}

	_, err := product.IsValid()

	require.Equal(t, "ID: 1h2uj12k22k does not validate as uuidv4", err.Error())
}

func TestProductIsInvalidWhenInvalidStatus(t *testing.T) {
	product := application.Product{
		ID:     "1h2uj12k22k",
		Name:   "Avatar",
		Status: "INVALID",
		Price:  10,
	}

	_, err := product.IsValid()

	require.Equal(t, "the status must be enabled or disabled", err.Error())
}

func TestProductIsInvalidWhenPriceBellowZero(t *testing.T) {
	product := application.Product{
		ID:     "1h2uj12k22k",
		Name:   "Avatar",
		Status: application.DISABLED,
		Price:  -10,
	}

	_, err := product.IsValid()

	require.Equal(t, "the price must be greater or equal zero", err.Error())
}
