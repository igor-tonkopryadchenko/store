// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/addons_v1.proto

package service

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ListSectionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListSectionsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSectionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSectionsRequestMultiError, or nil if none found.
func (m *ListSectionsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSectionsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetLocale() == nil {
		err := ListSectionsRequestValidationError{
			field:  "Locale",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetLocale()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListSectionsRequestValidationError{
					field:  "Locale",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListSectionsRequestValidationError{
					field:  "Locale",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocale()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListSectionsRequestValidationError{
				field:  "Locale",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetWeek() == nil {
		err := ListSectionsRequestValidationError{
			field:  "Week",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetWeek()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListSectionsRequestValidationError{
					field:  "Week",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListSectionsRequestValidationError{
					field:  "Week",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWeek()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListSectionsRequestValidationError{
				field:  "Week",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListSectionsRequestMultiError(errors)
	}
	return nil
}

// ListSectionsRequestMultiError is an error wrapping multiple validation
// errors returned by ListSectionsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListSectionsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSectionsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSectionsRequestMultiError) AllErrors() []error { return m }

// ListSectionsRequestValidationError is the validation error returned by
// ListSectionsRequest.Validate if the designated constraints aren't met.
type ListSectionsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSectionsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSectionsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSectionsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSectionsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSectionsRequestValidationError) ErrorName() string {
	return "ListSectionsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListSectionsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSectionsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSectionsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSectionsRequestValidationError{}

// Validate checks the field values on ListSectionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListSectionsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSectionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSectionsResponseMultiError, or nil if none found.
func (m *ListSectionsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSectionsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSections() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListSectionsResponseValidationError{
						field:  fmt.Sprintf("Sections[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListSectionsResponseValidationError{
						field:  fmt.Sprintf("Sections[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSectionsResponseValidationError{
					field:  fmt.Sprintf("Sections[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListSectionsResponseMultiError(errors)
	}
	return nil
}

// ListSectionsResponseMultiError is an error wrapping multiple validation
// errors returned by ListSectionsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListSectionsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSectionsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSectionsResponseMultiError) AllErrors() []error { return m }

// ListSectionsResponseValidationError is the validation error returned by
// ListSectionsResponse.Validate if the designated constraints aren't met.
type ListSectionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSectionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSectionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSectionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSectionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSectionsResponseValidationError) ErrorName() string {
	return "ListSectionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListSectionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSectionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSectionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSectionsResponseValidationError{}

// Validate checks the field values on UpdateSection with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateSection) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSection with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateSectionMultiError, or
// nil if none found.
func (m *UpdateSection) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSection) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetIndex(); val <= 0 || val >= 2030 {
		err := UpdateSectionValidationError{
			field:  "Index",
			reason: "value must be inside range (0, 2030)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	// no validation rules for ImageUri

	if len(errors) > 0 {
		return UpdateSectionMultiError(errors)
	}
	return nil
}

// UpdateSectionMultiError is an error wrapping multiple validation errors
// returned by UpdateSection.ValidateAll() if the designated constraints
// aren't met.
type UpdateSectionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSectionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSectionMultiError) AllErrors() []error { return m }

// UpdateSectionValidationError is the validation error returned by
// UpdateSection.Validate if the designated constraints aren't met.
type UpdateSectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSectionValidationError) ErrorName() string { return "UpdateSectionValidationError" }

// Error satisfies the builtin error interface
func (e UpdateSectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSectionValidationError{}

// Validate checks the field values on UpdateSectionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSectionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSectionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSectionRequestMultiError, or nil if none found.
func (m *UpdateSectionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSectionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Handle

	if m.GetUpdateSection() == nil {
		err := UpdateSectionRequestValidationError{
			field:  "UpdateSection",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUpdateSection()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSectionRequestValidationError{
					field:  "UpdateSection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSectionRequestValidationError{
					field:  "UpdateSection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateSection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSectionRequestValidationError{
				field:  "UpdateSection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateSectionRequestMultiError(errors)
	}
	return nil
}

// UpdateSectionRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateSectionRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateSectionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSectionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSectionRequestMultiError) AllErrors() []error { return m }

// UpdateSectionRequestValidationError is the validation error returned by
// UpdateSectionRequest.Validate if the designated constraints aren't met.
type UpdateSectionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSectionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSectionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSectionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSectionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSectionRequestValidationError) ErrorName() string {
	return "UpdateSectionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSectionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSectionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSectionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSectionRequestValidationError{}

// Validate checks the field values on NoResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NoResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NoResponseMultiError, or
// nil if none found.
func (m *NoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *NoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return NoResponseMultiError(errors)
	}
	return nil
}

// NoResponseMultiError is an error wrapping multiple validation errors
// returned by NoResponse.ValidateAll() if the designated constraints aren't met.
type NoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoResponseMultiError) AllErrors() []error { return m }

// NoResponseValidationError is the validation error returned by
// NoResponse.Validate if the designated constraints aren't met.
type NoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoResponseValidationError) ErrorName() string { return "NoResponseValidationError" }

// Error satisfies the builtin error interface
func (e NoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoResponseValidationError{}
