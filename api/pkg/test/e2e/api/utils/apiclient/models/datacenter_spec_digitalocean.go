// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatacenterSpecDigitalocean DatacenterSpecDigitalocean describes a DigitalOcean datacenter
// swagger:model DatacenterSpecDigitalocean
type DatacenterSpecDigitalocean struct {

	// Datacenter location, e.g. "ams3". A list of existing datacenters can be found
	// at https://www.digitalocean.com/docs/platform/availability-matrix/
	Region string `json:"region,omitempty"`
}

// Validate validates this datacenter spec digitalocean
func (m *DatacenterSpecDigitalocean) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpecDigitalocean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpecDigitalocean) UnmarshalBinary(b []byte) error {
	var res DatacenterSpecDigitalocean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
