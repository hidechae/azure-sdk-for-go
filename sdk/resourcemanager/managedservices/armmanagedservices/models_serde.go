//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagedservices

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Authorization.
func (a Authorization) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "delegatedRoleDefinitionIds", a.DelegatedRoleDefinitionIDs)
	populate(objectMap, "principalId", a.PrincipalID)
	populate(objectMap, "principalIdDisplayName", a.PrincipalIDDisplayName)
	populate(objectMap, "roleDefinitionId", a.RoleDefinitionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Authorization.
func (a *Authorization) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "delegatedRoleDefinitionIds":
			err = unpopulate(val, "DelegatedRoleDefinitionIDs", &a.DelegatedRoleDefinitionIDs)
			delete(rawMsg, key)
		case "principalId":
			err = unpopulate(val, "PrincipalID", &a.PrincipalID)
			delete(rawMsg, key)
		case "principalIdDisplayName":
			err = unpopulate(val, "PrincipalIDDisplayName", &a.PrincipalIDDisplayName)
			delete(rawMsg, key)
		case "roleDefinitionId":
			err = unpopulate(val, "RoleDefinitionID", &a.RoleDefinitionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EligibleApprover.
func (e EligibleApprover) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "principalId", e.PrincipalID)
	populate(objectMap, "principalIdDisplayName", e.PrincipalIDDisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EligibleApprover.
func (e *EligibleApprover) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &e.PrincipalID)
			delete(rawMsg, key)
		case "principalIdDisplayName":
			err = unpopulate(val, "PrincipalIDDisplayName", &e.PrincipalIDDisplayName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EligibleAuthorization.
func (e EligibleAuthorization) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "justInTimeAccessPolicy", e.JustInTimeAccessPolicy)
	populate(objectMap, "principalId", e.PrincipalID)
	populate(objectMap, "principalIdDisplayName", e.PrincipalIDDisplayName)
	populate(objectMap, "roleDefinitionId", e.RoleDefinitionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EligibleAuthorization.
func (e *EligibleAuthorization) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "justInTimeAccessPolicy":
			err = unpopulate(val, "JustInTimeAccessPolicy", &e.JustInTimeAccessPolicy)
			delete(rawMsg, key)
		case "principalId":
			err = unpopulate(val, "PrincipalID", &e.PrincipalID)
			delete(rawMsg, key)
		case "principalIdDisplayName":
			err = unpopulate(val, "PrincipalIDDisplayName", &e.PrincipalIDDisplayName)
			delete(rawMsg, key)
		case "roleDefinitionId":
			err = unpopulate(val, "RoleDefinitionID", &e.RoleDefinitionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDefinition.
func (e ErrorDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDefinition.
func (e *ErrorDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JustInTimeAccessPolicy.
func (j JustInTimeAccessPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "managedByTenantApprovers", j.ManagedByTenantApprovers)
	populate(objectMap, "maximumActivationDuration", j.MaximumActivationDuration)
	populate(objectMap, "multiFactorAuthProvider", j.MultiFactorAuthProvider)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JustInTimeAccessPolicy.
func (j *JustInTimeAccessPolicy) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", j, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "managedByTenantApprovers":
			err = unpopulate(val, "ManagedByTenantApprovers", &j.ManagedByTenantApprovers)
			delete(rawMsg, key)
		case "maximumActivationDuration":
			err = unpopulate(val, "MaximumActivationDuration", &j.MaximumActivationDuration)
			delete(rawMsg, key)
		case "multiFactorAuthProvider":
			err = unpopulate(val, "MultiFactorAuthProvider", &j.MultiFactorAuthProvider)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", j, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MarketplaceRegistrationDefinition.
func (m MarketplaceRegistrationDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "plan", m.Plan)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MarketplaceRegistrationDefinition.
func (m *MarketplaceRegistrationDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "plan":
			err = unpopulate(val, "Plan", &m.Plan)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &m.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MarketplaceRegistrationDefinitionList.
func (m MarketplaceRegistrationDefinitionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MarketplaceRegistrationDefinitionList.
func (m *MarketplaceRegistrationDefinitionList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &m.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MarketplaceRegistrationDefinitionProperties.
func (m MarketplaceRegistrationDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "authorizations", m.Authorizations)
	populate(objectMap, "eligibleAuthorizations", m.EligibleAuthorizations)
	populate(objectMap, "managedByTenantId", m.ManagedByTenantID)
	populate(objectMap, "offerDisplayName", m.OfferDisplayName)
	populate(objectMap, "planDisplayName", m.PlanDisplayName)
	populate(objectMap, "publisherDisplayName", m.PublisherDisplayName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MarketplaceRegistrationDefinitionProperties.
func (m *MarketplaceRegistrationDefinitionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authorizations":
			err = unpopulate(val, "Authorizations", &m.Authorizations)
			delete(rawMsg, key)
		case "eligibleAuthorizations":
			err = unpopulate(val, "EligibleAuthorizations", &m.EligibleAuthorizations)
			delete(rawMsg, key)
		case "managedByTenantId":
			err = unpopulate(val, "ManagedByTenantID", &m.ManagedByTenantID)
			delete(rawMsg, key)
		case "offerDisplayName":
			err = unpopulate(val, "OfferDisplayName", &m.OfferDisplayName)
			delete(rawMsg, key)
		case "planDisplayName":
			err = unpopulate(val, "PlanDisplayName", &m.PlanDisplayName)
			delete(rawMsg, key)
		case "publisherDisplayName":
			err = unpopulate(val, "PublisherDisplayName", &m.PublisherDisplayName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Operation.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationList.
func (o *OperationList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Plan.
func (p Plan) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "product", p.Product)
	populate(objectMap, "publisher", p.Publisher)
	populate(objectMap, "version", p.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Plan.
func (p *Plan) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "product":
			err = unpopulate(val, "Product", &p.Product)
			delete(rawMsg, key)
		case "publisher":
			err = unpopulate(val, "Publisher", &p.Publisher)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &p.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationAssignment.
func (r RegistrationAssignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationAssignment.
func (r *RegistrationAssignment) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &r.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationAssignmentList.
func (r RegistrationAssignmentList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationAssignmentList.
func (r *RegistrationAssignmentList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &r.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &r.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationAssignmentProperties.
func (r RegistrationAssignmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "registrationDefinition", r.RegistrationDefinition)
	populate(objectMap, "registrationDefinitionId", r.RegistrationDefinitionID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationAssignmentProperties.
func (r *RegistrationAssignmentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
			delete(rawMsg, key)
		case "registrationDefinition":
			err = unpopulate(val, "RegistrationDefinition", &r.RegistrationDefinition)
			delete(rawMsg, key)
		case "registrationDefinitionId":
			err = unpopulate(val, "RegistrationDefinitionID", &r.RegistrationDefinitionID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationAssignmentPropertiesRegistrationDefinition.
func (r RegistrationAssignmentPropertiesRegistrationDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "plan", r.Plan)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationAssignmentPropertiesRegistrationDefinition.
func (r *RegistrationAssignmentPropertiesRegistrationDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "plan":
			err = unpopulate(val, "Plan", &r.Plan)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &r.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationAssignmentPropertiesRegistrationDefinitionProperties.
func (r RegistrationAssignmentPropertiesRegistrationDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "authorizations", r.Authorizations)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "eligibleAuthorizations", r.EligibleAuthorizations)
	populate(objectMap, "managedByTenantId", r.ManagedByTenantID)
	populate(objectMap, "managedByTenantName", r.ManagedByTenantName)
	populate(objectMap, "manageeTenantId", r.ManageeTenantID)
	populate(objectMap, "manageeTenantName", r.ManageeTenantName)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "registrationDefinitionName", r.RegistrationDefinitionName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationAssignmentPropertiesRegistrationDefinitionProperties.
func (r *RegistrationAssignmentPropertiesRegistrationDefinitionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authorizations":
			err = unpopulate(val, "Authorizations", &r.Authorizations)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &r.Description)
			delete(rawMsg, key)
		case "eligibleAuthorizations":
			err = unpopulate(val, "EligibleAuthorizations", &r.EligibleAuthorizations)
			delete(rawMsg, key)
		case "managedByTenantId":
			err = unpopulate(val, "ManagedByTenantID", &r.ManagedByTenantID)
			delete(rawMsg, key)
		case "managedByTenantName":
			err = unpopulate(val, "ManagedByTenantName", &r.ManagedByTenantName)
			delete(rawMsg, key)
		case "manageeTenantId":
			err = unpopulate(val, "ManageeTenantID", &r.ManageeTenantID)
			delete(rawMsg, key)
		case "manageeTenantName":
			err = unpopulate(val, "ManageeTenantName", &r.ManageeTenantName)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
			delete(rawMsg, key)
		case "registrationDefinitionName":
			err = unpopulate(val, "RegistrationDefinitionName", &r.RegistrationDefinitionName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationDefinition.
func (r RegistrationDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "plan", r.Plan)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationDefinition.
func (r *RegistrationDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "plan":
			err = unpopulate(val, "Plan", &r.Plan)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &r.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationDefinitionList.
func (r RegistrationDefinitionList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationDefinitionList.
func (r *RegistrationDefinitionList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &r.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &r.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationDefinitionProperties.
func (r RegistrationDefinitionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "authorizations", r.Authorizations)
	populate(objectMap, "description", r.Description)
	populate(objectMap, "eligibleAuthorizations", r.EligibleAuthorizations)
	populate(objectMap, "managedByTenantId", r.ManagedByTenantID)
	populate(objectMap, "managedByTenantName", r.ManagedByTenantName)
	populate(objectMap, "manageeTenantId", r.ManageeTenantID)
	populate(objectMap, "manageeTenantName", r.ManageeTenantName)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "registrationDefinitionName", r.RegistrationDefinitionName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationDefinitionProperties.
func (r *RegistrationDefinitionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authorizations":
			err = unpopulate(val, "Authorizations", &r.Authorizations)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &r.Description)
			delete(rawMsg, key)
		case "eligibleAuthorizations":
			err = unpopulate(val, "EligibleAuthorizations", &r.EligibleAuthorizations)
			delete(rawMsg, key)
		case "managedByTenantId":
			err = unpopulate(val, "ManagedByTenantID", &r.ManagedByTenantID)
			delete(rawMsg, key)
		case "managedByTenantName":
			err = unpopulate(val, "ManagedByTenantName", &r.ManagedByTenantName)
			delete(rawMsg, key)
		case "manageeTenantId":
			err = unpopulate(val, "ManageeTenantID", &r.ManageeTenantID)
			delete(rawMsg, key)
		case "manageeTenantName":
			err = unpopulate(val, "ManageeTenantName", &r.ManageeTenantName)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
			delete(rawMsg, key)
		case "registrationDefinitionName":
			err = unpopulate(val, "RegistrationDefinitionName", &r.RegistrationDefinitionName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
