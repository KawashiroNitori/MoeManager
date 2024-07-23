package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Picture holds the schema definition for the Picture entity.
type Picture struct {
	ent.Schema
}

// Fields of the Picture.
func (Picture) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").Unique(),
		field.String("original_filename"),
		field.String("dir"),
		field.String("digest"),
		field.Bool("is_upscaled").Default(false),
		field.Int("original_width"),
		field.Int("original_height"),
		field.Int("upscaled_width").Default(0),
		field.Int("upscaled_height").Default(0),
		field.Int("upscale_ratio").Default(1),
		field.String("error_message").Default(""),
		field.Enum("status").Values("unknown", "processing", "done", "error").Default("unknown"),
		field.Time("created_at").Default(time.Now()),
	}
}

func (Picture) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("original_filename"),
		index.Fields("digest"),
	}
}

// Edges of the Picture.
func (Picture) Edges() []ent.Edge {
	return nil
}
