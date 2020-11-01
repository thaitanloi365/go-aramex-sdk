package aramex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaitanloi365/go-aramex-sdk/types"
)

func TestLocationFetchCities(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})

	result, err := aramex.FetchCities(context.Background(), &CitiesFetchingRequest{
		CountryCode: types.CountryCodeSG,
	})

	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestLocationFetchStates(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})

	result, err := aramex.FetchStates(context.Background(), &StatesFetchingRequest{
		CountryCode: types.CountryCodeSG,
	})

	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestLocationFetchOffices(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})

	result, err := aramex.FetchOffices(context.Background(), &OfficesFetchingRequest{
		CountryCode: types.CountryCodeSG,
	})

	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestLocationFetchCountries(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})

	result, err := aramex.FetchCountries(context.Background(), nil)

	assert.NoError(t, err)

	aramex.printJSON(&result)
}

func TestLocationValidateAddress(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})

	result, err := aramex.ValidateAddress(context.Background(), &AddressValidationRequest{
		Address: &Address{
			CountryCode:         types.CountryCodeJO,
			City:                "Kerak",
			StateOrProvinceCode: "Kerak",
		},
	})

	assert.NoError(t, err)

	aramex.printJSON(&result)
}
