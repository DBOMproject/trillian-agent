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

package models

import "github.com/go-openapi/swag"

//Record defines the structure for storing a record in trillian
type Record struct {
	AuditDefinition

	// Revision
	// Required: true
	Revision int64 `json:"revision"`

	// Previous Revision
	// Required: true
	PreviousRevision int64 `json:"previousRevision"`
}

// MarshalBinary interface implementation
func (m *Record) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Record) UnmarshalBinary(b []byte) error {
	var res Record
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
