package handler

import (
	"bytes"
	"encoding/json"
	"go-test-basic/common"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Initialize
	common.SetEnv()
	m.Run()
	// Cleanup
	common.TeardownEnv()
}

func TestCreateItem(t *testing.T) {
	t.Log("Before ApplyFuncReturn")
	gomonkey.ApplyFuncReturn(common.GetConfig, common.Config{MaxLength: 10})
	t.Log("After ApplyFuncReturn")

	// gomonkey.ApplyFunc(common.GetConfig, func() common.Config {
	// 	return common.Config{MaxLength: 10}
	// })

	t.Run("name to long", func(t *testing.T) {
		common.InitTestDBTwo(t)

		// Create request and ResponseWriter
		req := CreateRequest{Name: "ttttttttttttttttttttttttttttttt", Description: "test desc"}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Marshal request
		reqBytes, err := json.Marshal(req)
		require.NoError(t, err)
		c.Request, _ = http.NewRequest("POST", "/create", bytes.NewReader(reqBytes))

		CreateItem(c)
		require.Equal(t, 400, w.Code)
	})

	t.Run("success", func(t *testing.T) {
		common.InitTestDBTwo(t)

		// Create request and ResponseWriter
		req := CreateRequest{Name: "test", Description: "test desc"}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Marshal request
		reqBytes, err := json.Marshal(req)
		require.NoError(t, err)
		c.Request, _ = http.NewRequest("POST", "/create", bytes.NewReader(reqBytes))
		CreateItem(c)

		// Check response
		require.Equal(t, 200, w.Code)
		var res CreateResponse
		err = json.Unmarshal(w.Body.Bytes(), &res)
		require.NoError(t, err)

		// require.Equal(t, 1, res.ID)
	})
}
