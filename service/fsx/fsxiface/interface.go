// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package fsxiface provides an interface to enable mocking the Amazon FSx service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package fsxiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/fsx"
)

// FSxAPI provides an interface to enable mocking the
// fsx.FSx service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon FSx.
//	func myFunc(svc fsxiface.FSxAPI) bool {
//	    // Make svc.AssociateFileSystemAliases request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := fsx.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockFSxClient struct {
//	    fsxiface.FSxAPI
//	}
//	func (m *mockFSxClient) AssociateFileSystemAliases(input *fsx.AssociateFileSystemAliasesInput) (*fsx.AssociateFileSystemAliasesOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockFSxClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type FSxAPI interface {
	AssociateFileSystemAliases(*fsx.AssociateFileSystemAliasesInput) (*fsx.AssociateFileSystemAliasesOutput, error)
	AssociateFileSystemAliasesWithContext(aws.Context, *fsx.AssociateFileSystemAliasesInput, ...request.Option) (*fsx.AssociateFileSystemAliasesOutput, error)
	AssociateFileSystemAliasesRequest(*fsx.AssociateFileSystemAliasesInput) (*request.Request, *fsx.AssociateFileSystemAliasesOutput)

	CancelDataRepositoryTask(*fsx.CancelDataRepositoryTaskInput) (*fsx.CancelDataRepositoryTaskOutput, error)
	CancelDataRepositoryTaskWithContext(aws.Context, *fsx.CancelDataRepositoryTaskInput, ...request.Option) (*fsx.CancelDataRepositoryTaskOutput, error)
	CancelDataRepositoryTaskRequest(*fsx.CancelDataRepositoryTaskInput) (*request.Request, *fsx.CancelDataRepositoryTaskOutput)

	CopyBackup(*fsx.CopyBackupInput) (*fsx.CopyBackupOutput, error)
	CopyBackupWithContext(aws.Context, *fsx.CopyBackupInput, ...request.Option) (*fsx.CopyBackupOutput, error)
	CopyBackupRequest(*fsx.CopyBackupInput) (*request.Request, *fsx.CopyBackupOutput)

	CreateBackup(*fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error)
	CreateBackupWithContext(aws.Context, *fsx.CreateBackupInput, ...request.Option) (*fsx.CreateBackupOutput, error)
	CreateBackupRequest(*fsx.CreateBackupInput) (*request.Request, *fsx.CreateBackupOutput)

	CreateDataRepositoryAssociation(*fsx.CreateDataRepositoryAssociationInput) (*fsx.CreateDataRepositoryAssociationOutput, error)
	CreateDataRepositoryAssociationWithContext(aws.Context, *fsx.CreateDataRepositoryAssociationInput, ...request.Option) (*fsx.CreateDataRepositoryAssociationOutput, error)
	CreateDataRepositoryAssociationRequest(*fsx.CreateDataRepositoryAssociationInput) (*request.Request, *fsx.CreateDataRepositoryAssociationOutput)

	CreateDataRepositoryTask(*fsx.CreateDataRepositoryTaskInput) (*fsx.CreateDataRepositoryTaskOutput, error)
	CreateDataRepositoryTaskWithContext(aws.Context, *fsx.CreateDataRepositoryTaskInput, ...request.Option) (*fsx.CreateDataRepositoryTaskOutput, error)
	CreateDataRepositoryTaskRequest(*fsx.CreateDataRepositoryTaskInput) (*request.Request, *fsx.CreateDataRepositoryTaskOutput)

	CreateFileCache(*fsx.CreateFileCacheInput) (*fsx.CreateFileCacheOutput, error)
	CreateFileCacheWithContext(aws.Context, *fsx.CreateFileCacheInput, ...request.Option) (*fsx.CreateFileCacheOutput, error)
	CreateFileCacheRequest(*fsx.CreateFileCacheInput) (*request.Request, *fsx.CreateFileCacheOutput)

	CreateFileSystem(*fsx.CreateFileSystemInput) (*fsx.CreateFileSystemOutput, error)
	CreateFileSystemWithContext(aws.Context, *fsx.CreateFileSystemInput, ...request.Option) (*fsx.CreateFileSystemOutput, error)
	CreateFileSystemRequest(*fsx.CreateFileSystemInput) (*request.Request, *fsx.CreateFileSystemOutput)

	CreateFileSystemFromBackup(*fsx.CreateFileSystemFromBackupInput) (*fsx.CreateFileSystemFromBackupOutput, error)
	CreateFileSystemFromBackupWithContext(aws.Context, *fsx.CreateFileSystemFromBackupInput, ...request.Option) (*fsx.CreateFileSystemFromBackupOutput, error)
	CreateFileSystemFromBackupRequest(*fsx.CreateFileSystemFromBackupInput) (*request.Request, *fsx.CreateFileSystemFromBackupOutput)

	CreateSnapshot(*fsx.CreateSnapshotInput) (*fsx.CreateSnapshotOutput, error)
	CreateSnapshotWithContext(aws.Context, *fsx.CreateSnapshotInput, ...request.Option) (*fsx.CreateSnapshotOutput, error)
	CreateSnapshotRequest(*fsx.CreateSnapshotInput) (*request.Request, *fsx.CreateSnapshotOutput)

	CreateStorageVirtualMachine(*fsx.CreateStorageVirtualMachineInput) (*fsx.CreateStorageVirtualMachineOutput, error)
	CreateStorageVirtualMachineWithContext(aws.Context, *fsx.CreateStorageVirtualMachineInput, ...request.Option) (*fsx.CreateStorageVirtualMachineOutput, error)
	CreateStorageVirtualMachineRequest(*fsx.CreateStorageVirtualMachineInput) (*request.Request, *fsx.CreateStorageVirtualMachineOutput)

	CreateVolume(*fsx.CreateVolumeInput) (*fsx.CreateVolumeOutput, error)
	CreateVolumeWithContext(aws.Context, *fsx.CreateVolumeInput, ...request.Option) (*fsx.CreateVolumeOutput, error)
	CreateVolumeRequest(*fsx.CreateVolumeInput) (*request.Request, *fsx.CreateVolumeOutput)

	CreateVolumeFromBackup(*fsx.CreateVolumeFromBackupInput) (*fsx.CreateVolumeFromBackupOutput, error)
	CreateVolumeFromBackupWithContext(aws.Context, *fsx.CreateVolumeFromBackupInput, ...request.Option) (*fsx.CreateVolumeFromBackupOutput, error)
	CreateVolumeFromBackupRequest(*fsx.CreateVolumeFromBackupInput) (*request.Request, *fsx.CreateVolumeFromBackupOutput)

	DeleteBackup(*fsx.DeleteBackupInput) (*fsx.DeleteBackupOutput, error)
	DeleteBackupWithContext(aws.Context, *fsx.DeleteBackupInput, ...request.Option) (*fsx.DeleteBackupOutput, error)
	DeleteBackupRequest(*fsx.DeleteBackupInput) (*request.Request, *fsx.DeleteBackupOutput)

	DeleteDataRepositoryAssociation(*fsx.DeleteDataRepositoryAssociationInput) (*fsx.DeleteDataRepositoryAssociationOutput, error)
	DeleteDataRepositoryAssociationWithContext(aws.Context, *fsx.DeleteDataRepositoryAssociationInput, ...request.Option) (*fsx.DeleteDataRepositoryAssociationOutput, error)
	DeleteDataRepositoryAssociationRequest(*fsx.DeleteDataRepositoryAssociationInput) (*request.Request, *fsx.DeleteDataRepositoryAssociationOutput)

	DeleteFileCache(*fsx.DeleteFileCacheInput) (*fsx.DeleteFileCacheOutput, error)
	DeleteFileCacheWithContext(aws.Context, *fsx.DeleteFileCacheInput, ...request.Option) (*fsx.DeleteFileCacheOutput, error)
	DeleteFileCacheRequest(*fsx.DeleteFileCacheInput) (*request.Request, *fsx.DeleteFileCacheOutput)

	DeleteFileSystem(*fsx.DeleteFileSystemInput) (*fsx.DeleteFileSystemOutput, error)
	DeleteFileSystemWithContext(aws.Context, *fsx.DeleteFileSystemInput, ...request.Option) (*fsx.DeleteFileSystemOutput, error)
	DeleteFileSystemRequest(*fsx.DeleteFileSystemInput) (*request.Request, *fsx.DeleteFileSystemOutput)

	DeleteSnapshot(*fsx.DeleteSnapshotInput) (*fsx.DeleteSnapshotOutput, error)
	DeleteSnapshotWithContext(aws.Context, *fsx.DeleteSnapshotInput, ...request.Option) (*fsx.DeleteSnapshotOutput, error)
	DeleteSnapshotRequest(*fsx.DeleteSnapshotInput) (*request.Request, *fsx.DeleteSnapshotOutput)

	DeleteStorageVirtualMachine(*fsx.DeleteStorageVirtualMachineInput) (*fsx.DeleteStorageVirtualMachineOutput, error)
	DeleteStorageVirtualMachineWithContext(aws.Context, *fsx.DeleteStorageVirtualMachineInput, ...request.Option) (*fsx.DeleteStorageVirtualMachineOutput, error)
	DeleteStorageVirtualMachineRequest(*fsx.DeleteStorageVirtualMachineInput) (*request.Request, *fsx.DeleteStorageVirtualMachineOutput)

	DeleteVolume(*fsx.DeleteVolumeInput) (*fsx.DeleteVolumeOutput, error)
	DeleteVolumeWithContext(aws.Context, *fsx.DeleteVolumeInput, ...request.Option) (*fsx.DeleteVolumeOutput, error)
	DeleteVolumeRequest(*fsx.DeleteVolumeInput) (*request.Request, *fsx.DeleteVolumeOutput)

	DescribeBackups(*fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error)
	DescribeBackupsWithContext(aws.Context, *fsx.DescribeBackupsInput, ...request.Option) (*fsx.DescribeBackupsOutput, error)
	DescribeBackupsRequest(*fsx.DescribeBackupsInput) (*request.Request, *fsx.DescribeBackupsOutput)

	DescribeBackupsPages(*fsx.DescribeBackupsInput, func(*fsx.DescribeBackupsOutput, bool) bool) error
	DescribeBackupsPagesWithContext(aws.Context, *fsx.DescribeBackupsInput, func(*fsx.DescribeBackupsOutput, bool) bool, ...request.Option) error

	DescribeDataRepositoryAssociations(*fsx.DescribeDataRepositoryAssociationsInput) (*fsx.DescribeDataRepositoryAssociationsOutput, error)
	DescribeDataRepositoryAssociationsWithContext(aws.Context, *fsx.DescribeDataRepositoryAssociationsInput, ...request.Option) (*fsx.DescribeDataRepositoryAssociationsOutput, error)
	DescribeDataRepositoryAssociationsRequest(*fsx.DescribeDataRepositoryAssociationsInput) (*request.Request, *fsx.DescribeDataRepositoryAssociationsOutput)

	DescribeDataRepositoryAssociationsPages(*fsx.DescribeDataRepositoryAssociationsInput, func(*fsx.DescribeDataRepositoryAssociationsOutput, bool) bool) error
	DescribeDataRepositoryAssociationsPagesWithContext(aws.Context, *fsx.DescribeDataRepositoryAssociationsInput, func(*fsx.DescribeDataRepositoryAssociationsOutput, bool) bool, ...request.Option) error

	DescribeDataRepositoryTasks(*fsx.DescribeDataRepositoryTasksInput) (*fsx.DescribeDataRepositoryTasksOutput, error)
	DescribeDataRepositoryTasksWithContext(aws.Context, *fsx.DescribeDataRepositoryTasksInput, ...request.Option) (*fsx.DescribeDataRepositoryTasksOutput, error)
	DescribeDataRepositoryTasksRequest(*fsx.DescribeDataRepositoryTasksInput) (*request.Request, *fsx.DescribeDataRepositoryTasksOutput)

	DescribeDataRepositoryTasksPages(*fsx.DescribeDataRepositoryTasksInput, func(*fsx.DescribeDataRepositoryTasksOutput, bool) bool) error
	DescribeDataRepositoryTasksPagesWithContext(aws.Context, *fsx.DescribeDataRepositoryTasksInput, func(*fsx.DescribeDataRepositoryTasksOutput, bool) bool, ...request.Option) error

	DescribeFileCaches(*fsx.DescribeFileCachesInput) (*fsx.DescribeFileCachesOutput, error)
	DescribeFileCachesWithContext(aws.Context, *fsx.DescribeFileCachesInput, ...request.Option) (*fsx.DescribeFileCachesOutput, error)
	DescribeFileCachesRequest(*fsx.DescribeFileCachesInput) (*request.Request, *fsx.DescribeFileCachesOutput)

	DescribeFileCachesPages(*fsx.DescribeFileCachesInput, func(*fsx.DescribeFileCachesOutput, bool) bool) error
	DescribeFileCachesPagesWithContext(aws.Context, *fsx.DescribeFileCachesInput, func(*fsx.DescribeFileCachesOutput, bool) bool, ...request.Option) error

	DescribeFileSystemAliases(*fsx.DescribeFileSystemAliasesInput) (*fsx.DescribeFileSystemAliasesOutput, error)
	DescribeFileSystemAliasesWithContext(aws.Context, *fsx.DescribeFileSystemAliasesInput, ...request.Option) (*fsx.DescribeFileSystemAliasesOutput, error)
	DescribeFileSystemAliasesRequest(*fsx.DescribeFileSystemAliasesInput) (*request.Request, *fsx.DescribeFileSystemAliasesOutput)

	DescribeFileSystemAliasesPages(*fsx.DescribeFileSystemAliasesInput, func(*fsx.DescribeFileSystemAliasesOutput, bool) bool) error
	DescribeFileSystemAliasesPagesWithContext(aws.Context, *fsx.DescribeFileSystemAliasesInput, func(*fsx.DescribeFileSystemAliasesOutput, bool) bool, ...request.Option) error

	DescribeFileSystems(*fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error)
	DescribeFileSystemsWithContext(aws.Context, *fsx.DescribeFileSystemsInput, ...request.Option) (*fsx.DescribeFileSystemsOutput, error)
	DescribeFileSystemsRequest(*fsx.DescribeFileSystemsInput) (*request.Request, *fsx.DescribeFileSystemsOutput)

	DescribeFileSystemsPages(*fsx.DescribeFileSystemsInput, func(*fsx.DescribeFileSystemsOutput, bool) bool) error
	DescribeFileSystemsPagesWithContext(aws.Context, *fsx.DescribeFileSystemsInput, func(*fsx.DescribeFileSystemsOutput, bool) bool, ...request.Option) error

	DescribeSnapshots(*fsx.DescribeSnapshotsInput) (*fsx.DescribeSnapshotsOutput, error)
	DescribeSnapshotsWithContext(aws.Context, *fsx.DescribeSnapshotsInput, ...request.Option) (*fsx.DescribeSnapshotsOutput, error)
	DescribeSnapshotsRequest(*fsx.DescribeSnapshotsInput) (*request.Request, *fsx.DescribeSnapshotsOutput)

	DescribeSnapshotsPages(*fsx.DescribeSnapshotsInput, func(*fsx.DescribeSnapshotsOutput, bool) bool) error
	DescribeSnapshotsPagesWithContext(aws.Context, *fsx.DescribeSnapshotsInput, func(*fsx.DescribeSnapshotsOutput, bool) bool, ...request.Option) error

	DescribeStorageVirtualMachines(*fsx.DescribeStorageVirtualMachinesInput) (*fsx.DescribeStorageVirtualMachinesOutput, error)
	DescribeStorageVirtualMachinesWithContext(aws.Context, *fsx.DescribeStorageVirtualMachinesInput, ...request.Option) (*fsx.DescribeStorageVirtualMachinesOutput, error)
	DescribeStorageVirtualMachinesRequest(*fsx.DescribeStorageVirtualMachinesInput) (*request.Request, *fsx.DescribeStorageVirtualMachinesOutput)

	DescribeStorageVirtualMachinesPages(*fsx.DescribeStorageVirtualMachinesInput, func(*fsx.DescribeStorageVirtualMachinesOutput, bool) bool) error
	DescribeStorageVirtualMachinesPagesWithContext(aws.Context, *fsx.DescribeStorageVirtualMachinesInput, func(*fsx.DescribeStorageVirtualMachinesOutput, bool) bool, ...request.Option) error

	DescribeVolumes(*fsx.DescribeVolumesInput) (*fsx.DescribeVolumesOutput, error)
	DescribeVolumesWithContext(aws.Context, *fsx.DescribeVolumesInput, ...request.Option) (*fsx.DescribeVolumesOutput, error)
	DescribeVolumesRequest(*fsx.DescribeVolumesInput) (*request.Request, *fsx.DescribeVolumesOutput)

	DescribeVolumesPages(*fsx.DescribeVolumesInput, func(*fsx.DescribeVolumesOutput, bool) bool) error
	DescribeVolumesPagesWithContext(aws.Context, *fsx.DescribeVolumesInput, func(*fsx.DescribeVolumesOutput, bool) bool, ...request.Option) error

	DisassociateFileSystemAliases(*fsx.DisassociateFileSystemAliasesInput) (*fsx.DisassociateFileSystemAliasesOutput, error)
	DisassociateFileSystemAliasesWithContext(aws.Context, *fsx.DisassociateFileSystemAliasesInput, ...request.Option) (*fsx.DisassociateFileSystemAliasesOutput, error)
	DisassociateFileSystemAliasesRequest(*fsx.DisassociateFileSystemAliasesInput) (*request.Request, *fsx.DisassociateFileSystemAliasesOutput)

	ListTagsForResource(*fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *fsx.ListTagsForResourceInput, ...request.Option) (*fsx.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*fsx.ListTagsForResourceInput) (*request.Request, *fsx.ListTagsForResourceOutput)

	ListTagsForResourcePages(*fsx.ListTagsForResourceInput, func(*fsx.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *fsx.ListTagsForResourceInput, func(*fsx.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ReleaseFileSystemNfsV3Locks(*fsx.ReleaseFileSystemNfsV3LocksInput) (*fsx.ReleaseFileSystemNfsV3LocksOutput, error)
	ReleaseFileSystemNfsV3LocksWithContext(aws.Context, *fsx.ReleaseFileSystemNfsV3LocksInput, ...request.Option) (*fsx.ReleaseFileSystemNfsV3LocksOutput, error)
	ReleaseFileSystemNfsV3LocksRequest(*fsx.ReleaseFileSystemNfsV3LocksInput) (*request.Request, *fsx.ReleaseFileSystemNfsV3LocksOutput)

	RestoreVolumeFromSnapshot(*fsx.RestoreVolumeFromSnapshotInput) (*fsx.RestoreVolumeFromSnapshotOutput, error)
	RestoreVolumeFromSnapshotWithContext(aws.Context, *fsx.RestoreVolumeFromSnapshotInput, ...request.Option) (*fsx.RestoreVolumeFromSnapshotOutput, error)
	RestoreVolumeFromSnapshotRequest(*fsx.RestoreVolumeFromSnapshotInput) (*request.Request, *fsx.RestoreVolumeFromSnapshotOutput)

	StartMisconfiguredStateRecovery(*fsx.StartMisconfiguredStateRecoveryInput) (*fsx.StartMisconfiguredStateRecoveryOutput, error)
	StartMisconfiguredStateRecoveryWithContext(aws.Context, *fsx.StartMisconfiguredStateRecoveryInput, ...request.Option) (*fsx.StartMisconfiguredStateRecoveryOutput, error)
	StartMisconfiguredStateRecoveryRequest(*fsx.StartMisconfiguredStateRecoveryInput) (*request.Request, *fsx.StartMisconfiguredStateRecoveryOutput)

	TagResource(*fsx.TagResourceInput) (*fsx.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *fsx.TagResourceInput, ...request.Option) (*fsx.TagResourceOutput, error)
	TagResourceRequest(*fsx.TagResourceInput) (*request.Request, *fsx.TagResourceOutput)

	UntagResource(*fsx.UntagResourceInput) (*fsx.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *fsx.UntagResourceInput, ...request.Option) (*fsx.UntagResourceOutput, error)
	UntagResourceRequest(*fsx.UntagResourceInput) (*request.Request, *fsx.UntagResourceOutput)

	UpdateDataRepositoryAssociation(*fsx.UpdateDataRepositoryAssociationInput) (*fsx.UpdateDataRepositoryAssociationOutput, error)
	UpdateDataRepositoryAssociationWithContext(aws.Context, *fsx.UpdateDataRepositoryAssociationInput, ...request.Option) (*fsx.UpdateDataRepositoryAssociationOutput, error)
	UpdateDataRepositoryAssociationRequest(*fsx.UpdateDataRepositoryAssociationInput) (*request.Request, *fsx.UpdateDataRepositoryAssociationOutput)

	UpdateFileCache(*fsx.UpdateFileCacheInput) (*fsx.UpdateFileCacheOutput, error)
	UpdateFileCacheWithContext(aws.Context, *fsx.UpdateFileCacheInput, ...request.Option) (*fsx.UpdateFileCacheOutput, error)
	UpdateFileCacheRequest(*fsx.UpdateFileCacheInput) (*request.Request, *fsx.UpdateFileCacheOutput)

	UpdateFileSystem(*fsx.UpdateFileSystemInput) (*fsx.UpdateFileSystemOutput, error)
	UpdateFileSystemWithContext(aws.Context, *fsx.UpdateFileSystemInput, ...request.Option) (*fsx.UpdateFileSystemOutput, error)
	UpdateFileSystemRequest(*fsx.UpdateFileSystemInput) (*request.Request, *fsx.UpdateFileSystemOutput)

	UpdateSnapshot(*fsx.UpdateSnapshotInput) (*fsx.UpdateSnapshotOutput, error)
	UpdateSnapshotWithContext(aws.Context, *fsx.UpdateSnapshotInput, ...request.Option) (*fsx.UpdateSnapshotOutput, error)
	UpdateSnapshotRequest(*fsx.UpdateSnapshotInput) (*request.Request, *fsx.UpdateSnapshotOutput)

	UpdateStorageVirtualMachine(*fsx.UpdateStorageVirtualMachineInput) (*fsx.UpdateStorageVirtualMachineOutput, error)
	UpdateStorageVirtualMachineWithContext(aws.Context, *fsx.UpdateStorageVirtualMachineInput, ...request.Option) (*fsx.UpdateStorageVirtualMachineOutput, error)
	UpdateStorageVirtualMachineRequest(*fsx.UpdateStorageVirtualMachineInput) (*request.Request, *fsx.UpdateStorageVirtualMachineOutput)

	UpdateVolume(*fsx.UpdateVolumeInput) (*fsx.UpdateVolumeOutput, error)
	UpdateVolumeWithContext(aws.Context, *fsx.UpdateVolumeInput, ...request.Option) (*fsx.UpdateVolumeOutput, error)
	UpdateVolumeRequest(*fsx.UpdateVolumeInput) (*request.Request, *fsx.UpdateVolumeOutput)
}

var _ FSxAPI = (*fsx.FSx)(nil)
