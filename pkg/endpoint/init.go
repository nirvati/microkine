package endpoint

import (
	// SQLite is the only backend (etcd endpoints are passed through unchanged).
	_ "github.com/nirvati/microkine/pkg/drivers/sqlite"
)
