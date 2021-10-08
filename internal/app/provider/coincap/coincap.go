package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"monitoring-service/internal/app/model/entity"
	"monitoring-service/internal/app/provider"
	"net/http"
)

// RestCoinCapAPI implements provider.CryptocurrencyAPI interface
type RestCoinCapAPI struct {
	client http.Client
	url    string
}

// GetRates implements provider.CryptocurrencyAPI interface method
func (r *RestCoinCapAPI) GetRates(currencyID string) (*entity.Cryptocurrency, error) {
	var cur *entity.Cryptocurrency
	url := fmt.Sprintf("%s/%s", r.url, currencyID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		defer func() {
			if err := res.Body.Close(); err != nil {
				log.Println(err)
			}
		}()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return nil, errors.New(string(body))
	}
	restRep := RestCoinCapRes{}
	if err := json.NewDecoder(res.Body).Decode(&restRep); err != nil {
		return nil, err
	}
	if restRep.Data == nil {
		return nil, errors.New("unable to get rates")
	}
	cur = restRep.Data
	return cur, nil
}

// NewRestCoinCapAPI constructor
func NewRestCoinCapAPI() provider.CryptocurrencyAPI {
	return &RestCoinCapAPI{
		client: http.Client{},
		url:    "https://api.coincap.io/v2/rates",
	}
}
