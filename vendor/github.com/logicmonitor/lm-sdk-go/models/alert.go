// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Alert alert
// swagger:model Alert
type Alert struct {

	// The active SDT, if one exists
	// Read Only: true
	SDT interface{} `json:"SDT,omitempty"`

	// The comment submitted with the acknowledgement
	// Read Only: true
	AckComment string `json:"ackComment,omitempty"`

	// Whether or not the alert has been acknowledged
	// Read Only: true
	Acked *bool `json:"acked,omitempty"`

	// The user that acknowledged the alert
	// Read Only: true
	AckedBy string `json:"ackedBy,omitempty"`

	// The time (in epoch format) that the alert was acknowledged
	// Read Only: true
	AckedEpoch int64 `json:"ackedEpoch,omitempty"`

	// The value that triggered the alert
	// Read Only: true
	AlertValue string `json:"alertValue,omitempty"`

	// The escalation chain the alert was routed to
	// Read Only: true
	Chain string `json:"chain,omitempty"`

	// The id of the escalation chain the alert was routed to
	// Read Only: true
	ChainID int32 `json:"chainId,omitempty"`

	// The value that cleared the alert
	// Read Only: true
	ClearValue string `json:"clearValue,omitempty"`

	// Whether or not the alert has cleared
	// Read Only: true
	Cleared *bool `json:"cleared,omitempty"`

	// The property or token values that should display with the alert details. Note that if referencing tokens, you'll need to URL encode the # symbol.
	// Read Only: true
	CustomColumns interface{} `json:"customColumns,omitempty"`

	// The id of the datapoint in alert
	// Read Only: true
	DataPointID int32 `json:"dataPointId,omitempty"`

	// The name of the datapoint in alert
	// Read Only: true
	DataPointName string `json:"dataPointName,omitempty"`

	// The alert message, if needMessage=true is included in the query parameters
	// Read Only: true
	DetailMessage interface{} `json:"detailMessage,omitempty"`

	// The time (in epoch format) that the alert ended
	// Read Only: true
	EndEpoch int64 `json:"endEpoch,omitempty"`

	// The alert id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The description of the instance in alert
	// Read Only: true
	InstanceDescription string `json:"instanceDescription,omitempty"`

	// The id of the instance in alert
	// Read Only: true
	InstanceID int32 `json:"instanceId,omitempty"`

	// The name of the instance in alert
	// Read Only: true
	InstanceName string `json:"instanceName,omitempty"`

	// The internal id for the alert
	// Read Only: true
	InternalID string `json:"internalId,omitempty"`

	// Information about the groups the object is a member of
	// Read Only: true
	MonitorObjectGroups interface{} `json:"monitorObjectGroups,omitempty"`

	// The id of the object that the alert is associated with
	// Read Only: true
	MonitorObjectID int32 `json:"monitorObjectId,omitempty"`

	// The name of the object that the alert is associated with
	// Read Only: true
	MonitorObjectName string `json:"monitorObjectName,omitempty"`

	// monitor object type
	// Read Only: true
	MonitorObjectType string `json:"monitorObjectType,omitempty"`

	// The next recipient in the escalation chain for this alert
	// Read Only: true
	NextRecipient int32 `json:"nextRecipient,omitempty"`

	// The recipients that have received the alert
	// Read Only: true
	ReceivedList string `json:"receivedList,omitempty"`

	// The device specific LogicModule Id
	// Read Only: true
	ResourceID int32 `json:"resourceId,omitempty"`

	// The id of the datasource in alert
	// Read Only: true
	ResourceTemplateID int32 `json:"resourceTemplateId,omitempty"`

	// The name of the datasource in alert
	// Read Only: true
	ResourceTemplateName string `json:"resourceTemplateName,omitempty"`

	// The type of the logicmodule in alert
	// Read Only: true
	ResourceTemplateType string `json:"resourceTemplateType,omitempty"`

	// The rule the alert matches
	// Read Only: true
	Rule string `json:"rule,omitempty"`

	// The id of the rule the alert matches
	// Read Only: true
	RuleID int32 `json:"ruleId,omitempty"`

	// Whether or not the alert was triggered during an SDT
	// Read Only: true
	SDTED *bool `json:"sdted,omitempty"`

	// The alert severity, where 2=warning, 3=error and 4=critical
	// Read Only: true
	Severity int32 `json:"severity,omitempty"`

	// The time (in epoch format) that the alert started
	// Read Only: true
	StartEpoch int64 `json:"startEpoch,omitempty"`

	// The id of the sub time based chain
	// Read Only: true
	SubChainID int32 `json:"subChainId,omitempty"`

	// The threshold associated with the object in alert
	// Read Only: true
	Threshold string `json:"threshold,omitempty"`

	// The type of alert
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this alert
func (m *Alert) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alert) UnmarshalBinary(b []byte) error {
	var res Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
