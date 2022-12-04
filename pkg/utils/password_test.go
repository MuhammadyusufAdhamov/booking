package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPassword(t *testing.T) {
	password := "12345678"

	hashedPassword, err := HashPassword(password)
	fmt.Println(hashedPassword)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}
