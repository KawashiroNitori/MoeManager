// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KawashiroNitori/MoeManager/internal/ent/picture"
	"github.com/KawashiroNitori/MoeManager/internal/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePicture = "Picture"
)

// PictureMutation represents an operation that mutates the Picture nodes in the graph.
type PictureMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	filename           *string
	original_filename  *string
	dir                *string
	digest             *string
	is_upscaled        *bool
	original_width     *int
	addoriginal_width  *int
	original_height    *int
	addoriginal_height *int
	upscaled_width     *int
	addupscaled_width  *int
	upscaled_height    *int
	addupscaled_height *int
	upscale_ratio      *int
	addupscale_ratio   *int
	error_message      *string
	status             *picture.Status
	created_at         *time.Time
	clearedFields      map[string]struct{}
	done               bool
	oldValue           func(context.Context) (*Picture, error)
	predicates         []predicate.Picture
}

var _ ent.Mutation = (*PictureMutation)(nil)

// pictureOption allows management of the mutation configuration using functional options.
type pictureOption func(*PictureMutation)

// newPictureMutation creates new mutation for the Picture entity.
func newPictureMutation(c config, op Op, opts ...pictureOption) *PictureMutation {
	m := &PictureMutation{
		config:        c,
		op:            op,
		typ:           TypePicture,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPictureID sets the ID field of the mutation.
func withPictureID(id int) pictureOption {
	return func(m *PictureMutation) {
		var (
			err   error
			once  sync.Once
			value *Picture
		)
		m.oldValue = func(ctx context.Context) (*Picture, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Picture.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPicture sets the old Picture of the mutation.
func withPicture(node *Picture) pictureOption {
	return func(m *PictureMutation) {
		m.oldValue = func(context.Context) (*Picture, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PictureMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PictureMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PictureMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PictureMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Picture.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetFilename sets the "filename" field.
func (m *PictureMutation) SetFilename(s string) {
	m.filename = &s
}

// Filename returns the value of the "filename" field in the mutation.
func (m *PictureMutation) Filename() (r string, exists bool) {
	v := m.filename
	if v == nil {
		return
	}
	return *v, true
}

// OldFilename returns the old "filename" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldFilename(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFilename is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFilename requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFilename: %w", err)
	}
	return oldValue.Filename, nil
}

// ResetFilename resets all changes to the "filename" field.
func (m *PictureMutation) ResetFilename() {
	m.filename = nil
}

// SetOriginalFilename sets the "original_filename" field.
func (m *PictureMutation) SetOriginalFilename(s string) {
	m.original_filename = &s
}

// OriginalFilename returns the value of the "original_filename" field in the mutation.
func (m *PictureMutation) OriginalFilename() (r string, exists bool) {
	v := m.original_filename
	if v == nil {
		return
	}
	return *v, true
}

// OldOriginalFilename returns the old "original_filename" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldOriginalFilename(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOriginalFilename is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOriginalFilename requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOriginalFilename: %w", err)
	}
	return oldValue.OriginalFilename, nil
}

// ResetOriginalFilename resets all changes to the "original_filename" field.
func (m *PictureMutation) ResetOriginalFilename() {
	m.original_filename = nil
}

// SetDir sets the "dir" field.
func (m *PictureMutation) SetDir(s string) {
	m.dir = &s
}

// Dir returns the value of the "dir" field in the mutation.
func (m *PictureMutation) Dir() (r string, exists bool) {
	v := m.dir
	if v == nil {
		return
	}
	return *v, true
}

// OldDir returns the old "dir" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldDir(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDir is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDir requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDir: %w", err)
	}
	return oldValue.Dir, nil
}

// ResetDir resets all changes to the "dir" field.
func (m *PictureMutation) ResetDir() {
	m.dir = nil
}

// SetDigest sets the "digest" field.
func (m *PictureMutation) SetDigest(s string) {
	m.digest = &s
}

// Digest returns the value of the "digest" field in the mutation.
func (m *PictureMutation) Digest() (r string, exists bool) {
	v := m.digest
	if v == nil {
		return
	}
	return *v, true
}

// OldDigest returns the old "digest" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldDigest(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDigest is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDigest requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDigest: %w", err)
	}
	return oldValue.Digest, nil
}

