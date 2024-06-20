// Code generated by ent, DO NOT EDIT.

package player

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the player type in the database.
	Label = "player"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldPlayerID holds the string denoting the player_id field in the database.
	FieldPlayerID = "player_id"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// Table holds the table name of the player in the database.
	Table = "player"
)

// Columns holds all SQL columns for player fields.
var Columns = []string{
	FieldID,
	FieldAccountID,
	FieldPlayerID,
	FieldNickname,
	FieldGender,
	FieldAvatar,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultGender holds the default value on creation for the "gender" field.
	DefaultGender int32
)

// OrderOption defines the ordering options for the Player queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByPlayerID orders the results by the player_id field.
func ByPlayerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlayerID, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}