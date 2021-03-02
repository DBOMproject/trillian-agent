package restapi

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"trillian-agent/models"
	"trillian-agent/restapi/operations"
	"trillian-agent/test"
	client "trillian-agent/trillian"

	tclient "github.com/google/trillian/client"

	"github.com/go-openapi/loads"
	"github.com/google/trillian"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
)

// Test_main contains tests for the agent config logger
func TestGetRecord(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/test-record", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestGetRecordChannelError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/error-channel/records/test-record", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestGetRecordChannelNotFound(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/random-channel/records/test-record", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestGetRecordError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/error-record", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
func TestGetRecordNotFound(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/random-record", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestGetRecordGetRevError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionErrorMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/test-record", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
func TestGetRecordGetClientError(t *testing.T) {
	getChannelClient = getChannelClientErrorMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/test-record", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestGetRecordGetClientError2(t *testing.T) {
	getChannelClient = getChannelClientError2Mock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/test-record", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestAuditRecord(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/test-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestAuditRecordChannelError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/error-channel/records/test-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestAuditRecordChannelNotFound(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/random-channel/records/test-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestAuditRecordError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/error-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
func TestAuditRecordNotFound(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/random-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestAuditRecord2Error(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock2
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/error-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
func TestAuditRecordNotFound2(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock2
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/random-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestAuditRecordGetRevError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionErrorMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/test-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
func TestAuditRecordGetClientError(t *testing.T) {
	getChannelClient = getChannelClientErrorMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/test-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestAuditRecordGetClientError2(t *testing.T) {
	getChannelClient = getChannelClientError2Mock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	req, err := http.NewRequest("GET", "/channels/test-channel/records/test-record/audit", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func getChannelClientMock(ctx context.Context, trillAdminClient trillian.TrillianAdminClient, trillMapClient trillian.TrillianMapClient, channelMapID int64, tracer opentracing.Tracer) (*tclient.MapClient, error) {

	return &tclient.MapClient{Conn: test.NewTrillianMapMockClient(nil, false, false, false)}, nil
}

func getChannelClientErrorMock(ctx context.Context, trillAdminClient trillian.TrillianAdminClient, trillMapClient trillian.TrillianMapClient, channelMapID int64, tracer opentracing.Tracer) (*tclient.MapClient, error) {
	if channelMapID == 1536 {
		return nil, errors.New("test-error")
	}
	return &tclient.MapClient{Conn: test.NewTrillianMapMockClient(nil, false, false, false)}, nil
}

func getChannelClientError2Mock(ctx context.Context, trillAdminClient trillian.TrillianAdminClient, trillMapClient trillian.TrillianMapClient, channelMapID int64, tracer opentracing.Tracer) (*tclient.MapClient, error) {
	if channelMapID != 1536 {
		return nil, errors.New("test-error")
	}
	return &tclient.MapClient{Conn: test.NewTrillianMapMockClient(nil, false, false, false)}, nil
}

func getCurrentRevisionMock(c *client.MapClient, ctx context.Context, mapID int64, tracer opentracing.Tracer) (uint64, error) {
	return uint64(1654), nil
}

func getCurrentRevisionErrorMock(c *client.MapClient, ctx context.Context, mapID int64, tracer opentracing.Tracer) (uint64, error) {
	return 0, errors.New("test-error")
}

func GetChannelMock(ctx context.Context, client *client.MapClient, channelID string, tracer opentracing.Tracer) (*models.Channel, error) {
	if channelID == "test-channel" {
		return &models.Channel{ChannelID: "test-channel", MapID: 1536}, nil
	} else if channelID == "error-channel" {
		return nil, errors.New("test-error")
	}
	return nil, nil
}

func GetRecordMock(ctx context.Context, client *client.MapClient, recordID string, revision int64, tracer opentracing.Tracer) (*models.Record, error) {
	payload := map[string]interface{}{
		"test": "test",
	}
	payload2 := map[string]interface{}{
		"test": "test2",
	}
	if recordID == "test-record" {
		if revision == 1 {
			return &models.Record{Revision: 1, PreviousRevision: 0, AuditDefinition: models.AuditDefinition{Payload: payload}}, nil
		}
		return &models.Record{Revision: 2, PreviousRevision: 1, AuditDefinition: models.AuditDefinition{Payload: payload2}}, nil
	} else if recordID == "error-record" {
		return nil, errors.New("test-error")
	}
	return nil, nil
}

func GetRecordMock2(ctx context.Context, client *client.MapClient, recordID string, revision int64, tracer opentracing.Tracer) (*models.Record, error) {
	payload := map[string]interface{}{
		"test": "test",
	}
	payload2 := map[string]interface{}{
		"test": "test2",
	}
	if recordID == "test-record" {
		if revision == 1 {
			return &models.Record{Revision: 1, PreviousRevision: 0, AuditDefinition: models.AuditDefinition{Payload: payload}}, nil
		}
		return &models.Record{Revision: 2, PreviousRevision: 1, AuditDefinition: models.AuditDefinition{Payload: payload2}}, nil
	} else if recordID == "error-record" {
		if revision != 1 {
			return &models.Record{Revision: 2, PreviousRevision: 1, AuditDefinition: models.AuditDefinition{Payload: payload}}, nil
		}
		return nil, errors.New("test-error")
	}
	if revision != 1 {
		return &models.Record{Revision: 2, PreviousRevision: 1, AuditDefinition: models.AuditDefinition{Payload: payload}}, nil
	}
	return nil, nil
}
