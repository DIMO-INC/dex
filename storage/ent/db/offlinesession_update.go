// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dexidp/dex/storage/ent/db/offlinesession"
	"github.com/dexidp/dex/storage/ent/db/predicate"
)

// OfflineSessionUpdate is the builder for updating OfflineSession entities.
type OfflineSessionUpdate struct {
	config
	hooks    []Hook
	mutation *OfflineSessionMutation
}

// Where appends a list predicates to the OfflineSessionUpdate builder.
func (osu *OfflineSessionUpdate) Where(ps ...predicate.OfflineSession) *OfflineSessionUpdate {
	osu.mutation.Where(ps...)
	return osu
}

// SetUserID sets the "user_id" field.
func (osu *OfflineSessionUpdate) SetUserID(s string) *OfflineSessionUpdate {
	osu.mutation.SetUserID(s)
	return osu
}

// SetConnID sets the "conn_id" field.
func (osu *OfflineSessionUpdate) SetConnID(s string) *OfflineSessionUpdate {
	osu.mutation.SetConnID(s)
	return osu
}

// SetRefresh sets the "refresh" field.
func (osu *OfflineSessionUpdate) SetRefresh(b []byte) *OfflineSessionUpdate {
	osu.mutation.SetRefresh(b)
	return osu
}

// SetConnectorData sets the "connector_data" field.
func (osu *OfflineSessionUpdate) SetConnectorData(b []byte) *OfflineSessionUpdate {
	osu.mutation.SetConnectorData(b)
	return osu
}

// ClearConnectorData clears the value of the "connector_data" field.
func (osu *OfflineSessionUpdate) ClearConnectorData() *OfflineSessionUpdate {
	osu.mutation.ClearConnectorData()
	return osu
}

// Mutation returns the OfflineSessionMutation object of the builder.
func (osu *OfflineSessionUpdate) Mutation() *OfflineSessionMutation {
	return osu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osu *OfflineSessionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(osu.hooks) == 0 {
		if err = osu.check(); err != nil {
			return 0, err
		}
		affected, err = osu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OfflineSessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = osu.check(); err != nil {
				return 0, err
			}
			osu.mutation = mutation
			affected, err = osu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(osu.hooks) - 1; i >= 0; i-- {
			if osu.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = osu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (osu *OfflineSessionUpdate) SaveX(ctx context.Context) int {
	affected, err := osu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osu *OfflineSessionUpdate) Exec(ctx context.Context) error {
	_, err := osu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osu *OfflineSessionUpdate) ExecX(ctx context.Context) {
	if err := osu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osu *OfflineSessionUpdate) check() error {
	if v, ok := osu.mutation.UserID(); ok {
		if err := offlinesession.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf("db: validator failed for field \"user_id\": %w", err)}
		}
	}
	if v, ok := osu.mutation.ConnID(); ok {
		if err := offlinesession.ConnIDValidator(v); err != nil {
			return &ValidationError{Name: "conn_id", err: fmt.Errorf("db: validator failed for field \"conn_id\": %w", err)}
		}
	}
	return nil
}

func (osu *OfflineSessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   offlinesession.Table,
			Columns: offlinesession.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: offlinesession.FieldID,
			},
		},
	}
	if ps := osu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: offlinesession.FieldUserID,
		})
	}
	if value, ok := osu.mutation.ConnID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: offlinesession.FieldConnID,
		})
	}
	if value, ok := osu.mutation.Refresh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: offlinesession.FieldRefresh,
		})
	}
	if value, ok := osu.mutation.ConnectorData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: offlinesession.FieldConnectorData,
		})
	}
	if osu.mutation.ConnectorDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: offlinesession.FieldConnectorData,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, osu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{offlinesession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OfflineSessionUpdateOne is the builder for updating a single OfflineSession entity.
type OfflineSessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OfflineSessionMutation
}

// SetUserID sets the "user_id" field.
func (osuo *OfflineSessionUpdateOne) SetUserID(s string) *OfflineSessionUpdateOne {
	osuo.mutation.SetUserID(s)
	return osuo
}

// SetConnID sets the "conn_id" field.
func (osuo *OfflineSessionUpdateOne) SetConnID(s string) *OfflineSessionUpdateOne {
	osuo.mutation.SetConnID(s)
	return osuo
}

// SetRefresh sets the "refresh" field.
func (osuo *OfflineSessionUpdateOne) SetRefresh(b []byte) *OfflineSessionUpdateOne {
	osuo.mutation.SetRefresh(b)
	return osuo
}

// SetConnectorData sets the "connector_data" field.
func (osuo *OfflineSessionUpdateOne) SetConnectorData(b []byte) *OfflineSessionUpdateOne {
	osuo.mutation.SetConnectorData(b)
	return osuo
}

// ClearConnectorData clears the value of the "connector_data" field.
func (osuo *OfflineSessionUpdateOne) ClearConnectorData() *OfflineSessionUpdateOne {
	osuo.mutation.ClearConnectorData()
	return osuo
}

// Mutation returns the OfflineSessionMutation object of the builder.
func (osuo *OfflineSessionUpdateOne) Mutation() *OfflineSessionMutation {
	return osuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osuo *OfflineSessionUpdateOne) Select(field string, fields ...string) *OfflineSessionUpdateOne {
	osuo.fields = append([]string{field}, fields...)
	return osuo
}

// Save executes the query and returns the updated OfflineSession entity.
func (osuo *OfflineSessionUpdateOne) Save(ctx context.Context) (*OfflineSession, error) {
	var (
		err  error
		node *OfflineSession
	)
	if len(osuo.hooks) == 0 {
		if err = osuo.check(); err != nil {
			return nil, err
		}
		node, err = osuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OfflineSessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = osuo.check(); err != nil {
				return nil, err
			}
			osuo.mutation = mutation
			node, err = osuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(osuo.hooks) - 1; i >= 0; i-- {
			if osuo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = osuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (osuo *OfflineSessionUpdateOne) SaveX(ctx context.Context) *OfflineSession {
	node, err := osuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osuo *OfflineSessionUpdateOne) Exec(ctx context.Context) error {
	_, err := osuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osuo *OfflineSessionUpdateOne) ExecX(ctx context.Context) {
	if err := osuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osuo *OfflineSessionUpdateOne) check() error {
	if v, ok := osuo.mutation.UserID(); ok {
		if err := offlinesession.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf("db: validator failed for field \"user_id\": %w", err)}
		}
	}
	if v, ok := osuo.mutation.ConnID(); ok {
		if err := offlinesession.ConnIDValidator(v); err != nil {
			return &ValidationError{Name: "conn_id", err: fmt.Errorf("db: validator failed for field \"conn_id\": %w", err)}
		}
	}
	return nil
}

func (osuo *OfflineSessionUpdateOne) sqlSave(ctx context.Context) (_node *OfflineSession, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   offlinesession.Table,
			Columns: offlinesession.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: offlinesession.FieldID,
			},
		},
	}
	id, ok := osuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing OfflineSession.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := osuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, offlinesession.FieldID)
		for _, f := range fields {
			if !offlinesession.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != offlinesession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: offlinesession.FieldUserID,
		})
	}
	if value, ok := osuo.mutation.ConnID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: offlinesession.FieldConnID,
		})
	}
	if value, ok := osuo.mutation.Refresh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: offlinesession.FieldRefresh,
		})
	}
	if value, ok := osuo.mutation.ConnectorData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: offlinesession.FieldConnectorData,
		})
	}
	if osuo.mutation.ConnectorDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: offlinesession.FieldConnectorData,
		})
	}
	_node = &OfflineSession{config: osuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{offlinesession.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
