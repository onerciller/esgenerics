package esgenerics

// FieldExtractor is a function type that extracts a field (key) from an entity.
type FieldExtractor[E Entity[M], M any] func(E) any

// Entity is an interface that must be implemented by all entities.
// It is used to convert the entity to a model.
type Entity[M any] interface {
	ToModel() M
}

// Result is the result of a search query.
type Result[E Entity[M], M any] struct {
	Hits Hits[E, M] `json:"hits"`
}

// Hits is the list of hits of a search query.
type Hits[E Entity[M], M any] struct {
	Hits  []Hit[E, M] `json:"hits"`
	Total Total       `json:"total"`
}

// ToFieldList converts the hits to a list of fields.
func (h Hits[E, M]) ToFieldList(extractor FieldExtractor[E, M]) []any {
	var fields []any
	for _, hit := range h.Hits {
		fields = append(fields, extractor(hit.Source))
	}
	return fields
}

// ToModelList converts the hits to a list of models.
func (h Hits[E, M]) ToModelList() []M {
	var models []M
	for _, hit := range h.Hits {
		models = append(models, hit.Source.ToModel())
	}
	return models
}

// LastItem returns the last item in the hits as a model.
// It returns a boolean indicating whether a last item was found.
func (h Hits[E, M]) LastItem() (M, bool) {
	hitsCount := len(h.Hits)
	if hitsCount == 0 {
		var zeroValue M
		return zeroValue, false
	}

	lastHit := h.Hits[hitsCount-1]
	return lastHit.Source.ToModel(), true
}

// Hit is a single hit of a search query.
type Hit[E Entity[M], M any] struct {
	Source E `json:"_source"`
}

// Total is the total number of hits of a search query.
type Total struct {
	Value int64 `json:"value"`
}