// ResetDigest resets all changes to the "digest" field.
func (m *PictureMutation) ResetDigest() {
	m.digest = nil
}

// SetIsUpscaled sets the "is_upscaled" field.
func (m *PictureMutation) SetIsUpscaled(b bool) {
	m.is_upscaled = &b
}

// IsUpscaled returns the value of the "is_upscaled" field in the mutation.
func (m *PictureMutation) IsUpscaled() (r bool, exists bool) {
	v := m.is_upscaled
	if v == nil {
		return
	}
	return *v, true
}

// OldIsUpscaled returns the old "is_upscaled" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldIsUpscaled(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsUpscaled is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsUpscaled requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsUpscaled: %w", err)
	}
	return oldValue.IsUpscaled, nil
}

// ResetIsUpscaled resets all changes to the "is_upscaled" field.
func (m *PictureMutation) ResetIsUpscaled() {
	m.is_upscaled = nil
}

// SetOriginalWidth sets the "original_width" field.
func (m *PictureMutation) SetOriginalWidth(i int) {
	m.original_width = &i
	m.addoriginal_width = nil
}

// OriginalWidth returns the value of the "original_width" field in the mutation.
func (m *PictureMutation) OriginalWidth() (r int, exists bool) {
	v := m.original_width
	if v == nil {
		return
	}
	return *v, true
}

// OldOriginalWidth returns the old "original_width" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldOriginalWidth(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOriginalWidth is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOriginalWidth requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOriginalWidth: %w", err)
	}
	return oldValue.OriginalWidth, nil
}

// AddOriginalWidth adds i to the "original_width" field.
func (m *PictureMutation) AddOriginalWidth(i int) {
	if m.addoriginal_width != nil {
		*m.addoriginal_width += i
	} else {
		m.addoriginal_width = &i
	}
}

// AddedOriginalWidth returns the value that was added to the "original_width" field in this mutation.
func (m *PictureMutation) AddedOriginalWidth() (r int, exists bool) {
	v := m.addoriginal_width
	if v == nil {
		return
	}
	return *v, true
}

// ResetOriginalWidth resets all changes to the "original_width" field.
func (m *PictureMutation) ResetOriginalWidth() {
	m.original_width = nil
	m.addoriginal_width = nil
}

// SetOriginalHeight sets the "original_height" field.
func (m *PictureMutation) SetOriginalHeight(i int) {
	m.original_height = &i
	m.addoriginal_height = nil
}

// OriginalHeight returns the value of the "original_height" field in the mutation.
func (m *PictureMutation) OriginalHeight() (r int, exists bool) {
	v := m.original_height
	if v == nil {
		return
	}
	return *v, true
}

// OldOriginalHeight returns the old "original_height" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldOriginalHeight(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOriginalHeight is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOriginalHeight requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOriginalHeight: %w", err)
	}
	return oldValue.OriginalHeight, nil
}

// AddOriginalHeight adds i to the "original_height" field.
func (m *PictureMutation) AddOriginalHeight(i int) {
	if m.addoriginal_height != nil {
		*m.addoriginal_height += i
	} else {
		m.addoriginal_height = &i
	}
}

// AddedOriginalHeight returns the value that was added to the "original_height" field in this mutation.
func (m *PictureMutation) AddedOriginalHeight() (r int, exists bool) {
	v := m.addoriginal_height
	if v == nil {
		return
	}
	return *v, true
}

// ResetOriginalHeight resets all changes to the "original_height" field.
func (m *PictureMutation) ResetOriginalHeight() {
	m.original_height = nil
	m.addoriginal_height = nil
}

// SetUpscaledWidth sets the "upscaled_width" field.
func (m *PictureMutation) SetUpscaledWidth(i int) {
	m.upscaled_width = &i
	m.addupscaled_width = nil
}

