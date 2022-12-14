// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e


package types

// FunctionScore type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e/specification/_types/query_dsl/compound.ts#L107-L127
type FunctionScore struct {
	Exp              *DecayFunction                 `json:"exp,omitempty"`
	FieldValueFactor *FieldValueFactorScoreFunction `json:"field_value_factor,omitempty"`
	Filter           *Query                         `json:"filter,omitempty"`
	Gauss            *DecayFunction                 `json:"gauss,omitempty"`
	Linear           *DecayFunction                 `json:"linear,omitempty"`
	RandomScore      *RandomScoreFunction           `json:"random_score,omitempty"`
	ScriptScore      *ScriptScoreFunction           `json:"script_score,omitempty"`
	Weight           *float64                       `json:"weight,omitempty"`
}

// NewFunctionScore returns a FunctionScore.
func NewFunctionScore() *FunctionScore {
	r := &FunctionScore{}

	return r
}