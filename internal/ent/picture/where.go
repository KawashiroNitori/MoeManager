// Code generated by ent, DO NOT EDIT.

package picture

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/KawashiroNitori/MoeManager/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldID, id))
}

// Filename applies equality check predicate on the "filename" field. It's identical to FilenameEQ.
func Filename(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldFilename, v))
}

// OriginalFilename applies equality check predicate on the "original_filename" field. It's identical to OriginalFilenameEQ.
func OriginalFilename(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldOriginalFilename, v))
}

// Dir applies equality check predicate on the "dir" field. It's identical to DirEQ.
func Dir(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldDir, v))
}

// Digest applies equality check predicate on the "digest" field. It's identical to DigestEQ.
func Digest(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldDigest, v))
}

// IsUpscaled applies equality check predicate on the "is_upscaled" field. It's identical to IsUpscaledEQ.
func IsUpscaled(v bool) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldIsUpscaled, v))
}

// OriginalWidth applies equality check predicate on the "original_width" field. It's identical to OriginalWidthEQ.
func OriginalWidth(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldOriginalWidth, v))
}

// OriginalHeight applies equality check predicate on the "original_height" field. It's identical to OriginalHeightEQ.
func OriginalHeight(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldOriginalHeight, v))
}

// UpscaledWidth applies equality check predicate on the "upscaled_width" field. It's identical to UpscaledWidthEQ.
func UpscaledWidth(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldUpscaledWidth, v))
}

// UpscaledHeight applies equality check predicate on the "upscaled_height" field. It's identical to UpscaledHeightEQ.
func UpscaledHeight(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldUpscaledHeight, v))
}

// UpscaleRatio applies equality check predicate on the "upscale_ratio" field. It's identical to UpscaleRatioEQ.
func UpscaleRatio(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldUpscaleRatio, v))
}

// ErrorMessage applies equality check predicate on the "error_message" field. It's identical to ErrorMessageEQ.
func ErrorMessage(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldErrorMessage, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldCreatedAt, v))
}

// FilenameEQ applies the EQ predicate on the "filename" field.
func FilenameEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldFilename, v))
}

// FilenameNEQ applies the NEQ predicate on the "filename" field.
func FilenameNEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldFilename, v))
}

// FilenameIn applies the In predicate on the "filename" field.
func FilenameIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldFilename, vs...))
}

// FilenameNotIn applies the NotIn predicate on the "filename" field.
func FilenameNotIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldFilename, vs...))
}

// FilenameGT applies the GT predicate on the "filename" field.
func FilenameGT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldFilename, v))
}

// FilenameGTE applies the GTE predicate on the "filename" field.
func FilenameGTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldFilename, v))
}

// FilenameLT applies the LT predicate on the "filename" field.
func FilenameLT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldFilename, v))
}

// FilenameLTE applies the LTE predicate on the "filename" field.
func FilenameLTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldFilename, v))
}

// FilenameContains applies the Contains predicate on the "filename" field.
func FilenameContains(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContains(FieldFilename, v))
}

// FilenameHasPrefix applies the HasPrefix predicate on the "filename" field.
func FilenameHasPrefix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasPrefix(FieldFilename, v))
}

// FilenameHasSuffix applies the HasSuffix predicate on the "filename" field.
func FilenameHasSuffix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasSuffix(FieldFilename, v))
}

// FilenameEqualFold applies the EqualFold predicate on the "filename" field.
func FilenameEqualFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEqualFold(FieldFilename, v))
}

// FilenameContainsFold applies the ContainsFold predicate on the "filename" field.
func FilenameContainsFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContainsFold(FieldFilename, v))
}

// OriginalFilenameEQ applies the EQ predicate on the "original_filename" field.
func OriginalFilenameEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldOriginalFilename, v))
}

// OriginalFilenameNEQ applies the NEQ predicate on the "original_filename" field.
func OriginalFilenameNEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldOriginalFilename, v))
}

// OriginalFilenameIn applies the In predicate on the "original_filename" field.
func OriginalFilenameIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldOriginalFilename, vs...))
}

// OriginalFilenameNotIn applies the NotIn predicate on the "original_filename" field.
func OriginalFilenameNotIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldOriginalFilename, vs...))
}

// OriginalFilenameGT applies the GT predicate on the "original_filename" field.
func OriginalFilenameGT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldOriginalFilename, v))
}