// UpscaledWidth returns the value of the "upscaled_width" field in the mutation.
func (m *PictureMutation) UpscaledWidth() (r int, exists bool) {
	v := m.upscaled_width
	if v == nil {
		return
	}
	return *v, true
}

// OldUpscaledWidth returns the old "upscaled_width" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldUpscaledWidth(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpscaledWidth is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpscaledWidth requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpscaledWidth: %w", err)
	}
	return oldValue.UpscaledWidth, nil
}

// AddUpscaledWidth adds i to the "upscaled_width" field.
func (m *PictureMutation) AddUpscaledWidth(i int) {
	if m.addupscaled_width != nil {
		*m.addupscaled_width += i
	} else {
		m.addupscaled_width = &i
	}
}

// AddedUpscaledWidth returns the value that was added to the "upscaled_width" field in this mutation.
func (m *PictureMutation) AddedUpscaledWidth() (r int, exists bool) {
	v := m.addupscaled_width
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpscaledWidth resets all changes to the "upscaled_width" field.
func (m *PictureMutation) ResetUpscaledWidth() {
	m.upscaled_width = nil
	m.addupscaled_width = nil
}

// SetUpscaledHeight sets the "upscaled_height" field.
func (m *PictureMutation) SetUpscaledHeight(i int) {
	m.upscaled_height = &i
	m.addupscaled_height = nil
}

// UpscaledHeight returns the value of the "upscaled_height" field in the mutation.
func (m *PictureMutation) UpscaledHeight() (r int, exists bool) {
	v := m.upscaled_height
	if v == nil {
		return
	}
	return *v, true
}

// OldUpscaledHeight returns the old "upscaled_height" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldUpscaledHeight(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpscaledHeight is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpscaledHeight requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpscaledHeight: %w", err)
	}
	return oldValue.UpscaledHeight, nil
}

// AddUpscaledHeight adds i to the "upscaled_height" field.
func (m *PictureMutation) AddUpscaledHeight(i int) {
	if m.addupscaled_height != nil {
		*m.addupscaled_height += i
	} else {
		m.addupscaled_height = &i
	}
}

// AddedUpscaledHeight returns the value that was added to the "upscaled_height" field in this mutation.
func (m *PictureMutation) AddedUpscaledHeight() (r int, exists bool) {
	v := m.addupscaled_height
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpscaledHeight resets all changes to the "upscaled_height" field.
func (m *PictureMutation) ResetUpscaledHeight() {
	m.upscaled_height = nil
	m.addupscaled_height = nil
}

// SetUpscaleRatio sets the "upscale_ratio" field.
func (m *PictureMutation) SetUpscaleRatio(i int) {
	m.upscale_ratio = &i
	m.addupscale_ratio = nil
}

// UpscaleRatio returns the value of the "upscale_ratio" field in the mutation.
func (m *PictureMutation) UpscaleRatio() (r int, exists bool) {
	v := m.upscale_ratio
	if v == nil {
		return
	}
	return *v, true
}

// OldUpscaleRatio returns the old "upscale_ratio" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldUpscaleRatio(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpscaleRatio is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpscaleRatio requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpscaleRatio: %w", err)
	}
	return oldValue.UpscaleRatio, nil
}

// AddUpscaleRatio adds i to the "upscale_ratio" field.
func (m *PictureMutation) AddUpscaleRatio(i int) {
	if m.addupscale_ratio != nil {
		*m.addupscale_ratio += i
	} else {
		m.addupscale_ratio = &i
	}
}

// AddedUpscaleRatio returns the value that was added to the "upscale_ratio" field in this mutation.
func (m *PictureMutation) AddedUpscaleRatio() (r int, exists bool) {
	v := m.addupscale_ratio
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpscaleRatio resets all changes to the "upscale_ratio" field.
func (m *PictureMutation) ResetUpscaleRatio() {
	m.upscale_ratio = nil
	m.addupscale_ratio = nil
}

// SetErrorMessage sets the "error_message" field.
func (m *PictureMutation) SetErrorMessage(s string) {
	m.error_message = &s
}

// ErrorMessage returns the value of the "error_message" field in the mutation.
func (m *PictureMutation) ErrorMessage() (r string, exists bool) {
	v := m.error_message
	if v == nil {
		return
	}
	return *v, true
}

// OldErrorMessage returns the old "error_message" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldErrorMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldErrorMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldErrorMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldErrorMessage: %w", err)
	}
	return oldValue.ErrorMessage, nil
}

