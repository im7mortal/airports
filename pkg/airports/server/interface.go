// Created by Petr Lozhkin

package server

import (
	"net/http"
)

type SDKServer interface {
	GetMainEngine() http.Handler
}
