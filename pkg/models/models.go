// Copyright 2024 SAP SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

type TimestampVersion struct {
	Version string `json:"version"`
}

type MetadataField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Metadata []MetadataField

type InRequest struct {
	Source  Source           `json:"source"`
	Version TimestampVersion `json:"version"`
}

type InResponse struct {
	Version  TimestampVersion `json:"version"`
	Metadata Metadata         `json:"metadata"`
}

type CheckRequest struct {
	Source  Source           `json:"source"`
	Version TimestampVersion `json:"version"`
}

type CheckResponse []TimestampVersion

type Source struct{}
