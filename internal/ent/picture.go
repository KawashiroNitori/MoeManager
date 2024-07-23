// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KawashiroNitori/MoeManager/internal/ent/picture"
)

// Picture is the model entity for the Picture schema.
type Picture struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Filename holds the value of the "filename" field.
	Filename string `json:"filename,omitempty"`
	// OriginalFilename holds the value of the "original_filename" field.
	OriginalFilename string `json:"original_filename,omitempty"`
	// Dir holds the value of the "dir" field.
	Dir string `json:"dir,omitempty"`
	// Digest holds the value of the "digest" field.
	Digest string `json:"digest,omitempty"`
	// IsUpscaled holds the value of the "is_upscaled" field.
	IsUpscaled bool `json:"is_upscaled,omitempty"`
	// OriginalWidth holds the value of the "original_width" field.
	OriginalWidth int `json:"original_width,omitempty"`
	// OriginalHeight holds the value of the "original_height" field.
	OriginalHeight int `json:"original_height,omitempty"`
	// UpscaledWidth holds the value of the "upscaled_width" field.
	UpscaledWidth int `json:"upscaled_width,omitempty"`
	// UpscaledHeight holds the value of the "upscaled_height" field.
	UpscaledHeight int `json:"upscaled_height,omitempty"`
	// UpscaleRatio holds the value of the "upscale_ratio" field.
	UpscaleRatio int `json:"upscale_ratio,omitempty"`
	// ErrorMessage holds the value of the "error_message" field.
	ErrorMessage string `json:"error_message,omitempty"`
	// Status holds the value of the "status" field.
	Status picture.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Picture) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case picture.FieldIsUpscaled:
			values[i] = new(sql.NullBool)
		case picture.FieldID, picture.FieldOriginalWidth, picture.FieldOriginalHeight, picture.FieldUpscaledWidth, picture.FieldUpscaledHeight, picture.FieldUpscaleRatio:
			values[i] = new(sql.NullInt64)
		case picture.FieldFilename, picture.FieldOriginalFilename, picture.FieldDir, picture.FieldDigest, picture.FieldErrorMessage, picture.FieldStatus:
			values[i] = new(sql.NullString)
		case picture.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Picture fields.
func (pi *Picture) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case picture.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pi.ID = int(value.Int64)
		case picture.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				pi.Filename = value.String
			}
		case picture.FieldOriginalFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field original_filename", values[i])
			} else if value.Valid {
				pi.OriginalFilename = value.String
			}
		case picture.FieldDir:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dir", values[i])
			} else if value.Valid {
				pi.Dir = value.String
			}
		case picture.FieldDigest:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field digest", values[i])
			} else if value.Valid {
				pi.Digest = value.String
			}
		case picture.FieldIsUpscaled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_upscaled", values[i])
			} else if value.Valid {
				pi.IsUpscaled = value.Bool
			}
		case picture.FieldOriginalWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field original_width", values[i])
			} else if value.Valid {
				pi.OriginalWidth = int(value.Int64)
			}
		case picture.FieldOriginalHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field original_height", values[i])
			} else if value.Valid {
				pi.OriginalHeight = int(value.Int64)
			}
		case picture.FieldUpscaledWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field upscaled_width", values[i])
			} else if value.Valid {
				pi.UpscaledWidth = int(value.Int64)
			}
		case picture.FieldUpscaledHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field upscaled_height", values[i])
			} else if value.Valid {
				pi.UpscaledHeight = int(value.Int64)
			}
		case picture.FieldUpscaleRatio:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field upscale_ratio", values[i])
			} else if value.Valid {
				pi.UpscaleRatio = int(value.Int64)
			}
		case picture.FieldErrorMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_message", values[i])
			} else if value.Valid {
				pi.ErrorMessage = value.String
			}
		case picture.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pi.Status = picture.Status(value.String)
			}
		case picture.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pi.CreatedAt = value.Time
			}
		default:
			pi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Picture.
// This includes values selected through modifiers, order, etc.
func (pi *Picture) Value(name string) (ent.Value, error) {
	return pi.selectValues.Get(name)
}

// Update returns a builder for updating this Picture.
// Note that you need to call Picture.Unwrap() before calling this method if this Picture
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *Picture) Update() *PictureUpdateOne {
	return NewPictureClient(pi.config).UpdateOne(pi)
}

// Unwrap unwraps the Picture entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *Picture) Unwrap() *Picture {
	_tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: Picture is not a transactional entity")
	}
	pi.config.driver = _tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *Picture) String() string {
	var builder strings.Builder
	builder.WriteString("Picture(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pi.ID))
	builder.WriteString("filename=")
	builder.WriteString(pi.Filename)
	builder.WriteString(", ")
	builder.WriteString("original_filename=")
	builder.WriteString(pi.OriginalFilename)
	builder.WriteString(", ")
	builder.WriteString("dir=")
	builder.WriteString(pi.Dir)
	builder.WriteString(", ")
	builder.WriteString("digest=")
	builder.WriteString(pi.Digest)
	builder.WriteString(", ")
	builder.WriteString("is_upscaled=")
	builder.WriteString(fmt.Sprintf("%v", pi.IsUpscaled))
	builder.WriteString(", ")
	builder.WriteString("original_width=")
	builder.WriteString(fmt.Sprintf("%v", pi.OriginalWidth))
	builder.WriteString(", ")
	builder.WriteString("original_height=")
	builder.WriteString(fmt.Sprintf("%v", pi.OriginalHeight))
	builder.WriteString(", ")
	builder.WriteString("upscaled_width=")
	builder.WriteString(fmt.Sprintf("%v", pi.UpscaledWidth))
	builder.WriteString(", ")
	builder.WriteString("upscaled_height=")
	builder.WriteString(fmt.Sprintf("%v", pi.UpscaledHeight))
	builder.WriteString(", ")
	builder.WriteString("upscale_ratio=")
	builder.WriteString(fmt.Sprintf("%v", pi.UpscaleRatio))
	builder.WriteString(", ")
	builder.WriteString("error_message=")
	builder.WriteString(pi.ErrorMessage)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pi.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Pictures is a parsable slice of Picture.
type Pictures []*Picture
