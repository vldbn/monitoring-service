package coincapmock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockCoinCapAPI_GetRates(t *testing.T) {
	t.Run("get-rates", func(t *testing.T) {
		api := NewMockCoinCapAPI()
		cur, err := api.GetRates("bitcoin")
		assert.NoError(t, err)
		assert.NotNil(t, cur)
		assert.Equal(t, "bitcoin", cur.CurrencyId)
	})
}
