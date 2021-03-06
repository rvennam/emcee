// Licensed Materials - Property of IBM
// (C) Copyright IBM Corp. 2019. All Rights Reserved.

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

package pkg

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v2"
)

// ToYAML prints YAML to stdout
func ToYAML(data *OpenAPI, w io.Writer) error {
	d, err := yaml.Marshal(&data)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "%s", d)
	return nil
}
