package repository

import (
	"bigtable_api/entity"
	"context"
	"errors"
	"log"
	"strconv"

	"cloud.google.com/go/bigtable"
)

type ClimateRepository struct {
	ClientInstance *bigtable.Client
}

func NewClimateRepository(clientInstance *bigtable.Client) *ClimateRepository {
	return &ClimateRepository{ClientInstance: clientInstance}
}

func (r *ClimateRepository) ReadPrefix(ctx context.Context, table, prefix string, filters map[string]string) ([]entity.BigtableOutput, error) {
	log.Printf("Reading from table %s with prefix %s", table, prefix)

	tbl := r.ClientInstance.Open(table)

	var filterList []bigtable.Filter

	if regexp, ok := filters["regexp"]; ok {
		filterList = append(filterList, bigtable.RowKeyFilter(regexp))
	}

	if version, ok := filters["version"]; ok {
		versionInt, err := strconv.Atoi(version)
		if err != nil {
			return nil, errors.New("wrong version filter")
		}
		filterList = append(filterList, bigtable.LatestNFilter(versionInt))
	}

	// filters 1 version by default
	filter := bigtable.LatestNFilter(1)

	if len(filterList) == 1 {
		filter = filterList[0]
	} else if len(filterList) > 1 {
		filter = bigtable.ChainFilters(filterList...)
	}

	var result []entity.BigtableOutput
	err := tbl.ReadRows(ctx, bigtable.PrefixRange(prefix),
		func(row bigtable.Row) bool {
			for _, cols := range row {
				for _, col := range cols {
					output := entity.BigtableOutput{
						Key:     row.Key(),
						Created: col.Timestamp.Time().UTC(),
						Value:   string(col.Value),
					}
					result = append(result, output)
				}
			}
			return true
		}, bigtable.RowFilter(filter))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ClimateRepository) ReadRows(ctx context.Context, table string, areas, dates []string, filters map[string]string) ([]entity.BigtableOutput, error) {
	log.Printf("Reading from table %s with areas: %s and dates: %s", table, areas, dates)

	tbl := r.ClientInstance.Open(table)

	filter, err := getFilter(filters)
	if err != nil {
		return nil, err
	}

	var output []entity.BigtableOutput
	if len(dates) > 1 {
		output, err = readRowRange(ctx, dates, areas, filter, tbl)
	} else {
		output, err = readNoRange(ctx, dates[0], areas, filter, tbl)
	}
	if err != nil {
		return nil, err
	}
	return output, nil
}

func readNoRange(ctx context.Context, date string, areas []string, filter bigtable.Filter, tbl *bigtable.Table) ([]entity.BigtableOutput, error) {
	var rowList bigtable.RowList
	for _, area := range areas {
		rowList = append(rowList, "w/"+area+"/"+date)
	}

	var result []entity.BigtableOutput
	err := tbl.ReadRows(ctx, rowList,
		func(row bigtable.Row) bool {
			for _, cols := range row {
				for _, col := range cols {
					output := entity.BigtableOutput{
						Key:     row.Key(),
						Created: col.Timestamp.Time().UTC(),
						Value:   string(col.Value),
					}
					result = append(result, output)
				}
			}
			return true
		}, bigtable.RowFilter(filter))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func readRowRange(ctx context.Context, dates, areas []string, filter bigtable.Filter, tbl *bigtable.Table) ([]entity.BigtableOutput, error) {
	var rowRangeList bigtable.RowRangeList
	for _, area := range areas {
		rowRangeList = append(rowRangeList, bigtable.NewRange("w/"+area+"/"+dates[0], "w/"+area+"/"+dates[1]))
	}

	var result []entity.BigtableOutput
	err := tbl.ReadRows(ctx, rowRangeList,
		func(row bigtable.Row) bool {
			for _, cols := range row {
				for _, col := range cols {
					output := entity.BigtableOutput{
						Key:     row.Key(),
						Created: col.Timestamp.Time().UTC(),
						Value:   string(col.Value),
					}
					result = append(result, output)
				}
			}
			return true
		}, bigtable.RowFilter(filter))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getFilter(filters map[string]string) (bigtable.Filter, error) {
	var filterList []bigtable.Filter

	if version, ok := filters["version"]; ok {
		versionInt, err := strconv.Atoi(version)
		if err != nil {
			return nil, errors.New("wrong version filter")
		}
		filterList = append(filterList, bigtable.LatestNFilter(versionInt))
	} else {
		// filters 1 version by default
		filterList = append(filterList, bigtable.LatestNFilter(1))
	}

	if regexp, ok := filters["regexp"]; ok {
		filterList = append(filterList, bigtable.RowKeyFilter(regexp))
	}

	var filter bigtable.Filter
	if len(filterList) == 1 {
		filter = filterList[0]
	} else if len(filterList) > 1 {
		filter = bigtable.ChainFilters(filterList...)
	}
	return filter, nil
}
