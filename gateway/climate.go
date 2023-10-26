package gateway

import (
	"bigtable_api/entity"
	"context"
)

type ClimateGateway interface {
	ReadPrefix(ctx context.Context, table, prefix string, filters map[string]string) ([]entity.BigtableOutput, error)
	ReadRows(ctx context.Context, table string, areas, dates []string, filters map[string]string) ([]entity.BigtableOutput, error)
}
