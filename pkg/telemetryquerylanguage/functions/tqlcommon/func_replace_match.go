// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tqlcommon // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/telemetryquerylanguage/functions/tqlcommon"

import (
	"fmt"

	"github.com/gobwas/glob"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/telemetryquerylanguage/tql"
)

func ReplaceMatch(target tql.GetSetter, pattern string, replacement string) (tql.ExprFunc, error) {
	glob, err := glob.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("the pattern supplied to replace_match is not a valid pattern: %w", err)
	}
	return func(ctx tql.TransformContext) interface{} {
		val := target.Get(ctx)
		if val == nil {
			return nil
		}
		if valStr, ok := val.(string); ok {
			if glob.Match(valStr) {
				target.Set(ctx, replacement)
			}
		}
		return nil
	}, nil
}