// ResetErrorMessage resets all changes to the "error_message" field.
func (m *PictureMutation) ResetErrorMessage() {
	m.error_message = nil
}

// SetStatus sets the "status" field.
func (m *PictureMutation) SetStatus(pi picture.Status) {
	m.status = &pi
}

// Status returns the value of the "status" field in the mutation.
func (m *PictureMutation) Status() (r picture.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldStatus(ctx context.Context) (v picture.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *PictureMutation) ResetStatus() {
	m.status = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *PictureMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *PictureMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *PictureMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the PictureMutation builder.
func (m *PictureMutation) Where(ps ...predicate.Picture) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PictureMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PictureMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Picture, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PictureMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PictureMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Picture).
func (m *PictureMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PictureMutation) Fields() []string {
	fields := make([]string, 0, 13)
	if m.filename != nil {
		fields = append(fields, picture.FieldFilename)
	}
	if m.original_filename != nil {
		fields = append(fields, picture.FieldOriginalFilename)
	}
	if m.dir != nil {
		fields = append(fields, picture.FieldDir)
	}
	if m.digest != nil {
		fields = append(fields, picture.FieldDigest)
	}
	if m.is_upscaled != nil {
		fields = append(fields, picture.FieldIsUpscaled)
	}
	if m.original_width != nil {
		fields = append(fields, picture.FieldOriginalWidth)
	}
	if m.original_height != nil {
		fields = append(fields, picture.FieldOriginalHeight)
	}
	if m.upscaled_width != nil {
		fields = append(fields, picture.FieldUpscaledWidth)
	}
	if m.upscaled_height != nil {
		fields = append(fields, picture.FieldUpscaledHeight)
	}
	if m.upscale_ratio != nil {
		fields = append(fields, picture.FieldUpscaleRatio)
	}
	if m.error_message != nil {
		fields = append(fields, picture.FieldErrorMessage)
	}
	if m.status != nil {
		fields = append(fields, picture.FieldStatus)
	}
	if m.created_at != nil {
		fields = append(fields, picture.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PictureMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case picture.FieldFilename:
		return m.Filename()
	case picture.FieldOriginalFilename:
		return m.OriginalFilename()
	case picture.FieldDir:
		return m.Dir()
	case picture.FieldDigest:
		return m.Digest()
	case picture.FieldIsUpscaled:
		return m.IsUpscaled()
	case picture.FieldOriginalWidth:
		return m.OriginalWidth()
	case picture.FieldOriginalHeight:
		return m.OriginalHeight()
	case picture.FieldUpscaledWidth:
		return m.UpscaledWidth()
	case picture.FieldUpscaledHeight:
		return m.UpscaledHeight()
	case picture.FieldUpscaleRatio:
		return m.UpscaleRatio()
	case picture.FieldErrorMessage:
		return m.ErrorMessage()
	case picture.FieldStatus:
		return m.Status()
	case picture.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PictureMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case picture.FieldFilename:
		return m.OldFilename(ctx)
	case picture.FieldOriginalFilename:
		return m.OldOriginalFilename(ctx)
	case picture.FieldDir:
		return m.OldDir(ctx)
	case picture.FieldDigest:
		return m.OldDigest(ctx)
	case picture.FieldIsUpscaled:
		return m.OldIsUpscaled(ctx)
	case picture.FieldOriginalWidth:
		return m.OldOriginalWidth(ctx)
	case picture.FieldOriginalHeight:
		return m.OldOriginalHeight(ctx)
	case picture.FieldUpscaledWidth:
		return m.OldUpscaledWidth(ctx)
	case picture.FieldUpscaledHeight:
		return m.OldUpscaledHeight(ctx)
	case picture.FieldUpscaleRatio:
		return m.OldUpscaleRatio(ctx)
	case picture.FieldErrorMessage:
		return m.OldErrorMessage(ctx)
	case picture.FieldStatus:
		return m.OldStatus(ctx)
	case picture.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Picture field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PictureMutation) SetField(name string, value ent.Value) error {
	switch name {
	case picture.FieldFilename:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFilename(v)
		return nil
	case picture.FieldOriginalFilename:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOriginalFilename(v)
		return nil
	case picture.FieldDir:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDir(v)
		return nil
	case picture.FieldDigest:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDigest(v)
		return nil
	case picture.FieldIsUpscaled:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsUpscaled(v)
		return nil
	case picture.FieldOriginalWidth:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOriginalWidth(v)
		return nil
	case picture.FieldOriginalHeight:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOriginalHeight(v)
		return nil
	case picture.FieldUpscaledWidth:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpscaledWidth(v)
		return nil
	case picture.FieldUpscaledHeight:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpscaledHeight(v)
		return nil
	case picture.FieldUpscaleRatio:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpscaleRatio(v)
		return nil
	case picture.FieldErrorMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetErrorMessage(v)
		return nil
	case picture.FieldStatus:
		v, ok := value.(picture.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case picture.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Picture field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PictureMutation) AddedFields() []string {
	var fields []string
	if m.addoriginal_width != nil {
		fields = append(fields, picture.FieldOriginalWidth)
	}
	if m.addoriginal_height != nil {
		fields = append(fields, picture.FieldOriginalHeight)
	}
	if m.addupscaled_width != nil {
		fields = append(fields, picture.FieldUpscaledWidth)
	}
	if m.addupscaled_height != nil {
		fields = append(fields, picture.FieldUpscaledHeight)
	}
	if m.addupscale_ratio != nil {
		fields = append(fields, picture.FieldUpscaleRatio)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PictureMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case picture.FieldOriginalWidth:
		return m.AddedOriginalWidth()
	case picture.FieldOriginalHeight:
		return m.AddedOriginalHeight()
	case picture.FieldUpscaledWidth:
		return m.AddedUpscaledWidth()
	case picture.FieldUpscaledHeight:
		return m.AddedUpscaledHeight()
	case picture.FieldUpscaleRatio:
		return m.AddedUpscaleRatio()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PictureMutation) AddField(name string, value ent.Value) error {
	switch name {
	case picture.FieldOriginalWidth:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddOriginalWidth(v)
		return nil
	case picture.FieldOriginalHeight:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddOriginalHeight(v)
		return nil
	case picture.FieldUpscaledWidth:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpscaledWidth(v)
		return nil
	case picture.FieldUpscaledHeight:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpscaledHeight(v)
		return nil
	case picture.FieldUpscaleRatio:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpscaleRatio(v)
		return nil
	}
	return fmt.Errorf("unknown Picture numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PictureMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PictureMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PictureMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Picture nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PictureMutation) ResetField(name string) error {
	switch name {
	case picture.FieldFilename:
		m.ResetFilename()
		return nil
	case picture.FieldOriginalFilename:
		m.ResetOriginalFilename()
		return nil
	case picture.FieldDir:
		m.ResetDir()
		return nil
	case picture.FieldDigest:
		m.ResetDigest()
		return nil
	case picture.FieldIsUpscaled:
		m.ResetIsUpscaled()
		return nil
	case picture.FieldOriginalWidth:
		m.ResetOriginalWidth()
		return nil
	case picture.FieldOriginalHeight:
		m.ResetOriginalHeight()
		return nil
	case picture.FieldUpscaledWidth:
		m.ResetUpscaledWidth()
		return nil
	case picture.FieldUpscaledHeight:
		m.ResetUpscaledHeight()
		return nil
	case picture.FieldUpscaleRatio:
		m.ResetUpscaleRatio()
		return nil
	case picture.FieldErrorMessage:
		m.ResetErrorMessage()
		return nil
	case picture.FieldStatus:
		m.ResetStatus()
		return nil
	case picture.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Picture field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PictureMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PictureMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PictureMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PictureMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PictureMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PictureMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PictureMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Picture unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PictureMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Picture edge %s", name)
}
