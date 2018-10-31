package model

import (
	"time"
)

// User is a user of the system and can also be the owner of 
// a property that the emeters belongs to.
type User struct {
	ID string `json:"-"`

	Name       string      `json:"name"`
	UserID     string      `json:"userid"`
	Properties []*Property `json:"-"`
}

// Property can be a home or other building the user owns.
type Property struct {
	ID string `json:"-"`

	Name    string    `json:"name"`
	EMeters []*EMeter `json:"-"`
}

// EMeter is the meter that the measurement came from.
type EMeter struct {
	ID string `json:"-"`

	Model        *EMeterModel   `json:"-"`
	Measurements []*Measurement `json:"-"`
}

// EMeterModel is the model of the emeter.
type EMeterModel struct {
	ID string `json:"-"`

	Manufacturer *Manufacturer `json:"-"`
	Meters       []*EMeter     `json:"-"`
}

// Measurement represents the measured readings for an energy meter at
// a given time.
type Measurement struct {
	ID string `json:"-"`

	Time     *time.Time `json:"time"`
	Readings []*Reading `json:"readings"`
}

// Reading is a value struct, does not have an identifier and can only 
// exist as part of a Measurement.
type Reading struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Manufacturer is the producer of the emeter.
type Manufacturer struct {
	ID string `json:"-"`

	Models []*EMeterModel `json:"-"`
}

// GetID returns the ID of the user.
func (u *User) GetID() string {
	return u.ID
}

// SetID sets the id of the user.
func (u *User) SetID(id string) error {
	u.ID = id
	return nil
}

// GetName returns the name, not required in this case for users would be the default name.
func (u *User) GetName() string {
	return "users"
}

func (p *Property) GetID() string {
	return p.ID
}

func (p *Property) SetID(id string) error {
	p.ID = id
	return nil
}

// GetName returns the name, not required in this case for Propertys would be the default name.
func (p *Property) GetName() string {
	return "properties"
}

func (e *EMeter) GetID() string {
	return e.ID
}

func (e *EMeter) SetID(id string) error {
	e.ID = id
	return nil
}

// GetName returns the name, not required in this case for EMeters would be the default name.
func (e *EMeter) GetName() string {
	return "properties"
}

func (m *Manufacturer) GetID() string {
	return m.ID
}

func (m *Manufacturer) SetID(id string) error {
	m.ID = id
	return nil
}

// GetName returns the name, not required in this case for Manufacturers would be the default name.
func (m *Manufacturer) GetName() string {
	return "properties"
}

func (e *EMeterModel) GetID() string {
	return e.ID
}

func (e *EMeterModel) SetID(id string) error {
	e.ID = id
	return nil
}

// GetName returns the name, not required in this case for EMeterModels would be the default name.
func (e *EMeterModel) GetName() string {
	return "properties"
}

func (e *Measurement) GetID() string {
	return e.ID
}

func (e *Measurement) SetID(id string) error {
	e.ID = id
	return nil
}

// GetName returns the name, not required in this case for Measurements would be the default name.
func (e *Measurement) GetName() string {
	return "properties"
}
