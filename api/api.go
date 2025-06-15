package api

import (
	"encoding/json"
)

type (
	Request struct {
		JsonRPC string          `json:"jsonrpc"`
		Method  string          `json:"method"`
		Params  json.RawMessage `json:"params,omitempty"`
		ID      json.RawMessage `json:"id,omitempty"`
	}

	Response struct {
		JsonRPC string          `json:"jsonrpc"`
		Result  json.RawMessage `json:"result,omitempty"`
		Error   *RpcError       `json:"error,omitempty"`
		ID      json.RawMessage `json:"id,omitempty"`
	}

	RpcError struct {
		Code    int             `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data,omitempty"`
	}

	// TODO server side sent notifications, how to fit them into this?
	Handler func(*Request) *Response

	BatchRequest  []*Request
	BatchResponse []*Response
)

const (
	JsonRPC = "2.0"
)
