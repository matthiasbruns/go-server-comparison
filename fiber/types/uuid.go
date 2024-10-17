package types

import (
	"bytes"
	"reflect"

	"github.com/google/uuid"
)

func UUIDNew() UUID {
	return UUID(uuid.New())
}

func UUIDParse(v string) (UUID, error) {
	id, err := uuid.Parse(v)
	if err != nil {
		return UUID{}, err
	}
	return UUID(id), nil
}

func UUIDMustParse(v string) UUID {
	return UUID(uuid.MustParse(v))
}

type UUID uuid.UUID

// ToUUID returns the underlying UUID type.
func (u UUID) ToUUID() uuid.UUID {
	return uuid.UUID(u)
}

// NotNil tests whether the UUID is not nil (instance) or a Nil-UUID.
// Use this with pointer types, because it is nil-pointer safe.
func (u *UUID) NotNil() bool {
	return u != nil && u.NotZero()
}

// NotZero tests whether the UUID is not a Nil-UUID.
func (u UUID) NotZero() bool {
	return !reflect.ValueOf(u).IsZero()
}

// String returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx , or "" if uuid is invalid.
func (u UUID) String() string {
	return u.ToUUID().String()
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (u UUID) MarshalBinary() ([]byte, error) {
	return u.ToUUID().MarshalBinary()
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (u *UUID) UnmarshalBinary(data []byte) error {
	return (*uuid.UUID)(u).UnmarshalBinary(data)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (u UUID) MarshalText() ([]byte, error) {
	return u.ToUUID().MarshalText()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (u *UUID) UnmarshalText(data []byte) error {
	return (*uuid.UUID)(u).UnmarshalText(data)
}

// MarshalJSON marshals the UUID into a JSON string.
func (u UUID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + u.ToUUID().String() + `"`), nil
}

// UnmarshalJSON reverts the logic of MarshalJSON.
func (u *UUID) UnmarshalJSON(data []byte) error {
	return (*uuid.UUID)(u).UnmarshalText(bytes.Trim(data, `"`))
}