// OriginalFilenameGTE applies the GTE predicate on the "original_filename" field.
func OriginalFilenameGTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldOriginalFilename, v))
}

// OriginalFilenameLT applies the LT predicate on the "original_filename" field.
func OriginalFilenameLT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldOriginalFilename, v))
}

// OriginalFilenameLTE applies the LTE predicate on the "original_filename" field.
func OriginalFilenameLTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldOriginalFilename, v))
}

// OriginalFilenameContains applies the Contains predicate on the "original_filename" field.
func OriginalFilenameContains(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContains(FieldOriginalFilename, v))
}

// OriginalFilenameHasPrefix applies the HasPrefix predicate on the "original_filename" field.
func OriginalFilenameHasPrefix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasPrefix(FieldOriginalFilename, v))
}

// OriginalFilenameHasSuffix applies the HasSuffix predicate on the "original_filename" field.
func OriginalFilenameHasSuffix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasSuffix(FieldOriginalFilename, v))
}

// OriginalFilenameEqualFold applies the EqualFold predicate on the "original_filename" field.
func OriginalFilenameEqualFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEqualFold(FieldOriginalFilename, v))
}

// OriginalFilenameContainsFold applies the ContainsFold predicate on the "original_filename" field.
func OriginalFilenameContainsFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContainsFold(FieldOriginalFilename, v))
}

// DirEQ applies the EQ predicate on the "dir" field.
func DirEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldDir, v))
}

// DirNEQ applies the NEQ predicate on the "dir" field.
func DirNEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldDir, v))
}

// DirIn applies the In predicate on the "dir" field.
func DirIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldDir, vs...))
}

// DirNotIn applies the NotIn predicate on the "dir" field.
func DirNotIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldDir, vs...))
}

// DirGT applies the GT predicate on the "dir" field.
func DirGT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldDir, v))
}

// DirGTE applies the GTE predicate on the "dir" field.
func DirGTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldDir, v))
}

// DirLT applies the LT predicate on the "dir" field.
func DirLT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldDir, v))
}

// DirLTE applies the LTE predicate on the "dir" field.
func DirLTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldDir, v))
}

// DirContains applies the Contains predicate on the "dir" field.
func DirContains(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContains(FieldDir, v))
}

// DirHasPrefix applies the HasPrefix predicate on the "dir" field.
func DirHasPrefix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasPrefix(FieldDir, v))
}

// DirHasSuffix applies the HasSuffix predicate on the "dir" field.
func DirHasSuffix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasSuffix(FieldDir, v))
}

// DirEqualFold applies the EqualFold predicate on the "dir" field.
func DirEqualFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEqualFold(FieldDir, v))
}

// DirContainsFold applies the ContainsFold predicate on the "dir" field.
func DirContainsFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContainsFold(FieldDir, v))
}

// DigestEQ applies the EQ predicate on the "digest" field.
func DigestEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldDigest, v))
}

// DigestNEQ applies the NEQ predicate on the "digest" field.
func DigestNEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldDigest, v))
}

// DigestIn applies the In predicate on the "digest" field.
func DigestIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldDigest, vs...))
}

// DigestNotIn applies the NotIn predicate on the "digest" field.
func DigestNotIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldDigest, vs...))
}

// DigestGT applies the GT predicate on the "digest" field.
func DigestGT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldDigest, v))
}

// DigestGTE applies the GTE predicate on the "digest" field.
func DigestGTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldDigest, v))
}

// DigestLT applies the LT predicate on the "digest" field.
func DigestLT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldDigest, v))
}

// DigestLTE applies the LTE predicate on the "digest" field.
func DigestLTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldDigest, v))
}

// DigestContains applies the Contains predicate on the "digest" field.
func DigestContains(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContains(FieldDigest, v))
}

// DigestHasPrefix applies the HasPrefix predicate on the "digest" field.
func DigestHasPrefix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasPrefix(FieldDigest, v))
}

// DigestHasSuffix applies the HasSuffix predicate on the "digest" field.
func DigestHasSuffix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasSuffix(FieldDigest, v))
}

// DigestEqualFold applies the EqualFold predicate on the "digest" field.
func DigestEqualFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEqualFold(FieldDigest, v))
}

// DigestContainsFold applies the ContainsFold predicate on the "digest" field.
func DigestContainsFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContainsFold(FieldDigest, v))
}

// IsUpscaledEQ applies the EQ predicate on the "is_upscaled" field.
func IsUpscaledEQ(v bool) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldIsUpscaled, v))
}

