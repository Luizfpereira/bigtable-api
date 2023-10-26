package usecase

import (
	"bigtable_api/entity"
	"bigtable_api/gateway"
	"context"
)

type ClimateUsecase struct {
	gateway gateway.ClimateGateway
}

func NewClimateUsecase(gateway gateway.ClimateGateway) *ClimateUsecase {
	return &ClimateUsecase{gateway: gateway}
}

func (c *ClimateUsecase) ReadPrefix(ctx context.Context, table string, filters map[string]string, prefixes ...string) ([]entity.BigtableOutput, error) {
	var prefix string
	for _, prefixPart := range prefixes {
		if prefixPart != "" {
			if prefix == "" {
				prefix += prefixPart

			} else {
				prefix += "/" + prefixPart
			}
		}
	}
	output, err := c.gateway.ReadPrefix(ctx, table, prefix, filters)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c *ClimateUsecase) Read(ctx context.Context, table string, filters map[string]string, areas, dates []string) ([]entity.BigtableOutput, error) {
	output, err := c.gateway.ReadRows(ctx, table, areas, dates, filters)
	if err != nil {
		return nil, err
	}
	return output, nil
}
