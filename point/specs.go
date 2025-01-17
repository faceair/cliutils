// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package point

const (
	defaultObjectName    = "default"
	defaultLoggingStatus = "unknown"
)

var (
	DefaultMeasurementName = "__default"

	KeyTime        = NewKey("time", KeyType_I)
	KeyMeasurement = NewKey("measurement", KeyType_S)
	KeySource      = NewKey("source", KeyType_S)
	KeyClass       = NewKey("class", KeyType_S)
	KeyDate        = NewKey("date", KeyType_I)

	KeyName   = NewKey("name", KeyType_D, []byte(defaultObjectName))
	KeyStatus = NewKey("status", KeyType_D, []byte(defaultLoggingStatus))
)

var (
	// For logging, we use measurement-name as source value
	// in kodo, so there should not be any tag/field named
	// with `source`.
	//
	// For object, we use measurement-name as class value
	// in kodo, so there should not be any tag/field named
	// with `class`.
	requiredKeys = map[Category][]*Key{
		Logging: {KeyStatus},
		Object:  {KeyName},
		// TODO: others data type not set...
	}

	disabledKeys = map[Category][]*Key{
		Logging: {KeySource, KeyDate},
		Object:  {KeyClass, KeyDate},
		// TODO: others data type not set...
	}

	DefaultEncoding = LineProtocol
)
