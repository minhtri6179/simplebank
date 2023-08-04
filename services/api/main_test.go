package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/minhtri67/simplebank/db/sqlc"
	"github.com/minhtri67/simplebank/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store sqlc.Store) *Server {
	config := util.Config{
		TokenSymmetric: util.RandomString(32),
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
