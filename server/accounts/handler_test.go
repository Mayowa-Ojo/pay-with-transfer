package accounts

import (
	"context"
	"database/sql"
	"pay-with-transfer/cache"
	"pay-with-transfer/config"
	"pay-with-transfer/store"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFetchSingleAccount(t *testing.T) {
	ctx := context.Background()
	cfg := config.Config{}
	mockCtrl := gomock.NewController(t)
	mockStore := store.NewMockStore(mockCtrl)
	mockCache := cache.NewMockCache(mockCtrl)

	h := New(mockStore, mockCache, cfg)

	t.Run("account-exists", func(t *testing.T) {
		accountID := uuid.New()
		expected := &store.Account{
			ID: accountID,
		}
		mockStore.EXPECT().GetAccountByID(ctx, accountID.String()).Return(expected, nil).Times(1)

		result, err := h.FetchSingleAccount(ctx, accountID.String())
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected.ID.String(), result.ID.String())
	})
	t.Run("account-does-not-exist", func(t *testing.T) {
		accountID := uuid.New()
		mockStore.EXPECT().GetAccountByID(ctx, accountID.String()).Return(nil, sql.ErrNoRows).Times(1)

		result, err := h.FetchSingleAccount(ctx, accountID.String())
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}
