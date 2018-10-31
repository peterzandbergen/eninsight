package model

import (
	"testing"

	"github.com/manyminds/api2go/jsonapi"
)

func TestUserInterfaces(t *testing.T) {
	var _ jsonapi.MarshalIdentifier = &User{}
	var _ jsonapi.UnmarshalIdentifier = &User{}
	var _ jsonapi.EntityNamer = &User{}
}

func TestPropertyInterfaces(t *testing.T) {
	var _ jsonapi.MarshalIdentifier = &Property{}
	var _ jsonapi.UnmarshalIdentifier = &Property{}
	var _ jsonapi.EntityNamer = &Property{}
}

func TestEMeterInterfaces(t *testing.T) {
	var _ jsonapi.MarshalIdentifier = &EMeter{}
	var _ jsonapi.UnmarshalIdentifier = &EMeter{}
	var _ jsonapi.EntityNamer = &EMeter{}
}

func TestEMeterModelInterfaces(t *testing.T) {
	var _ jsonapi.MarshalIdentifier = &EMeterModel{}
	var _ jsonapi.UnmarshalIdentifier = &EMeterModel{}
	var _ jsonapi.EntityNamer = &EMeterModel{}
}

func TestMeasurementInterfaces(t *testing.T) {
	var _ jsonapi.MarshalIdentifier = &Measurement{}
	var _ jsonapi.UnmarshalIdentifier = &Measurement{}
	var _ jsonapi.EntityNamer = &Measurement{}
}

func TestManufacturerInterfaces(t *testing.T) {
	var _ jsonapi.MarshalIdentifier = &Manufacturer{}
	var _ jsonapi.UnmarshalIdentifier = &Manufacturer{}
	var _ jsonapi.EntityNamer = &Manufacturer{}
}