// IsUpscaledNEQ applies the NEQ predicate on the "is_upscaled" field.
func IsUpscaledNEQ(v bool) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldIsUpscaled, v))
}

// OriginalWidthEQ applies the EQ predicate on the "original_width" field.
func OriginalWidthEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldOriginalWidth, v))
}

// OriginalWidthNEQ applies the NEQ predicate on the "original_width" field.
func OriginalWidthNEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldOriginalWidth, v))
}

// OriginalWidthIn applies the In predicate on the "original_width" field.
func OriginalWidthIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldOriginalWidth, vs...))
}

// OriginalWidthNotIn applies the NotIn predicate on the "original_width" field.
func OriginalWidthNotIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldOriginalWidth, vs...))
}

// OriginalWidthGT applies the GT predicate on the "original_width" field.
func OriginalWidthGT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldOriginalWidth, v))
}

// OriginalWidthGTE applies the GTE predicate on the "original_width" field.
func OriginalWidthGTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldOriginalWidth, v))
}

// OriginalWidthLT applies the LT predicate on the "original_width" field.
func OriginalWidthLT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldOriginalWidth, v))
}

// OriginalWidthLTE applies the LTE predicate on the "original_width" field.
func OriginalWidthLTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldOriginalWidth, v))
}

// OriginalHeightEQ applies the EQ predicate on the "original_height" field.
func OriginalHeightEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldOriginalHeight, v))
}

// OriginalHeightNEQ applies the NEQ predicate on the "original_height" field.
func OriginalHeightNEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldOriginalHeight, v))
}

// OriginalHeightIn applies the In predicate on the "original_height" field.
func OriginalHeightIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldOriginalHeight, vs...))
}

// OriginalHeightNotIn applies the NotIn predicate on the "original_height" field.
func OriginalHeightNotIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldOriginalHeight, vs...))
}

// OriginalHeightGT applies the GT predicate on the "original_height" field.
func OriginalHeightGT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldOriginalHeight, v))
}

// OriginalHeightGTE applies the GTE predicate on the "original_height" field.
func OriginalHeightGTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldOriginalHeight, v))
}

// OriginalHeightLT applies the LT predicate on the "original_height" field.
func OriginalHeightLT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldOriginalHeight, v))
}

// OriginalHeightLTE applies the LTE predicate on the "original_height" field.
func OriginalHeightLTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldOriginalHeight, v))
}

// UpscaledWidthEQ applies the EQ predicate on the "upscaled_width" field.
func UpscaledWidthEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldUpscaledWidth, v))
}

// UpscaledWidthNEQ applies the NEQ predicate on the "upscaled_width" field.
func UpscaledWidthNEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldUpscaledWidth, v))
}

// UpscaledWidthIn applies the In predicate on the "upscaled_width" field.
func UpscaledWidthIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldUpscaledWidth, vs...))
}

// UpscaledWidthNotIn applies the NotIn predicate on the "upscaled_width" field.
func UpscaledWidthNotIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldUpscaledWidth, vs...))
}

// UpscaledWidthGT applies the GT predicate on the "upscaled_width" field.
func UpscaledWidthGT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldUpscaledWidth, v))
}

// UpscaledWidthGTE applies the GTE predicate on the "upscaled_width" field.
func UpscaledWidthGTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldUpscaledWidth, v))
}

// UpscaledWidthLT applies the LT predicate on the "upscaled_width" field.
func UpscaledWidthLT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldUpscaledWidth, v))
}

// UpscaledWidthLTE applies the LTE predicate on the "upscaled_width" field.
func UpscaledWidthLTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldUpscaledWidth, v))
}

// UpscaledHeightEQ applies the EQ predicate on the "upscaled_height" field.
func UpscaledHeightEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldUpscaledHeight, v))
}

// UpscaledHeightNEQ applies the NEQ predicate on the "upscaled_height" field.
func UpscaledHeightNEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldUpscaledHeight, v))
}

// UpscaledHeightIn applies the In predicate on the "upscaled_height" field.
func UpscaledHeightIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldUpscaledHeight, vs...))
}

// UpscaledHeightNotIn applies the NotIn predicate on the "upscaled_height" field.
func UpscaledHeightNotIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldUpscaledHeight, vs...))
}

// UpscaledHeightGT applies the GT predicate on the "upscaled_height" field.
func UpscaledHeightGT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldUpscaledHeight, v))
}

