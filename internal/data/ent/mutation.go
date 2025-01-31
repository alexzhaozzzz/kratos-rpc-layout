// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/account"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/player"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccount = "Account"
	TypePlayer  = "Player"
)

// AccountMutation represents an operation that mutates the Account nodes in the graph.
type AccountMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	pass          *string
	email         *string
	phone         *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Account, error)
	predicates    []predicate.Account
}

var _ ent.Mutation = (*AccountMutation)(nil)

// accountOption allows management of the mutation configuration using functional options.
type accountOption func(*AccountMutation)

// newAccountMutation creates new mutation for the Account entity.
func newAccountMutation(c config, op Op, opts ...accountOption) *AccountMutation {
	m := &AccountMutation{
		config:        c,
		op:            op,
		typ:           TypeAccount,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccountID sets the ID field of the mutation.
func withAccountID(id int64) accountOption {
	return func(m *AccountMutation) {
		var (
			err   error
			once  sync.Once
			value *Account
		)
		m.oldValue = func(ctx context.Context) (*Account, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Account.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccount sets the old Account of the mutation.
func withAccount(node *Account) accountOption {
	return func(m *AccountMutation) {
		m.oldValue = func(context.Context) (*Account, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccountMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccountMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Account entities.
func (m *AccountMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AccountMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AccountMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Account.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetPass sets the "pass" field.
func (m *AccountMutation) SetPass(s string) {
	m.pass = &s
}

// Pass returns the value of the "pass" field in the mutation.
func (m *AccountMutation) Pass() (r string, exists bool) {
	v := m.pass
	if v == nil {
		return
	}
	return *v, true
}

// OldPass returns the old "pass" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldPass(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPass is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPass requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPass: %w", err)
	}
	return oldValue.Pass, nil
}

// ClearPass clears the value of the "pass" field.
func (m *AccountMutation) ClearPass() {
	m.pass = nil
	m.clearedFields[account.FieldPass] = struct{}{}
}

// PassCleared returns if the "pass" field was cleared in this mutation.
func (m *AccountMutation) PassCleared() bool {
	_, ok := m.clearedFields[account.FieldPass]
	return ok
}

// ResetPass resets all changes to the "pass" field.
func (m *AccountMutation) ResetPass() {
	m.pass = nil
	delete(m.clearedFields, account.FieldPass)
}

// SetEmail sets the "email" field.
func (m *AccountMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *AccountMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ClearEmail clears the value of the "email" field.
func (m *AccountMutation) ClearEmail() {
	m.email = nil
	m.clearedFields[account.FieldEmail] = struct{}{}
}

// EmailCleared returns if the "email" field was cleared in this mutation.
func (m *AccountMutation) EmailCleared() bool {
	_, ok := m.clearedFields[account.FieldEmail]
	return ok
}

// ResetEmail resets all changes to the "email" field.
func (m *AccountMutation) ResetEmail() {
	m.email = nil
	delete(m.clearedFields, account.FieldEmail)
}

// SetPhone sets the "phone" field.
func (m *AccountMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *AccountMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldPhone(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ClearPhone clears the value of the "phone" field.
func (m *AccountMutation) ClearPhone() {
	m.phone = nil
	m.clearedFields[account.FieldPhone] = struct{}{}
}

// PhoneCleared returns if the "phone" field was cleared in this mutation.
func (m *AccountMutation) PhoneCleared() bool {
	_, ok := m.clearedFields[account.FieldPhone]
	return ok
}

// ResetPhone resets all changes to the "phone" field.
func (m *AccountMutation) ResetPhone() {
	m.phone = nil
	delete(m.clearedFields, account.FieldPhone)
}

// Where appends a list predicates to the AccountMutation builder.
func (m *AccountMutation) Where(ps ...predicate.Account) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AccountMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AccountMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Account, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AccountMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AccountMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Account).
func (m *AccountMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccountMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.pass != nil {
		fields = append(fields, account.FieldPass)
	}
	if m.email != nil {
		fields = append(fields, account.FieldEmail)
	}
	if m.phone != nil {
		fields = append(fields, account.FieldPhone)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccountMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case account.FieldPass:
		return m.Pass()
	case account.FieldEmail:
		return m.Email()
	case account.FieldPhone:
		return m.Phone()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccountMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case account.FieldPass:
		return m.OldPass(ctx)
	case account.FieldEmail:
		return m.OldEmail(ctx)
	case account.FieldPhone:
		return m.OldPhone(ctx)
	}
	return nil, fmt.Errorf("unknown Account field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) SetField(name string, value ent.Value) error {
	switch name {
	case account.FieldPass:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPass(v)
		return nil
	case account.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case account.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccountMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccountMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Account numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccountMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(account.FieldPass) {
		fields = append(fields, account.FieldPass)
	}
	if m.FieldCleared(account.FieldEmail) {
		fields = append(fields, account.FieldEmail)
	}
	if m.FieldCleared(account.FieldPhone) {
		fields = append(fields, account.FieldPhone)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccountMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccountMutation) ClearField(name string) error {
	switch name {
	case account.FieldPass:
		m.ClearPass()
		return nil
	case account.FieldEmail:
		m.ClearEmail()
		return nil
	case account.FieldPhone:
		m.ClearPhone()
		return nil
	}
	return fmt.Errorf("unknown Account nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccountMutation) ResetField(name string) error {
	switch name {
	case account.FieldPass:
		m.ResetPass()
		return nil
	case account.FieldEmail:
		m.ResetEmail()
		return nil
	case account.FieldPhone:
		m.ResetPhone()
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccountMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccountMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccountMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccountMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccountMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccountMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccountMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Account unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccountMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Account edge %s", name)
}

// PlayerMutation represents an operation that mutates the Player nodes in the graph.
type PlayerMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	account_id    *int64
	addaccount_id *int64
	player_id     *int64
	addplayer_id  *int64
	nickname      *string
	gender        *int32
	addgender     *int32
	avatar        *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Player, error)
	predicates    []predicate.Player
}

var _ ent.Mutation = (*PlayerMutation)(nil)

// playerOption allows management of the mutation configuration using functional options.
type playerOption func(*PlayerMutation)

// newPlayerMutation creates new mutation for the Player entity.
func newPlayerMutation(c config, op Op, opts ...playerOption) *PlayerMutation {
	m := &PlayerMutation{
		config:        c,
		op:            op,
		typ:           TypePlayer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPlayerID sets the ID field of the mutation.
func withPlayerID(id int64) playerOption {
	return func(m *PlayerMutation) {
		var (
			err   error
			once  sync.Once
			value *Player
		)
		m.oldValue = func(ctx context.Context) (*Player, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Player.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPlayer sets the old Player of the mutation.
func withPlayer(node *Player) playerOption {
	return func(m *PlayerMutation) {
		m.oldValue = func(context.Context) (*Player, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PlayerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PlayerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Player entities.
func (m *PlayerMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PlayerMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PlayerMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Player.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAccountID sets the "account_id" field.
func (m *PlayerMutation) SetAccountID(i int64) {
	m.account_id = &i
	m.addaccount_id = nil
}

// AccountID returns the value of the "account_id" field in the mutation.
func (m *PlayerMutation) AccountID() (r int64, exists bool) {
	v := m.account_id
	if v == nil {
		return
	}
	return *v, true
}

// OldAccountID returns the old "account_id" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldAccountID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccountID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccountID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccountID: %w", err)
	}
	return oldValue.AccountID, nil
}

// AddAccountID adds i to the "account_id" field.
func (m *PlayerMutation) AddAccountID(i int64) {
	if m.addaccount_id != nil {
		*m.addaccount_id += i
	} else {
		m.addaccount_id = &i
	}
}

// AddedAccountID returns the value that was added to the "account_id" field in this mutation.
func (m *PlayerMutation) AddedAccountID() (r int64, exists bool) {
	v := m.addaccount_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetAccountID resets all changes to the "account_id" field.
func (m *PlayerMutation) ResetAccountID() {
	m.account_id = nil
	m.addaccount_id = nil
}

// SetPlayerID sets the "player_id" field.
func (m *PlayerMutation) SetPlayerID(i int64) {
	m.player_id = &i
	m.addplayer_id = nil
}

// PlayerID returns the value of the "player_id" field in the mutation.
func (m *PlayerMutation) PlayerID() (r int64, exists bool) {
	v := m.player_id
	if v == nil {
		return
	}
	return *v, true
}

// OldPlayerID returns the old "player_id" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldPlayerID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPlayerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPlayerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPlayerID: %w", err)
	}
	return oldValue.PlayerID, nil
}

// AddPlayerID adds i to the "player_id" field.
func (m *PlayerMutation) AddPlayerID(i int64) {
	if m.addplayer_id != nil {
		*m.addplayer_id += i
	} else {
		m.addplayer_id = &i
	}
}

// AddedPlayerID returns the value that was added to the "player_id" field in this mutation.
func (m *PlayerMutation) AddedPlayerID() (r int64, exists bool) {
	v := m.addplayer_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetPlayerID resets all changes to the "player_id" field.
func (m *PlayerMutation) ResetPlayerID() {
	m.player_id = nil
	m.addplayer_id = nil
}

// SetNickname sets the "nickname" field.
func (m *PlayerMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the value of the "nickname" field in the mutation.
func (m *PlayerMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldNickname returns the old "nickname" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldNickname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNickname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickname: %w", err)
	}
	return oldValue.Nickname, nil
}

// ClearNickname clears the value of the "nickname" field.
func (m *PlayerMutation) ClearNickname() {
	m.nickname = nil
	m.clearedFields[player.FieldNickname] = struct{}{}
}

// NicknameCleared returns if the "nickname" field was cleared in this mutation.
func (m *PlayerMutation) NicknameCleared() bool {
	_, ok := m.clearedFields[player.FieldNickname]
	return ok
}

// ResetNickname resets all changes to the "nickname" field.
func (m *PlayerMutation) ResetNickname() {
	m.nickname = nil
	delete(m.clearedFields, player.FieldNickname)
}

// SetGender sets the "gender" field.
func (m *PlayerMutation) SetGender(i int32) {
	m.gender = &i
	m.addgender = nil
}

// Gender returns the value of the "gender" field in the mutation.
func (m *PlayerMutation) Gender() (r int32, exists bool) {
	v := m.gender
	if v == nil {
		return
	}
	return *v, true
}

// OldGender returns the old "gender" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldGender(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGender is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGender requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGender: %w", err)
	}
	return oldValue.Gender, nil
}

// AddGender adds i to the "gender" field.
func (m *PlayerMutation) AddGender(i int32) {
	if m.addgender != nil {
		*m.addgender += i
	} else {
		m.addgender = &i
	}
}

// AddedGender returns the value that was added to the "gender" field in this mutation.
func (m *PlayerMutation) AddedGender() (r int32, exists bool) {
	v := m.addgender
	if v == nil {
		return
	}
	return *v, true
}

// ClearGender clears the value of the "gender" field.
func (m *PlayerMutation) ClearGender() {
	m.gender = nil
	m.addgender = nil
	m.clearedFields[player.FieldGender] = struct{}{}
}

// GenderCleared returns if the "gender" field was cleared in this mutation.
func (m *PlayerMutation) GenderCleared() bool {
	_, ok := m.clearedFields[player.FieldGender]
	return ok
}

// ResetGender resets all changes to the "gender" field.
func (m *PlayerMutation) ResetGender() {
	m.gender = nil
	m.addgender = nil
	delete(m.clearedFields, player.FieldGender)
}

// SetAvatar sets the "avatar" field.
func (m *PlayerMutation) SetAvatar(s string) {
	m.avatar = &s
}

// Avatar returns the value of the "avatar" field in the mutation.
func (m *PlayerMutation) Avatar() (r string, exists bool) {
	v := m.avatar
	if v == nil {
		return
	}
	return *v, true
}

// OldAvatar returns the old "avatar" field's value of the Player entity.
// If the Player object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PlayerMutation) OldAvatar(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAvatar is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAvatar requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvatar: %w", err)
	}
	return oldValue.Avatar, nil
}

// ClearAvatar clears the value of the "avatar" field.
func (m *PlayerMutation) ClearAvatar() {
	m.avatar = nil
	m.clearedFields[player.FieldAvatar] = struct{}{}
}

// AvatarCleared returns if the "avatar" field was cleared in this mutation.
func (m *PlayerMutation) AvatarCleared() bool {
	_, ok := m.clearedFields[player.FieldAvatar]
	return ok
}

// ResetAvatar resets all changes to the "avatar" field.
func (m *PlayerMutation) ResetAvatar() {
	m.avatar = nil
	delete(m.clearedFields, player.FieldAvatar)
}

// Where appends a list predicates to the PlayerMutation builder.
func (m *PlayerMutation) Where(ps ...predicate.Player) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PlayerMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PlayerMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Player, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PlayerMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PlayerMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Player).
func (m *PlayerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PlayerMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.account_id != nil {
		fields = append(fields, player.FieldAccountID)
	}
	if m.player_id != nil {
		fields = append(fields, player.FieldPlayerID)
	}
	if m.nickname != nil {
		fields = append(fields, player.FieldNickname)
	}
	if m.gender != nil {
		fields = append(fields, player.FieldGender)
	}
	if m.avatar != nil {
		fields = append(fields, player.FieldAvatar)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PlayerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case player.FieldAccountID:
		return m.AccountID()
	case player.FieldPlayerID:
		return m.PlayerID()
	case player.FieldNickname:
		return m.Nickname()
	case player.FieldGender:
		return m.Gender()
	case player.FieldAvatar:
		return m.Avatar()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PlayerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case player.FieldAccountID:
		return m.OldAccountID(ctx)
	case player.FieldPlayerID:
		return m.OldPlayerID(ctx)
	case player.FieldNickname:
		return m.OldNickname(ctx)
	case player.FieldGender:
		return m.OldGender(ctx)
	case player.FieldAvatar:
		return m.OldAvatar(ctx)
	}
	return nil, fmt.Errorf("unknown Player field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlayerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case player.FieldAccountID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccountID(v)
		return nil
	case player.FieldPlayerID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPlayerID(v)
		return nil
	case player.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case player.FieldGender:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGender(v)
		return nil
	case player.FieldAvatar:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvatar(v)
		return nil
	}
	return fmt.Errorf("unknown Player field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PlayerMutation) AddedFields() []string {
	var fields []string
	if m.addaccount_id != nil {
		fields = append(fields, player.FieldAccountID)
	}
	if m.addplayer_id != nil {
		fields = append(fields, player.FieldPlayerID)
	}
	if m.addgender != nil {
		fields = append(fields, player.FieldGender)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PlayerMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case player.FieldAccountID:
		return m.AddedAccountID()
	case player.FieldPlayerID:
		return m.AddedPlayerID()
	case player.FieldGender:
		return m.AddedGender()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PlayerMutation) AddField(name string, value ent.Value) error {
	switch name {
	case player.FieldAccountID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAccountID(v)
		return nil
	case player.FieldPlayerID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPlayerID(v)
		return nil
	case player.FieldGender:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGender(v)
		return nil
	}
	return fmt.Errorf("unknown Player numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PlayerMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(player.FieldNickname) {
		fields = append(fields, player.FieldNickname)
	}
	if m.FieldCleared(player.FieldGender) {
		fields = append(fields, player.FieldGender)
	}
	if m.FieldCleared(player.FieldAvatar) {
		fields = append(fields, player.FieldAvatar)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PlayerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PlayerMutation) ClearField(name string) error {
	switch name {
	case player.FieldNickname:
		m.ClearNickname()
		return nil
	case player.FieldGender:
		m.ClearGender()
		return nil
	case player.FieldAvatar:
		m.ClearAvatar()
		return nil
	}
	return fmt.Errorf("unknown Player nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PlayerMutation) ResetField(name string) error {
	switch name {
	case player.FieldAccountID:
		m.ResetAccountID()
		return nil
	case player.FieldPlayerID:
		m.ResetPlayerID()
		return nil
	case player.FieldNickname:
		m.ResetNickname()
		return nil
	case player.FieldGender:
		m.ResetGender()
		return nil
	case player.FieldAvatar:
		m.ResetAvatar()
		return nil
	}
	return fmt.Errorf("unknown Player field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PlayerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PlayerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PlayerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PlayerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PlayerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PlayerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PlayerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Player unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PlayerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Player edge %s", name)
}
