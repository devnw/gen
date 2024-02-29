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

package cli

import (
	"fmt"
	"log/slog"
	"reflect"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Args map[string]*Arg

type Arg struct {
	Value       any    `json:"value" yaml:"value"`
	Description string `json:"description" yaml:"description"`
	Required    bool   `json:"required" yaml:"required"`
	GetFN       func() any
}

func (a *Arg) Get() any {
	if a.GetFN != nil {
		out := a.GetFN()
		if out == nil {
			return a.Value
		}

		return out
	}

	return a.Value
}

func (a *Arg) String() string {
	additional := ""

	if a.Required {
		additional = " (required)"
		if a.Value != reflect.Zero(reflect.TypeOf(a.Value)).Interface() {
			additional = fmt.Sprintf(
				" (default: %v, required)",
				a.Value,
			)
		}
	}

	return fmt.Sprintf(
		"%s%s",
		a.Description,
		additional,
	)
}

var ErrUnknownType = fmt.Errorf("unknown type")

//nolint:funlen // this is a necessary evil
func SetArgs(flags *pflag.FlagSet, prefix string, args Args) error {
	for key, val := range args {
		arg := fmt.Sprintf("%s.%s", prefix, key)
		switch v := val.Value.(type) {
		case string:
			flags.String(arg, v, val.String())
			val.GetFN = func() any { return viper.GetString(arg) }
		case []string:
			flags.StringSlice(arg, v, val.String())
			val.GetFN = func() any { return viper.GetStringSlice(arg) }
		case bool:
			flags.Bool(arg, v, val.String())
			val.GetFN = func() any { return viper.GetBool(arg) }
		case int:
			flags.Int(arg, v, val.String())
			val.GetFN = func() any { return viper.GetInt(arg) }
		case uint:
			flags.Uint(arg, v, val.String())
			val.GetFN = func() any { return viper.GetUint(arg) }
		case int8:
			flags.Int8(arg, v, val.String())
			val.GetFN = func() any { return viper.GetInt32(arg) }
		case uint8:
			flags.Uint8(arg, v, val.String())
			val.GetFN = func() any { return viper.GetUint32(arg) }
		case int16:
			flags.Int16(arg, v, val.String())
			val.GetFN = func() any { return viper.GetInt32(arg) }
		case uint16:
			flags.Uint16(arg, v, val.String())
			val.GetFN = func() any { return viper.GetUint32(arg) }
		case int32:
			flags.Int32(arg, v, val.String())
			val.GetFN = func() any { return viper.GetInt32(arg) }
		case uint32:
			flags.Uint32(arg, v, val.String())
			val.GetFN = func() any { return viper.GetUint32(arg) }
		case int64:
			flags.Int64(arg, v, val.String())
			val.GetFN = func() any { return viper.GetInt64(arg) }
		case uint64:
			flags.Uint64(arg, v, val.String())
			val.GetFN = func() any { return viper.GetUint64(arg) }
		case float32:
			flags.Float32(arg, v, val.String())
			val.GetFN = func() any { return viper.GetFloat64(arg) }
		case float64:
			flags.Float64(arg, v, val.String())
			val.GetFN = func() any { return viper.GetFloat64(arg) }
		case time.Duration:
			flags.Duration(arg, v, val.String())
			val.GetFN = func() any { return viper.GetDuration(arg) }
		default:
			slog.Error(
				"unknown type",
				"type", reflect.TypeOf(v),
			)
			return ErrUnknownType
		}

		err := viper.BindPFlag(arg, flags.Lookup(arg))
		if err != nil {
			return err
		}
	}

	return nil
}

func As[T any](in any) T {
	return in.(T)
}
