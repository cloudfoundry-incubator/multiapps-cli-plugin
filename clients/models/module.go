// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Module module
// swagger:model Module

type Module struct {

	// app name
	AppName string `json:"appName,omitempty"`

	// created on
	CreatedOn strfmt.Date `json:"createdOn,omitempty"`

	// module name
	ModuleName string `json:"moduleName,omitempty"`

	// provided dependency names
	ProvidedDependencyNames []string `json:"providedDependencyNames"`

	// services
	Services []string `json:"services"`

	// updated on
	UpdatedOn strfmt.Date `json:"updatedOn,omitempty"`

	// uris
	Uris []string `json:"uris"`
}

/* polymorph Module appName false */

/* polymorph Module createdOn false */

/* polymorph Module moduleName false */

/* polymorph Module providedDependencyNames false */

/* polymorph Module services false */

/* polymorph Module updatedOn false */

/* polymorph Module uris false */

// Validate validates this module
func (m *Module) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvidedDependencyNames(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUris(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Module) validateProvidedDependencyNames(formats strfmt.Registry) error {

	if swag.IsZero(m.ProvidedDependencyNames) { // not required
		return nil
	}

	return nil
}

func (m *Module) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	return nil
}

func (m *Module) validateUris(formats strfmt.Registry) error {

	if swag.IsZero(m.Uris) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Module) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Module) UnmarshalBinary(b []byte) error {
	var res Module
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
