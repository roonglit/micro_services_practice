package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	user, err := testQueries.CreateUser(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, "yao4@example.com", user.Email)
	require.Equal(t, "test123", user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

}

func TestGetUser(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), "yao3@example.com")
	require.NoError(t, err)
	require.NotNil(t, user)
	// require.Equal(t, "riw@example.com", user.Email)

}
