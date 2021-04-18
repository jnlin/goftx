package goftx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/jnlin/goftx/models"
)

const (
	apiGetCoins             = "/wallet/coins"
	apiGetBalances          = "/wallet/balances"
	apiGetBalancesAll       = "/wallet/all_balances"
	apiGetDepositAddress    = "/wallet/deposit_address/%s"
	apiGetDepositHistory    = "/wallet/deposits"
	apiGetWithdrawalHistory = "/wallet/withdrawals"
	apiRequestWithdrawal    = apiGetWithdrawalHistory
	apiGetAirdrops          = "/wallet/airdrops"
	apiGetSavedAddresses    = "/wallet/saved_addresses"
	apiCreateSavedAddresses = apiGetSavedAddresses
	apiDeleteSavedAddresses = apiGetSavedAddresses
)

type Wallet struct {
	client *Client
}

func (w *Wallet) GetBalances() ([]*models.Balance, error) {
	request, err := a.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetBalances),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := a.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.Balance
	if err = json.Unmarshal(response, &result); err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
