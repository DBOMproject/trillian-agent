package restapi

import (
	"bytes"
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

func TestAddRecord(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

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

func TestAddRecordInvalidType(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "BAD")

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
func TestAddRecordError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record-error"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

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
func TestAddRecordGetError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "error-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

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
func TestUpdateRecordConflictError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "test-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusConflict)
	}
}

func TestAddRecordRevError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}
	channelConfigMapID = -2
	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	channelConfigMapID = -1
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestAddRecordRevError2(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}
	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel-bad-map-id/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

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

func TestAddRecordGetClientError(t *testing.T) {
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
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

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
func TestAddRecordGetClientError2(t *testing.T) {
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
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

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
func TestAddRecordChannelError(t *testing.T) {
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

	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/error-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

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
func TestAddRecordCreateChannel(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createChannel = CreateChannelMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/new-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

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
func TestAddRecordCreateChannelError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createChannel = CreateChannelMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)

	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/new-channel-error/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "CREATE")

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

func TestUpdateRecord(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "test-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "UPDATE")

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

func TestUpdateRecordError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "update-record-error"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "UPDATE")

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
func TestUpdateRecordGetError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "error-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "UPDATE")

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
func TestUpdateRecordMissingError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "random-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "UPDATE")

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
func TestUpdateRecordMissingChannel(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "test-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/random-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "UPDATE")

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
func TestUpdateRecordRevError(t *testing.T) {
	getChannelClient = getChannelClientMock
	getCurrentRevision = getCurrentRevisionMock
	getChannel = GetChannelMock
	getRecord = GetRecordMock
	createRecord = CreateRecordMock
	swaggerSpec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTrillianAgentAPI(swaggerSpec)
	handler := configureAPI(api)
	assert.Equal(t, true, true)
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}
	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel-bad-map-id/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "UPDATE")

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
func TestUpdateRecordGetClientError(t *testing.T) {
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
	payload := map[string]interface{}{
		"test": "test",
	}
	recordID := "new-record"
	record := models.RecordDefinition{RecordID: &recordID, RecordIDPayload: payload}

	reqBody, _ := record.MarshalBinary()
	req, err := http.NewRequest("POST", "/channels/test-channel/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("commit-type", "UPDATE")

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
	if mapID == -2 || mapID == 321 {
		return 0, errors.New("test-error")
	}
	return uint64(1654), nil
}

func getCurrentRevisionErrorMock(c *client.MapClient, ctx context.Context, mapID int64, tracer opentracing.Tracer) (uint64, error) {
	return 0, errors.New("test-error")
}

func GetChannelMock(ctx context.Context, client *client.MapClient, channelID string, tracer opentracing.Tracer) (*models.Channel, error) {
	if channelID == "test-channel" {
		return &models.Channel{ChannelID: "test-channel", MapID: 1536}, nil
	} else if channelID == "test-channel-bad-map-id" {
		return &models.Channel{ChannelID: "test-channel", MapID: 321}, nil
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
	if recordID == "test-record" || recordID == "update-record-error" {
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

func CreateRecordMock(ctx context.Context, client *client.Client, revision int64, prevRevision int64, channelID string, commitType string, recordDef *models.RecordDefinition, tracer opentracing.Tracer) error {
	if *recordDef.RecordID == "new-record-error" || *recordDef.RecordID == "update-record-error" {
		return errors.New("create-channel-error")
	}
	return nil
}

func CreateChannelMock(ctx context.Context, trillAdminClient trillian.TrillianAdminClient, trillMapClient trillian.TrillianMapClient, trillMapWriteClient trillian.TrillianMapWriteClient, revision int64, channelMapID int64, channelID string, tracer opentracing.Tracer) (int64, error) {
	if channelID == "new-channel-error" {
		return -1, errors.New("create-channel-error")
	}
	return 651, nil
}
