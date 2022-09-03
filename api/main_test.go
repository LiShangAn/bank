package api

import (
	"fmt"
	"os"
	"testing"
	"time"

	db "github.com/LiShangAn/bank/db/sqlc"
	"github.com/LiShangAn/bank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestserver(t *testing.T, store db.IStore) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	fmt.Println("start test")
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
