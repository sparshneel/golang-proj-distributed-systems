package repository

type GetRequest struct {
	RecordIds []string
	SFilters []*SimpleFilter
    CFilters []*ComplexFilter
	Aggregations []*ColumnAggregation
}

type SimpleFilter struct{
	Column string
	Value string
	FilterFunc string
}

type ComplexFilter struct{}

type ColumnAggregation struct{}

