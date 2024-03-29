// Code generated by ent, DO NOT EDIT.

package ent

import (
	"assignment3/ent/request"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Request is the model entity for the Request schema.
type Request struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RequestURL holds the value of the "request_url" field.
	RequestURL string `json:"request_url,omitempty"`
	// RequestMethod holds the value of the "request_method" field.
	RequestMethod string `json:"request_method,omitempty"`
	// RequestHeaders holds the value of the "request_headers" field.
	RequestHeaders map[string]interface{} `json:"request_headers,omitempty"`
	// RequestBody holds the value of the "request_body" field.
	RequestBody map[string]interface{} `json:"request_body,omitempty"`
	// ResponseStatusCode holds the value of the "response_status_code" field.
	ResponseStatusCode int32 `json:"response_status_code,omitempty"`
	// ResponseBody holds the value of the "response_body" field.
	ResponseBody map[string]interface{} `json:"response_body,omitempty"`
	// RequestTimestamp holds the value of the "request_timestamp" field.
	RequestTimestamp time.Time `json:"request_timestamp,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Request) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case request.FieldRequestHeaders, request.FieldRequestBody, request.FieldResponseBody:
			values[i] = new([]byte)
		case request.FieldID, request.FieldResponseStatusCode:
			values[i] = new(sql.NullInt64)
		case request.FieldRequestURL, request.FieldRequestMethod:
			values[i] = new(sql.NullString)
		case request.FieldRequestTimestamp, request.FieldCreatedAt, request.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Request fields.
func (r *Request) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case request.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case request.FieldRequestURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field request_url", values[i])
			} else if value.Valid {
				r.RequestURL = value.String
			}
		case request.FieldRequestMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field request_method", values[i])
			} else if value.Valid {
				r.RequestMethod = value.String
			}
		case request.FieldRequestHeaders:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field request_headers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.RequestHeaders); err != nil {
					return fmt.Errorf("unmarshal field request_headers: %w", err)
				}
			}
		case request.FieldRequestBody:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field request_body", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.RequestBody); err != nil {
					return fmt.Errorf("unmarshal field request_body: %w", err)
				}
			}
		case request.FieldResponseStatusCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field response_status_code", values[i])
			} else if value.Valid {
				r.ResponseStatusCode = int32(value.Int64)
			}
		case request.FieldResponseBody:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field response_body", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.ResponseBody); err != nil {
					return fmt.Errorf("unmarshal field response_body: %w", err)
				}
			}
		case request.FieldRequestTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field request_timestamp", values[i])
			} else if value.Valid {
				r.RequestTimestamp = value.Time
			}
		case request.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case request.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Request.
// This includes values selected through modifiers, order, etc.
func (r *Request) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// Update returns a builder for updating this Request.
// Note that you need to call Request.Unwrap() before calling this method if this Request
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Request) Update() *RequestUpdateOne {
	return NewRequestClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Request entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Request) Unwrap() *Request {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Request is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Request) String() string {
	var builder strings.Builder
	builder.WriteString("Request(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("request_url=")
	builder.WriteString(r.RequestURL)
	builder.WriteString(", ")
	builder.WriteString("request_method=")
	builder.WriteString(r.RequestMethod)
	builder.WriteString(", ")
	builder.WriteString("request_headers=")
	builder.WriteString(fmt.Sprintf("%v", r.RequestHeaders))
	builder.WriteString(", ")
	builder.WriteString("request_body=")
	builder.WriteString(fmt.Sprintf("%v", r.RequestBody))
	builder.WriteString(", ")
	builder.WriteString("response_status_code=")
	builder.WriteString(fmt.Sprintf("%v", r.ResponseStatusCode))
	builder.WriteString(", ")
	builder.WriteString("response_body=")
	builder.WriteString(fmt.Sprintf("%v", r.ResponseBody))
	builder.WriteString(", ")
	builder.WriteString("request_timestamp=")
	builder.WriteString(r.RequestTimestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Requests is a parsable slice of Request.
type Requests []*Request
