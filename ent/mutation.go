// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/msal4/entgql/ent/assignment"
	"github.com/msal4/entgql/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAssignment = "Assignment"
)

// AssignmentMutation represents an operation that mutates the Assignment nodes in the graph.
type AssignmentMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	is_exam       *bool
	due_date      *time.Time
	duration      *time.Duration
	addduration   *time.Duration
	deleted_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Assignment, error)
	predicates    []predicate.Assignment
}

var _ ent.Mutation = (*AssignmentMutation)(nil)

// assignmentOption allows management of the mutation configuration using functional options.
type assignmentOption func(*AssignmentMutation)

// newAssignmentMutation creates new mutation for the Assignment entity.
func newAssignmentMutation(c config, op Op, opts ...assignmentOption) *AssignmentMutation {
	m := &AssignmentMutation{
		config:        c,
		op:            op,
		typ:           TypeAssignment,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAssignmentID sets the ID field of the mutation.
func withAssignmentID(id uuid.UUID) assignmentOption {
	return func(m *AssignmentMutation) {
		var (
			err   error
			once  sync.Once
			value *Assignment
		)
		m.oldValue = func(ctx context.Context) (*Assignment, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Assignment.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAssignment sets the old Assignment of the mutation.
func withAssignment(node *Assignment) assignmentOption {
	return func(m *AssignmentMutation) {
		m.oldValue = func(context.Context) (*Assignment, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AssignmentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AssignmentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Assignment entities.
func (m *AssignmentMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AssignmentMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetIsExam sets the "is_exam" field.
func (m *AssignmentMutation) SetIsExam(b bool) {
	m.is_exam = &b
}

// IsExam returns the value of the "is_exam" field in the mutation.
func (m *AssignmentMutation) IsExam() (r bool, exists bool) {
	v := m.is_exam
	if v == nil {
		return
	}
	return *v, true
}

// OldIsExam returns the old "is_exam" field's value of the Assignment entity.
// If the Assignment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AssignmentMutation) OldIsExam(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldIsExam is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldIsExam requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsExam: %w", err)
	}
	return oldValue.IsExam, nil
}

// ResetIsExam resets all changes to the "is_exam" field.
func (m *AssignmentMutation) ResetIsExam() {
	m.is_exam = nil
}

// SetDueDate sets the "due_date" field.
func (m *AssignmentMutation) SetDueDate(t time.Time) {
	m.due_date = &t
}

// DueDate returns the value of the "due_date" field in the mutation.
func (m *AssignmentMutation) DueDate() (r time.Time, exists bool) {
	v := m.due_date
	if v == nil {
		return
	}
	return *v, true
}

// OldDueDate returns the old "due_date" field's value of the Assignment entity.
// If the Assignment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AssignmentMutation) OldDueDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDueDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDueDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDueDate: %w", err)
	}
	return oldValue.DueDate, nil
}

// ResetDueDate resets all changes to the "due_date" field.
func (m *AssignmentMutation) ResetDueDate() {
	m.due_date = nil
}

// SetDuration sets the "duration" field.
func (m *AssignmentMutation) SetDuration(t time.Duration) {
	m.duration = &t
	m.addduration = nil
}

// Duration returns the value of the "duration" field in the mutation.
func (m *AssignmentMutation) Duration() (r time.Duration, exists bool) {
	v := m.duration
	if v == nil {
		return
	}
	return *v, true
}

// OldDuration returns the old "duration" field's value of the Assignment entity.
// If the Assignment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AssignmentMutation) OldDuration(ctx context.Context) (v time.Duration, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDuration is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDuration requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDuration: %w", err)
	}
	return oldValue.Duration, nil
}

// AddDuration adds t to the "duration" field.
func (m *AssignmentMutation) AddDuration(t time.Duration) {
	if m.addduration != nil {
		*m.addduration += t
	} else {
		m.addduration = &t
	}
}

// AddedDuration returns the value that was added to the "duration" field in this mutation.
func (m *AssignmentMutation) AddedDuration() (r time.Duration, exists bool) {
	v := m.addduration
	if v == nil {
		return
	}
	return *v, true
}

// ClearDuration clears the value of the "duration" field.
func (m *AssignmentMutation) ClearDuration() {
	m.duration = nil
	m.addduration = nil
	m.clearedFields[assignment.FieldDuration] = struct{}{}
}

// DurationCleared returns if the "duration" field was cleared in this mutation.
func (m *AssignmentMutation) DurationCleared() bool {
	_, ok := m.clearedFields[assignment.FieldDuration]
	return ok
}

// ResetDuration resets all changes to the "duration" field.
func (m *AssignmentMutation) ResetDuration() {
	m.duration = nil
	m.addduration = nil
	delete(m.clearedFields, assignment.FieldDuration)
}

// SetDeletedAt sets the "deleted_at" field.
func (m *AssignmentMutation) SetDeletedAt(t time.Time) {
	m.deleted_at = &t
}

// DeletedAt returns the value of the "deleted_at" field in the mutation.
func (m *AssignmentMutation) DeletedAt() (r time.Time, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old "deleted_at" field's value of the Assignment entity.
// If the Assignment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AssignmentMutation) OldDeletedAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDeletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (m *AssignmentMutation) ClearDeletedAt() {
	m.deleted_at = nil
	m.clearedFields[assignment.FieldDeletedAt] = struct{}{}
}

// DeletedAtCleared returns if the "deleted_at" field was cleared in this mutation.
func (m *AssignmentMutation) DeletedAtCleared() bool {
	_, ok := m.clearedFields[assignment.FieldDeletedAt]
	return ok
}

// ResetDeletedAt resets all changes to the "deleted_at" field.
func (m *AssignmentMutation) ResetDeletedAt() {
	m.deleted_at = nil
	delete(m.clearedFields, assignment.FieldDeletedAt)
}

// Where appends a list predicates to the AssignmentMutation builder.
func (m *AssignmentMutation) Where(ps ...predicate.Assignment) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *AssignmentMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Assignment).
func (m *AssignmentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AssignmentMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.is_exam != nil {
		fields = append(fields, assignment.FieldIsExam)
	}
	if m.due_date != nil {
		fields = append(fields, assignment.FieldDueDate)
	}
	if m.duration != nil {
		fields = append(fields, assignment.FieldDuration)
	}
	if m.deleted_at != nil {
		fields = append(fields, assignment.FieldDeletedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AssignmentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case assignment.FieldIsExam:
		return m.IsExam()
	case assignment.FieldDueDate:
		return m.DueDate()
	case assignment.FieldDuration:
		return m.Duration()
	case assignment.FieldDeletedAt:
		return m.DeletedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AssignmentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case assignment.FieldIsExam:
		return m.OldIsExam(ctx)
	case assignment.FieldDueDate:
		return m.OldDueDate(ctx)
	case assignment.FieldDuration:
		return m.OldDuration(ctx)
	case assignment.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Assignment field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AssignmentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case assignment.FieldIsExam:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsExam(v)
		return nil
	case assignment.FieldDueDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDueDate(v)
		return nil
	case assignment.FieldDuration:
		v, ok := value.(time.Duration)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDuration(v)
		return nil
	case assignment.FieldDeletedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Assignment field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AssignmentMutation) AddedFields() []string {
	var fields []string
	if m.addduration != nil {
		fields = append(fields, assignment.FieldDuration)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AssignmentMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case assignment.FieldDuration:
		return m.AddedDuration()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AssignmentMutation) AddField(name string, value ent.Value) error {
	switch name {
	case assignment.FieldDuration:
		v, ok := value.(time.Duration)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDuration(v)
		return nil
	}
	return fmt.Errorf("unknown Assignment numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AssignmentMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(assignment.FieldDuration) {
		fields = append(fields, assignment.FieldDuration)
	}
	if m.FieldCleared(assignment.FieldDeletedAt) {
		fields = append(fields, assignment.FieldDeletedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AssignmentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AssignmentMutation) ClearField(name string) error {
	switch name {
	case assignment.FieldDuration:
		m.ClearDuration()
		return nil
	case assignment.FieldDeletedAt:
		m.ClearDeletedAt()
		return nil
	}
	return fmt.Errorf("unknown Assignment nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AssignmentMutation) ResetField(name string) error {
	switch name {
	case assignment.FieldIsExam:
		m.ResetIsExam()
		return nil
	case assignment.FieldDueDate:
		m.ResetDueDate()
		return nil
	case assignment.FieldDuration:
		m.ResetDuration()
		return nil
	case assignment.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	}
	return fmt.Errorf("unknown Assignment field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AssignmentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AssignmentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AssignmentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AssignmentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AssignmentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AssignmentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AssignmentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Assignment unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AssignmentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Assignment edge %s", name)
}
