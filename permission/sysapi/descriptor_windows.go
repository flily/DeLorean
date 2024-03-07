//go:build windows

package sysapi

import (
	"golang.org/x/sys/windows"
)

type SecurityInformation uint32

const (
	OwnerSecurityInformation           SecurityInformation = 0x00000001
	GroupSecurityInformation           SecurityInformation = 0x00000002
	DaclSecurityInformation            SecurityInformation = 0x00000004
	SaclSecurityInformation            SecurityInformation = 0x00000008
	LabelSecurityInformation           SecurityInformation = 0x00000010
	AttributeSecurityInformation       SecurityInformation = 0x00000020
	ScopeSecurityInformation           SecurityInformation = 0x00000040
	BackupSecurityInformation          SecurityInformation = 0x00010000
	ProtectedDaclSecurityInformation   SecurityInformation = 0x80000000
	ProtectedSaclSecurityInformation   SecurityInformation = 0x40000000
	UnprotectedDaclSecurityInformation SecurityInformation = 0x20000000
	UnprotectedSaclSecurityInformation SecurityInformation = 0x10000000
)

type SecurityDescriptor struct {
	windows.SECURITY_DESCRIPTOR
}

type ObjectType uint32

const (
	ObjectTypeUnknownObject         ObjectType = 0
	ObjectTypeFileObject            ObjectType = 1
	ObjectTypeService               ObjectType = 2
	ObjectTypePrinter               ObjectType = 3
	ObjectTypeRegistryKey           ObjectType = 4
	ObjectTypeLMShare               ObjectType = 5
	ObjectTypeKernelObject          ObjectType = 6
	ObjectTypeWindowObject          ObjectType = 7
	ObjectTypeDSObject              ObjectType = 8
	ObjectTypeDSObjectAll           ObjectType = 9
	ObjectTypeProviderDefinedObject ObjectType = 10
	ObjectTypeWMIGUIDObject         ObjectType = 11
	ObjectTypeRegistryWow6432Key    ObjectType = 12
	ObjectTypeRegistryWow6464Key    ObjectType = 13
)

func WindowsGetNamedSecurityInfo(name string, objType ObjectType, securityInfo SecurityInformation) (*SecurityDescriptor, error) {

	sd, err := windows.GetNamedSecurityInfo(
		name,
		windows.SE_OBJECT_TYPE(objType),
		windows.SECURITY_INFORMATION(securityInfo),
	)

	return &SecurityDescriptor{SECURITY_DESCRIPTOR: *sd}, err
}
