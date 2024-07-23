// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/KawashiroNitori/MoeManager/internal/ent/picture"
	"github.com/KawashiroNitori/MoeManager/internal/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	pictureFields := schema.Picture{}.Fields()
	_ = pictureFields
	// pictureDescIsUpscaled is the schema descriptor for is_upscaled field.
	pictureDescIsUpscaled := pictureFields[4].Descriptor()
	// picture.DefaultIsUpscaled holds the default value on creation for the is_upscaled field.
	picture.DefaultIsUpscaled = pictureDescIsUpscaled.Default.(bool)
	// pictureDescUpscaledWidth is the schema descriptor for upscaled_width field.
	pictureDescUpscaledWidth := pictureFields[7].Descriptor()
	// picture.DefaultUpscaledWidth holds the default value on creation for the upscaled_width field.
	picture.DefaultUpscaledWidth = pictureDescUpscaledWidth.Default.(int)
	// pictureDescUpscaledHeight is the schema descriptor for upscaled_height field.
	pictureDescUpscaledHeight := pictureFields[8].Descriptor()
	// picture.DefaultUpscaledHeight holds the default value on creation for the upscaled_height field.
	picture.DefaultUpscaledHeight = pictureDescUpscaledHeight.Default.(int)
	// pictureDescUpscaleRatio is the schema descriptor for upscale_ratio field.
	pictureDescUpscaleRatio := pictureFields[9].Descriptor()
	// picture.DefaultUpscaleRatio holds the default value on creation for the upscale_ratio field.
	picture.DefaultUpscaleRatio = pictureDescUpscaleRatio.Default.(int)
	// pictureDescErrorMessage is the schema descriptor for error_message field.
	pictureDescErrorMessage := pictureFields[10].Descriptor()
	// picture.DefaultErrorMessage holds the default value on creation for the error_message field.
	picture.DefaultErrorMessage = pictureDescErrorMessage.Default.(string)
	// pictureDescCreatedAt is the schema descriptor for created_at field.
	pictureDescCreatedAt := pictureFields[12].Descriptor()
	// picture.DefaultCreatedAt holds the default value on creation for the created_at field.
	picture.DefaultCreatedAt = pictureDescCreatedAt.Default.(time.Time)
}