// UpscaledHeightGTE applies the GTE predicate on the "upscaled_height" field.
func UpscaledHeightGTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldUpscaledHeight, v))
}

// UpscaledHeightLT applies the LT predicate on the "upscaled_height" field.
func UpscaledHeightLT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldUpscaledHeight, v))
}

// UpscaledHeightLTE applies the LTE predicate on the "upscaled_height" field.
func UpscaledHeightLTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldUpscaledHeight, v))
}

// UpscaleRatioEQ applies the EQ predicate on the "upscale_ratio" field.
func UpscaleRatioEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldUpscaleRatio, v))
}

// UpscaleRatioNEQ applies the NEQ predicate on the "upscale_ratio" field.
func UpscaleRatioNEQ(v int) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldUpscaleRatio, v))
}

// UpscaleRatioIn applies the In predicate on the "upscale_ratio" field.
func UpscaleRatioIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldUpscaleRatio, vs...))
}

// UpscaleRatioNotIn applies the NotIn predicate on the "upscale_ratio" field.
func UpscaleRatioNotIn(vs ...int) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldUpscaleRatio, vs...))
}

// UpscaleRatioGT applies the GT predicate on the "upscale_ratio" field.
func UpscaleRatioGT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldUpscaleRatio, v))
}

// UpscaleRatioGTE applies the GTE predicate on the "upscale_ratio" field.
func UpscaleRatioGTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldUpscaleRatio, v))
}

// UpscaleRatioLT applies the LT predicate on the "upscale_ratio" field.
func UpscaleRatioLT(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldUpscaleRatio, v))
}

// UpscaleRatioLTE applies the LTE predicate on the "upscale_ratio" field.
func UpscaleRatioLTE(v int) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldUpscaleRatio, v))
}

// ErrorMessageEQ applies the EQ predicate on the "error_message" field.
func ErrorMessageEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldErrorMessage, v))
}

// ErrorMessageNEQ applies the NEQ predicate on the "error_message" field.
func ErrorMessageNEQ(v string) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldErrorMessage, v))
}

// ErrorMessageIn applies the In predicate on the "error_message" field.
func ErrorMessageIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldErrorMessage, vs...))
}

// ErrorMessageNotIn applies the NotIn predicate on the "error_message" field.
func ErrorMessageNotIn(vs ...string) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldErrorMessage, vs...))
}

// ErrorMessageGT applies the GT predicate on the "error_message" field.
func ErrorMessageGT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldErrorMessage, v))
}

// ErrorMessageGTE applies the GTE predicate on the "error_message" field.
func ErrorMessageGTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldErrorMessage, v))
}

// ErrorMessageLT applies the LT predicate on the "error_message" field.
func ErrorMessageLT(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldErrorMessage, v))
}

// ErrorMessageLTE applies the LTE predicate on the "error_message" field.
func ErrorMessageLTE(v string) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldErrorMessage, v))
}

// ErrorMessageContains applies the Contains predicate on the "error_message" field.
func ErrorMessageContains(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContains(FieldErrorMessage, v))
}

// ErrorMessageHasPrefix applies the HasPrefix predicate on the "error_message" field.
func ErrorMessageHasPrefix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasPrefix(FieldErrorMessage, v))
}

// ErrorMessageHasSuffix applies the HasSuffix predicate on the "error_message" field.
func ErrorMessageHasSuffix(v string) predicate.Picture {
	return predicate.Picture(sql.FieldHasSuffix(FieldErrorMessage, v))
}

// ErrorMessageEqualFold applies the EqualFold predicate on the "error_message" field.
func ErrorMessageEqualFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldEqualFold(FieldErrorMessage, v))
}

// ErrorMessageContainsFold applies the ContainsFold predicate on the "error_message" field.
func ErrorMessageContainsFold(v string) predicate.Picture {
	return predicate.Picture(sql.FieldContainsFold(FieldErrorMessage, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldStatus, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Picture {
	return predicate.Picture(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Picture {
	return predicate.Picture(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Picture {
	return predicate.Picture(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Picture {
	return predicate.Picture(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Picture {
	return predicate.Picture(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Picture {
	return predicate.Picture(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Picture {
	return predicate.Picture(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Picture {
	return predicate.Picture(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Picture) predicate.Picture {
	return predicate.Picture(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Picture) predicate.Picture {
	return predicate.Picture(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Picture) predicate.Picture {
	return predicate.Picture(sql.NotPredicates(p))
}