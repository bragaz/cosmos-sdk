package rest

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/slashing/internal/types"
)

// Concrete Swagger types used to generate REST documentation. Note, these types
// are not actually used but since all queries return a generic JSON raw message,
// they enabled typed documentation.
//
//nolint:deadcode,unused
type (
	validatorSignInfo struct {
		Height int64                      `json:"height"`
		Result types.ValidatorSigningInfo `json:"result"`
	}

	validatorsSigningInfo struct {
		Height int64                        `json:"height"`
		Result []types.ValidatorSigningInfo `json:"result"`
	}

	// alias the params type for use in swagger auto gendoc
	slashParams    []types.Params
	slashingParams struct {
		Height int64       `json:"height"`
		Result slashParams `json:"result"`
	}

	postUnjail struct {
		auth.StdTx
		Msgs []types.MsgUnjail `json:"msg"`
	}
)
