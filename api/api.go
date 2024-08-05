package api

import (
	"encoding/json"
	"net/http"
)

// Structs
type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type Error struct {
	Code int
	Message string
}