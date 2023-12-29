package esgenerics

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MathAllQueryHandler struct{}

func (MathAllQueryHandler) BuildQuery(ctx context.Context) QueryMap {
	return map[string]interface{}{"match_all": map[string]interface{}{}}
}

type ResponseStorageModel struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type StorageEntity struct {
	DisplayName string `json:"displayName"`
	StorageCode string `json:"storageCode"`
}

func (s StorageEntity) ToModel() ResponseStorageModel {
	return ResponseStorageModel{
		Name: s.DisplayName,
		Code: s.StorageCode,
	}
}

type ErrorReader struct{}

func (ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func TestExecuteSearch_Success(t *testing.T) {
	// Arrange
	mockClient := NewMockElasticSearchClient(t)
	store := NewStore[StorageEntity, ResponseStorageModel](mockClient, "test-index")

	// Create a mock response
	mockResponse := &Result[StorageEntity, ResponseStorageModel]{}
	jsonResponse, _ := json.Marshal(mockResponse)
	mockHTTPResponse := &esapi.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(string(jsonResponse))),
	}

	mockClient.On("Search", mock.Anything, "test-index", mock.Anything).Return(mockHTTPResponse, nil)

	// Act
	result, err := store.ExecuteSearch(context.Background(), MathAllQueryHandler{})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, mockResponse, result)
	mockClient.AssertExpectations(t)
}

func TestExecuteSearch_ElasticsearchError(t *testing.T) {
	// Arrange
	mockClient := NewMockElasticSearchClient(t)
	store := NewStore[StorageEntity, ResponseStorageModel](mockClient, "test-index")

	expectedErrorMessage := "elasticsearch search error"
	mockClient.On("Search", mock.Anything, "test-index", mock.Anything).Return(nil, errors.New(expectedErrorMessage))

	// Act
	result, err := store.ExecuteSearch(context.Background(), MathAllQueryHandler{})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), expectedErrorMessage)
	mockClient.AssertExpectations(t)
}

func TestExecuteSearch_JSONUnmarshalError(t *testing.T) {
	// Arrange
	mockClient := NewMockElasticSearchClient(t)
	store := NewStore[StorageEntity, ResponseStorageModel](mockClient, "test-index")

	// Create a response with invalid JSON
	mockHTTPResponse := &esapi.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("{invalid-json")),
	}

	mockClient.On("Search", mock.Anything, "test-index", mock.Anything).Return(mockHTTPResponse, nil)

	// Act
	result, err := store.ExecuteSearch(context.Background(), MathAllQueryHandler{})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	mockClient.AssertExpectations(t)
}

func TestExecuteSearch_ReadError(t *testing.T) {
	// Arrange
	mockClient := NewMockElasticSearchClient(t)
	store := NewStore[StorageEntity, ResponseStorageModel](mockClient, "test-index")

	mockHTTPResponse := &esapi.Response{
		StatusCode: 200,
		Body:       io.NopCloser(ErrorReader{}),
	}

	mockClient.On("Search", mock.Anything, "test-index", mock.Anything).Return(mockHTTPResponse, nil)

	// Act
	result, err := store.ExecuteSearch(context.Background(), MathAllQueryHandler{})

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error reading response body")
	mockClient.AssertExpectations(t)
}
