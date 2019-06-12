// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PostInstancesParamsBodyVolumeAttachmentsItemsVolumeEncryptionKey VolumeEncryptionKey
//
// The encryption key for the volume
// swagger:model postInstancesParamsBodyVolumeAttachmentsItemsVolumeEncryptionKey
type PostInstancesParamsBodyVolumeAttachmentsItemsVolumeEncryptionKey struct {

	// The CRN for the volume encryption key
	Crn string `json:"crn,omitempty"`
}

// Validate validates this post instances params body volume attachments items volume encryption key
func (m *PostInstancesParamsBodyVolumeAttachmentsItemsVolumeEncryptionKey) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostInstancesParamsBodyVolumeAttachmentsItemsVolumeEncryptionKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostInstancesParamsBodyVolumeAttachmentsItemsVolumeEncryptionKey) UnmarshalBinary(b []byte) error {
	var res PostInstancesParamsBodyVolumeAttachmentsItemsVolumeEncryptionKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}