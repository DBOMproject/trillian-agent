/*
 * Copyright 2020 Unisys Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package responses

import (
	"errors"
	"trillian-agent/logger"
	"trillian-agent/models"
	"trillian-agent/restapi/operations/record"
)

var log = logger.GetLogger("Responses:Error")

//ChannelNotFound is the message to log if a channel is not found
var ChannelNotFound = "No Such Channel"

//ResourceNotFound is the message to log if a resource is not found
var ResourceNotFound = "No Such Resource"

//ResourceExists is the message to log if a resource already exists
var ResourceExists = "Resource Already Exists"

//InvalidCommitType is the message to log if an invalid commit type is specified
var InvalidCommitType = "Invalid Commit Type"

//InternalError is the messsage to log if an internal erro occurs
var InternalError = "Internal Error"

//ErrAuditInternalServerError returns error when an internal error occurs
func ErrAuditInternalServerError(err error) *record.AuditRecordInternalServerError {
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.AuditRecordInternalServerError{Payload: &errRes}
	return &res
}

//ErrAuditChannelNotFound returns error for when a channel is not found
func ErrAuditChannelNotFound() *record.AuditRecordNotFound {
	err := errors.New(ChannelNotFound)
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.AuditRecordNotFound{Payload: &errRes}
	return &res
}

//ErrAuditResourceNotFound returns error for when a resource is not found
func ErrAuditResourceNotFound() *record.AuditRecordNotFound {
	err := errors.New(ResourceNotFound)
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.AuditRecordNotFound{Payload: &errRes}
	return &res
}

//ErrCommitInternalServerError returns rror when an internal error occurs
func ErrCommitInternalServerError(err error) *record.CommitRecordInternalServerError {
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.CommitRecordInternalServerError{Payload: &errRes}
	return &res
}

//ErrCommitRecordConflict returns error for when there is a resource conflict
func ErrCommitRecordConflict() *record.CommitRecordConflict {
	err := errors.New(ResourceExists)
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.CommitRecordConflict{Payload: &errRes}
	return &res
}

//ErrCommitChannelNotFound returns error for when a channel is not found
func ErrCommitChannelNotFound() *record.CommitRecordNotFound {
	err := errors.New(ChannelNotFound)
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.CommitRecordNotFound{Payload: &errRes}
	return &res
}

//ErrCommitResourceNotFound returns error for when a resource is not found
func ErrCommitResourceNotFound() *record.CommitRecordNotFound {
	err := errors.New(ResourceNotFound)
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.CommitRecordNotFound{Payload: &errRes}
	return &res
}

//ErrCommitInvalidCommitType returns error for when an invalid commit type is provided
func ErrCommitInvalidCommitType() *record.CommitRecordNotFound {
	err := errors.New(InvalidCommitType)
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.CommitRecordNotFound{Payload: &errRes}
	return &res
}

//ErrRetrieveRecordInternalServerError returns error when an internal error occurs
func ErrRetrieveRecordInternalServerError(err error) *record.RetrieveRecordInternalServerError {
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.RetrieveRecordInternalServerError{Payload: &errRes}
	return &res
}

//ErrRetrieveChannelNotFound returns error for when a channel is not found
func ErrRetrieveChannelNotFound() *record.RetrieveRecordNotFound {
	err := errors.New(ChannelNotFound)
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.RetrieveRecordNotFound{Payload: &errRes}
	return &res
}

//ErrRetrieveResourceNotFound returns error for when a resource is not found
func ErrRetrieveResourceNotFound() *record.RetrieveRecordNotFound {
	err := errors.New(ResourceNotFound)
	var status = err.Error()
	log.Err(err).Msg(status)
	var success = false
	var errRes = models.ErrorResponseDefinition{Status: &status, Success: &success}
	var res = record.RetrieveRecordNotFound{Payload: &errRes}
	return &res
}
