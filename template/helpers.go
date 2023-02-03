// Copyright 2018 Drone.IO Inc
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

package template

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/aymerick/raymond"
)

// RegisterSettings
// most of this can use template.RegisterSettings(DefaultFunctions)
func RegisterSettings(funcSettings map[string]interface{}) {
	for name, function := range sprig.GenericFuncMap() {
		if invalidHelper(name) {
			continue
		}

		funcSettings[name] = function
	}

	raymond.RegisterHelpers(funcSettings)
}

func invalidHelper(name string) bool {
	invalids := []string{
		"buildCustomCert",
		"decryptAES",
		"derivePassword",
		"encryptAES",
		"fail",
		"genCA",
		"genPrivateKey",
		"genSelfSignedCert",
		"genSignedCert",
		"hello",
		"mustAppend",
		"mustCompact",
		"mustDateModify",
		"mustDeepCopy",
		"mustFirst",
		"mustHas",
		"mustInitial",
		"mustLast",
		"mustMerge",
		"mustMergeOverwrite",
		"mustPrepend",
		"mustPush",
		"mustRegexFind",
		"mustRegexFindAll",
		"mustRegexMatch",
		"mustRegexReplaceAll",
		"mustRegexReplaceAllLiteral",
		"mustRegexSplit",
		"mustRest",
		"mustReverse",
		"mustSlice",
		"mustToDate",
		"mustToJson",
		"mustToPrettyJson",
		"mustToRawJson",
		"mustUniq",
		"mustWithout",
		"must_date_modify",
		"semver",
		"semverCompare",
		"trimall",
	}

	for _, invalid := range invalids {
		if name == invalid {
			return true
		}
	}

	return false
}
