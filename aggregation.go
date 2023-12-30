package esgenerics

// AggregationResponse is a generic struct for different types of aggregations.
type AggregationResponse struct {
	Aggregations map[string]AggregationResult `json:"aggregations"`
}

func (a *AggregationResponse) GetAggregationResult(name string) (AggregationResult, bool) {
	agg, ok := a.Aggregations[name]
	return agg, ok
}

// AggregationResult is a generic struct for a terms aggregation.
type AggregationResult struct {
	Buckets []AggregationBucket
}

// ToKeyCountMap converts the buckets in a terms aggregation to a map of key to count.
func (t *AggregationResult) ToKeyCountMap() (map[string]int, bool) {
	result := make(map[string]int)
	if len(t.Buckets) == 0 {
		return result, false
	}
	for _, bucket := range t.Buckets {
		result[bucket.Key] = bucket.Count
	}
	return result, true
}

// AggregationBucket is a generic struct for a bucket in a terms aggregation.
type AggregationBucket struct {
	Key   string `json:"key"`
	Count int    `json:"doc_count"`
}
