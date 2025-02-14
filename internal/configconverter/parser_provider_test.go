// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configconverter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configmapprovider"
)

func TestMemLimitBallastRemoverPP(t *testing.T) {
	tests := []struct {
		name  string
		fname string
		key   string
	}{
		{
			name:  "default",
			fname: "testdata/ballast_mem_limiter.yaml",
			key:   "processors::memory_limiter::ballast_size_mib",
		},
		{
			name:  "custom",
			fname: "testdata/ballast_mem_limiter_custom.yaml",
			key:   "processors::memory_limiter/foo::ballast_size_mib",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pp := &converterProvider{
				wrapped:     configmapprovider.NewFile(test.fname),
				cfgMapFuncs: []CfgMapFunc{RemoveBallastKey},
			}
			r, err := pp.Retrieve(context.Background(), nil)
			require.NoError(t, err)
			cfgMap, err := r.Get(context.Background())
			require.NoError(t, err)
			b := cfgMap.IsSet(test.key)
			assert.False(t, b)
		})
	}
}
