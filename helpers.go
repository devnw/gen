// Copyright 2022 Developer Network
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

package gen

import "context"

// defaultCtx is the default context used by the stream package. This is
// hardcoded to context.Background() but can be overridden by the unit tests.
//
//nolint:gochecknoglobals // this is a valid global
var defaultCtx = context.Background()

// _ctx returns a valid Context with CancelFunc even if it the
// supplied context is initially nil. If the supplied context
// is nil it uses the default context.
func _ctx(c context.Context) (context.Context, context.CancelFunc) {
	if c == nil {
		c = defaultCtx
	}

	return context.WithCancel(c)
}
