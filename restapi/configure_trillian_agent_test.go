package restapi

import (
	"bytes"
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"trillian-agent/mock"
	"trillian-agent/models"
	"trillian-agent/restapi/operations"
	client "trillian-agent/trillian"

	tclient "github.com/google/trillian/client"

	"github.com/go-openapi/loads"
	"github.com/google/trillian"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
)

//TestAddRecord tests successfully creating a record
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

//TestAddRecordInvalidType tests invalid commit type
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

//TestAddRecordError tests an error when creating a record
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

//TestAddRecordGetError tests a get record error when creating a record
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

//TestAddRecordConflictError tests a record conflict error when creating a record
func TestAddRecordConflictError(t *testing.T) {
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

//TestAddRecordRevError tests a get revision error when creating a record
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

//TestAddRecordRevError2 tests a get revision error when creating a record
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

//TestAddRecordGetClientError tests a get client error when creating a record
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

//TestAddRecordGetClientError2 tests a get client error when creating a record
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

//TestAddRecordChannelError tests a get channel error when creating a record
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

//TestAddRecordCreateChannel tests a create channel success when creating a record
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

//TestAddRecordCreateChannelError tests a create channel error when creating a record
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

//TestUpdateRecord tests successfully updating a record
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

//TestUpdateRecordError tests an error while updating a record
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

//TestUpdateRecordGetError tests a get record error while updating a record
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

//TestUpdateRecordMissingError tests a record not found error while updating a record
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

//TestUpdateRecordMissingChannel tests a channel not found error while updating a record
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

//TestUpdateRecordRevError tests a get revision error while updating a record
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

//TestUpdateRecordRevError tests a get client error while updating a record
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

//TestGetRecord tests successfully getting a record
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

//TestGetRecordChannelError tests a get channel error while getting a record
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

//TestGetRecordChannelNotFound tests a get channel not found error while getting a record
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

//TestGetRecordError tests a get record error while getting a record
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

//TestGetRecordNotFound tests a get record not found error while getting a record
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

//TestGetRecordGetClientError tests a get client error while getting a record
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

//TestGetRecordGetClientError2 tests a get client error while getting a record
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

//TestAuditRecord tests successfully auditing a record
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

//TestAuditRecordChannelError tests a get channel error auditing a record
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

//TestAuditRecordChannelNotFound tests a get channel  not found error auditing a record
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

//TestAuditRecordError tests a get record error auditing a record
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

//TestAuditRecordNotFound tests a get record not found error auditing a record
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

//TestAuditRecord2Error tests a get record error auditing a record
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

//TestAuditRecordNotFound2 tests a get record not found error auditing a record
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

//TestAuditRecordGetClientError tests a get client error auditing a record
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

//TestAuditRecordGetClientError2 tests a get client error auditing a record
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

	return &tclient.MapClient{Conn: mock.NewTrillianMapMockClient(nil, false, false, false)}, nil
}

func getChannelClientErrorMock(ctx context.Context, trillAdminClient trillian.TrillianAdminClient, trillMapClient trillian.TrillianMapClient, channelMapID int64, tracer opentracing.Tracer) (*tclient.MapClient, error) {
	if channelMapID == 1536 {
		return nil, errors.New("test-error")
	}
	return &tclient.MapClient{Conn: mock.NewTrillianMapMockClient(nil, false, false, false)}, nil
}

func getChannelClientError2Mock(ctx context.Context, trillAdminClient trillian.TrillianAdminClient, trillMapClient trillian.TrillianMapClient, channelMapID int64, tracer opentracing.Tracer) (*tclient.MapClient, error) {
	if channelMapID != 1536 {
		return nil, errors.New("test-error")
	}
	return &tclient.MapClient{Conn: mock.NewTrillianMapMockClient(nil, false, false, false)}, nil
}

func getCurrentRevisionMock(c *client.MapClient, ctx context.Context, mapID int64, tracer opentracing.Tracer) (uint64, error) {
	if mapID == -2 || mapID == 321 {
		return 0, errors.New("test-error")
	}
	return uint64(1654), nil
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
