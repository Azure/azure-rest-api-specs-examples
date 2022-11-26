package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_CloudServiceConfiguration.json
func ExamplePoolClient_Create_createPoolFullCloudServiceConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			ApplicationLicenses: []*string{
				to.Ptr("app-license0"),
				to.Ptr("app-license1")},
			ApplicationPackages: []*armbatch.ApplicationPackageReference{
				{
					ID:      to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
					Version: to.Ptr("asdf"),
				}},
			Certificates: []*armbatch.CertificateReference{
				{
					ID:            to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567"),
					StoreLocation: to.Ptr(armbatch.CertificateStoreLocationLocalMachine),
					StoreName:     to.Ptr("MY"),
					Visibility: []*armbatch.CertificateVisibility{
						to.Ptr(armbatch.CertificateVisibilityRemoteUser)},
				}},
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				CloudServiceConfiguration: &armbatch.CloudServiceConfiguration{
					OSFamily:  to.Ptr("4"),
					OSVersion: to.Ptr("WA-GUEST-OS-4.45_201708-01"),
				},
			},
			DisplayName:            to.Ptr("my-pool-name"),
			InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateEnabled),
			Metadata: []*armbatch.MetadataItem{
				{
					Name:  to.Ptr("metadata-1"),
					Value: to.Ptr("value-1"),
				},
				{
					Name:  to.Ptr("metadata-2"),
					Value: to.Ptr("value-2"),
				}},
			NetworkConfiguration: &armbatch.NetworkConfiguration{
				PublicIPAddressConfiguration: &armbatch.PublicIPAddressConfiguration{
					IPAddressIDs: []*string{
						to.Ptr("/subscriptions/subid1/resourceGroups/rg13/providers/Microsoft.Network/publicIPAddresses/ip135"),
						to.Ptr("/subscriptions/subid2/resourceGroups/rg24/providers/Microsoft.Network/publicIPAddresses/ip268")},
					Provision: to.Ptr(armbatch.IPAddressProvisioningTypeUserManaged),
				},
				SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
			},
			ScaleSettings: &armbatch.ScaleSettings{
				FixedScale: &armbatch.FixedScaleSettings{
					NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionTaskCompletion),
					ResizeTimeout:          to.Ptr("PT8M"),
					TargetDedicatedNodes:   to.Ptr[int32](6),
					TargetLowPriorityNodes: to.Ptr[int32](28),
				},
			},
			StartTask: &armbatch.StartTask{
				CommandLine: to.Ptr("cmd /c SET"),
				EnvironmentSettings: []*armbatch.EnvironmentSetting{
					{
						Name:  to.Ptr("MYSET"),
						Value: to.Ptr("1234"),
					}},
				MaxTaskRetryCount: to.Ptr[int32](6),
				ResourceFiles: []*armbatch.ResourceFile{
					{
						FileMode: to.Ptr("777"),
						FilePath: to.Ptr("c:\\temp\\gohere"),
						HTTPURL:  to.Ptr("https://testaccount.blob.core.windows.net/example-blob-file"),
					}},
				UserIdentity: &armbatch.UserIdentity{
					AutoUser: &armbatch.AutoUserSpecification{
						ElevationLevel: to.Ptr(armbatch.ElevationLevelAdmin),
						Scope:          to.Ptr(armbatch.AutoUserScopePool),
					},
				},
				WaitForSuccess: to.Ptr(true),
			},
			TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
				NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypePack),
			},
			TaskSlotsPerNode: to.Ptr[int32](13),
			UserAccounts: []*armbatch.UserAccount{
				{
					Name:           to.Ptr("username1"),
					ElevationLevel: to.Ptr(armbatch.ElevationLevelAdmin),
					LinuxUserConfiguration: &armbatch.LinuxUserConfiguration{
						Gid:           to.Ptr[int32](4567),
						SSHPrivateKey: to.Ptr("sshprivatekeyvalue"),
						UID:           to.Ptr[int32](1234),
					},
					Password: to.Ptr("<ExamplePassword>"),
				}},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
