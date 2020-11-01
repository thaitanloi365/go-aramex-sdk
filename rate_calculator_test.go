package aramex

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaitanloi365/go-aramex-sdk/types"
)

func TestRateCalculator(t *testing.T) {
	var aramex = New(&Config{
		IsLive:     false,
		ClientInfo: DefaultClientInfo,
	})

	result, err := aramex.CalculateRate(context.Background(), RateCalculatorRequest{
		OriginAddress: &Address{
			CountryCode:         types.CountryCodeJO,
			City:                "Amman",
			StateOrProvinceCode: "Amman",
		},
		DestinationAddress: &Address{
			CountryCode:         types.CountryCodeJO,
			City:                "Kerak",
			StateOrProvinceCode: "Kerak",
		},
		ShipmentDetails: &ShipmentDetails{
			NumberOfPieces: 1,
			ActualWeight: &Weight{
				Unit:  types.WeightUnitKG,
				Value: 0.5,
			},
		},
	})

	assert.NoError(t, err)

	aramex.printJSON(&result)

}
