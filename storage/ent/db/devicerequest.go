// Code generated by entc, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/dexidp/dex/storage/ent/db/devicerequest"
)

// DeviceRequest is the model entity for the DeviceRequest schema.
type DeviceRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserCode holds the value of the "user_code" field.
	UserCode string `json:"user_code,omitempty"`
	// DeviceCode holds the value of the "device_code" field.
	DeviceCode string `json:"device_code,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID string `json:"client_id,omitempty"`
	// ClientSecret holds the value of the "client_secret" field.
	ClientSecret string `json:"client_secret,omitempty"`
	// Scopes holds the value of the "scopes" field.
	Scopes []string `json:"scopes,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeviceRequest) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case devicerequest.FieldScopes:
			values[i] = new([]byte)
		case devicerequest.FieldID:
			values[i] = new(sql.NullInt64)
		case devicerequest.FieldUserCode, devicerequest.FieldDeviceCode, devicerequest.FieldClientID, devicerequest.FieldClientSecret:
			values[i] = new(sql.NullString)
		case devicerequest.FieldExpiry:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DeviceRequest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeviceRequest fields.
func (dr *DeviceRequest) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case devicerequest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dr.ID = int(value.Int64)
		case devicerequest.FieldUserCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_code", values[i])
			} else if value.Valid {
				dr.UserCode = value.String
			}
		case devicerequest.FieldDeviceCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_code", values[i])
			} else if value.Valid {
				dr.DeviceCode = value.String
			}
		case devicerequest.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				dr.ClientID = value.String
			}
		case devicerequest.FieldClientSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_secret", values[i])
			} else if value.Valid {
				dr.ClientSecret = value.String
			}
		case devicerequest.FieldScopes:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field scopes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dr.Scopes); err != nil {
					return fmt.Errorf("unmarshal field scopes: %w", err)
				}
			}
		case devicerequest.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				dr.Expiry = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DeviceRequest.
// Note that you need to call DeviceRequest.Unwrap() before calling this method if this DeviceRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (dr *DeviceRequest) Update() *DeviceRequestUpdateOne {
	return (&DeviceRequestClient{config: dr.config}).UpdateOne(dr)
}

// Unwrap unwraps the DeviceRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dr *DeviceRequest) Unwrap() *DeviceRequest {
	tx, ok := dr.config.driver.(*txDriver)
	if !ok {
		panic("db: DeviceRequest is not a transactional entity")
	}
	dr.config.driver = tx.drv
	return dr
}

// String implements the fmt.Stringer.
func (dr *DeviceRequest) String() string {
	var builder strings.Builder
	builder.WriteString("DeviceRequest(")
	builder.WriteString(fmt.Sprintf("id=%v", dr.ID))
	builder.WriteString(", user_code=")
	builder.WriteString(dr.UserCode)
	builder.WriteString(", device_code=")
	builder.WriteString(dr.DeviceCode)
	builder.WriteString(", client_id=")
	builder.WriteString(dr.ClientID)
	builder.WriteString(", client_secret=")
	builder.WriteString(dr.ClientSecret)
	builder.WriteString(", scopes=")
	builder.WriteString(fmt.Sprintf("%v", dr.Scopes))
	builder.WriteString(", expiry=")
	builder.WriteString(dr.Expiry.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DeviceRequests is a parsable slice of DeviceRequest.
type DeviceRequests []*DeviceRequest

func (dr DeviceRequests) config(cfg config) {
	for _i := range dr {
		dr[_i].config = cfg
	}
}