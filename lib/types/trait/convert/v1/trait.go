/*
Copyright 2023 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package traitv1

import (
	traitv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/trait/v1"
	"github.com/gravitational/teleport/lib/types/trait"
)

// FromProto converts an array of v1 traits into a map of string to string array.
func FromProto(traits []*traitv1.Trait) trait.Traits {
	out := map[string][]string{}
	for _, trait := range traits {
		out[trait.Key] = trait.Values
	}
	return out
}

// ToV1 converts a map of string to string array to an array of v1 traits.
func ToProto(traits trait.Traits) []*traitv1.Trait {
	out := make([]*traitv1.Trait, 0, len(traits))
	for key, values := range traits {
		out = append(out, &traitv1.Trait{
			Key:    key,
			Values: values,
		})
	}
	return out
}
