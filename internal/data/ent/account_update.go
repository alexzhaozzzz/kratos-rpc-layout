// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/account"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/predicate"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetPass sets the "pass" field.
func (au *AccountUpdate) SetPass(s string) *AccountUpdate {
	au.mutation.SetPass(s)
	return au
}

// SetNillablePass sets the "pass" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePass(s *string) *AccountUpdate {
	if s != nil {
		au.SetPass(*s)
	}
	return au
}

// ClearPass clears the value of the "pass" field.
func (au *AccountUpdate) ClearPass() *AccountUpdate {
	au.mutation.ClearPass()
	return au
}

// SetEmail sets the "email" field.
func (au *AccountUpdate) SetEmail(s string) *AccountUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (au *AccountUpdate) SetNillableEmail(s *string) *AccountUpdate {
	if s != nil {
		au.SetEmail(*s)
	}
	return au
}

// ClearEmail clears the value of the "email" field.
func (au *AccountUpdate) ClearEmail() *AccountUpdate {
	au.mutation.ClearEmail()
	return au
}

// SetPhone sets the "phone" field.
func (au *AccountUpdate) SetPhone(s string) *AccountUpdate {
	au.mutation.SetPhone(s)
	return au
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePhone(s *string) *AccountUpdate {
	if s != nil {
		au.SetPhone(*s)
	}
	return au
}

// ClearPhone clears the value of the "phone" field.
func (au *AccountUpdate) ClearPhone() *AccountUpdate {
	au.mutation.ClearPhone()
	return au
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Pass(); ok {
		_spec.SetField(account.FieldPass, field.TypeString, value)
	}
	if au.mutation.PassCleared() {
		_spec.ClearField(account.FieldPass, field.TypeString)
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
	}
	if au.mutation.EmailCleared() {
		_spec.ClearField(account.FieldEmail, field.TypeString)
	}
	if value, ok := au.mutation.Phone(); ok {
		_spec.SetField(account.FieldPhone, field.TypeString, value)
	}
	if au.mutation.PhoneCleared() {
		_spec.ClearField(account.FieldPhone, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetPass sets the "pass" field.
func (auo *AccountUpdateOne) SetPass(s string) *AccountUpdateOne {
	auo.mutation.SetPass(s)
	return auo
}

// SetNillablePass sets the "pass" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePass(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPass(*s)
	}
	return auo
}

// ClearPass clears the value of the "pass" field.
func (auo *AccountUpdateOne) ClearPass() *AccountUpdateOne {
	auo.mutation.ClearPass()
	return auo
}

// SetEmail sets the "email" field.
func (auo *AccountUpdateOne) SetEmail(s string) *AccountUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableEmail(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetEmail(*s)
	}
	return auo
}

// ClearEmail clears the value of the "email" field.
func (auo *AccountUpdateOne) ClearEmail() *AccountUpdateOne {
	auo.mutation.ClearEmail()
	return auo
}

// SetPhone sets the "phone" field.
func (auo *AccountUpdateOne) SetPhone(s string) *AccountUpdateOne {
	auo.mutation.SetPhone(s)
	return auo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePhone(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPhone(*s)
	}
	return auo
}

// ClearPhone clears the value of the "phone" field.
func (auo *AccountUpdateOne) ClearPhone() *AccountUpdateOne {
	auo.mutation.ClearPhone()
	return auo
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Pass(); ok {
		_spec.SetField(account.FieldPass, field.TypeString, value)
	}
	if auo.mutation.PassCleared() {
		_spec.ClearField(account.FieldPass, field.TypeString)
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
	}
	if auo.mutation.EmailCleared() {
		_spec.ClearField(account.FieldEmail, field.TypeString)
	}
	if value, ok := auo.mutation.Phone(); ok {
		_spec.SetField(account.FieldPhone, field.TypeString, value)
	}
	if auo.mutation.PhoneCleared() {
		_spec.ClearField(account.FieldPhone, field.TypeString)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
