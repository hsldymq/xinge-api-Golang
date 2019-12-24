package xinge

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalPushResponse(t *testing.T) {
	data := `{"seq":0,"push_id":"3634417576","ret_code":0,"environment":"product","err_msg":"","result":"[0]"}`
	resp := PushResponse{}
	require.NoError(t, json.Unmarshal([]byte(data), &resp))
}
