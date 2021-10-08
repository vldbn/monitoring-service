package response

import "monitoring-service/internal/app/model/entity"

// AuthLoginRes JWT tokens response struct
type AuthLoginRes struct {
	DefaultRes
	Username string         `json:"username,omitempty"`
	Tokens   *entity.Tokens `json:"tokens,omitempty"`
}
