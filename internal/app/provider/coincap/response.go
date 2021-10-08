package coincap

import "monitoring-service/internal/app/model/entity"

// RestCoinCapRes CoinCap HTTP response
type RestCoinCapRes struct {
	Data *entity.Cryptocurrency `json:"data,omitempty"`
}
