// Code generated by dto builder generator; DO NOT EDIT.

package sdk

import ()

func NewShowApplicationRoleRequest() *ShowApplicationRoleRequest {
	return &ShowApplicationRoleRequest{}
}

func (s *ShowApplicationRoleRequest) WithApplicationName(ApplicationName AccountObjectIdentifier) *ShowApplicationRoleRequest {
	s.ApplicationName = ApplicationName
	return s
}

func (s *ShowApplicationRoleRequest) WithLimit(Limit *LimitFrom) *ShowApplicationRoleRequest {
	s.Limit = Limit
	return s
}

func NewShowByIDApplicationRoleRequest(
	name DatabaseObjectIdentifier,
	ApplicationName AccountObjectIdentifier,
) *ShowByIDApplicationRoleRequest {
	s := ShowByIDApplicationRoleRequest{}
	s.name = name
	s.ApplicationName = ApplicationName
	return &s
}