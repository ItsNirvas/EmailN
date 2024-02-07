package endpoints

import (
	internalerrors "emailn/internal/internalErrors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BodyForTest struct {
	ID int
}

// SEPARATE FUNCTIONS TO TEST ERRORS (MIGHT NEED LATER BECAUSE IT MIGHT CAUSE SOME
// ERROR ON TESTING DUE TO BE MINIMIZED TO TESTING IT ALL IN ONLY ONE FUNCTION)

/*func Test_HandlerError_EdpointsReturnInternalError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}

func Test_HandlerError_EdpointsReturnDomainError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("Domain Error")
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "Domain Error")
}

func Test_HandlerError_EdpointReturnObjAndStatus(t *testing.T) {
	assert := assert.New(t)
	type BodyForTest struct {
		ID int
	}
	objExpected := BodyForTest{ID: 2}
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, 201, nil
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	objReturned := BodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &objReturned)
	assert.Equal(objExpected, objReturned)
}*/

// ALL TEST FUNCTIONS WERE MINIMIZED TO BE TESTED ALL IN ONE (THIS ONE BELOW)

func Test_HandlerError_EdpointsReturns(t *testing.T) {
	assert := assert.New(t)
	objExpected := BodyForTest{ID: 2}

	endpointInternalErr := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}
	endpointDomainErr := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("Domain Error")
	}
	endpointObjAndStatus := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, 201, nil
	}

	handlerFuncToInternalErr := HandlerError(endpointInternalErr)
	handlerFuncToDomainErr := HandlerError(endpointDomainErr)
	handlerFuncToObjAndStatus := HandlerError(endpointObjAndStatus)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFuncToInternalErr.ServeHTTP(res, req)
	handlerFuncToDomainErr.ServeHTTP(res, req)
	handlerFuncToObjAndStatus.ServeHTTP(res, req)

	if http.StatusInternalServerError == res.Code {
		// INTERNAL ERROR TEST
		assert.Equal(http.StatusInternalServerError, res.Code)
		assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
	}
	if http.StatusBadRequest == res.Code {
		// DOMAIN ERROR TEST
		assert.Equal(http.StatusBadRequest, res.Code)
		assert.Contains(res.Body.String(), "Domain Error")
	}
	if http.StatusCreated == res.Code {
		// RETURNED WHEN IT HAS NO ERRORS TEST
		assert.Equal(http.StatusCreated, res.Code)
		objReturned := BodyForTest{}
		json.Unmarshal(res.Body.Bytes(), &objReturned)
		assert.Equal(objExpected, objReturned)
	}
}
