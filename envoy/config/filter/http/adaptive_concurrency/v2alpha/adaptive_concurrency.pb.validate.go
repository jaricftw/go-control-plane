// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/adaptive_concurrency/v2alpha/adaptive_concurrency.proto

package v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on AdaptiveConcurrency with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AdaptiveConcurrency) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// AdaptiveConcurrencyValidationError is the validation error returned by
// AdaptiveConcurrency.Validate if the designated constraints aren't met.
type AdaptiveConcurrencyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdaptiveConcurrencyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdaptiveConcurrencyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdaptiveConcurrencyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdaptiveConcurrencyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdaptiveConcurrencyValidationError) ErrorName() string {
	return "AdaptiveConcurrencyValidationError"
}

// Error satisfies the builtin error interface
func (e AdaptiveConcurrencyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdaptiveConcurrency.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdaptiveConcurrencyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdaptiveConcurrencyValidationError{}
