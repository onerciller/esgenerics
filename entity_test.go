package esgenerics

import (
	"reflect"
	"testing"
)

func TestHitToModel(t *testing.T) {
	entity := StorageEntity{DisplayName: "Test Entity"}
	hit := Hit[StorageEntity, ResponseStorageModel]{Source: entity}

	model := hit.Source.ToModel()
	expected := ResponseStorageModel{Name: "Test Entity"}

	if !reflect.DeepEqual(model, expected) {
		t.Errorf("Hit ToModel was incorrect, got: %v, want: %v.", model, expected)
	}
}

func TestHitsToModels(t *testing.T) {
	entities := []Hit[StorageEntity, ResponseStorageModel]{
		{Source: StorageEntity{DisplayName: "Entity 1"}},
		{Source: StorageEntity{DisplayName: "Entity 2"}},
	}
	hits := Hits[StorageEntity, ResponseStorageModel]{Hits: entities}

	var models []ResponseStorageModel
	for _, hit := range hits.Hits {
		models = append(models, hit.Source.ToModel())
	}

	expected := []ResponseStorageModel{
		{Name: "Entity 1"},
		{Name: "Entity 2"},
	}

	if !reflect.DeepEqual(models, expected) {
		t.Errorf("Hits ToModels was incorrect, got: %v, want: %v.", models, expected)
	}
}
