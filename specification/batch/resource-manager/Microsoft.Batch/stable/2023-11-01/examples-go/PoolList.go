package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolList.json
func ExamplePoolClient_NewListByBatchAccountPager_listPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPoolClient().NewListByBatchAccountPager("default-azurebatch-japaneast", "sampleacct", &armbatch.PoolClientListByBatchAccountOptions{Maxresults: nil,
		Select: nil,
		Filter: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ListPoolsResult = armbatch.ListPoolsResult{
		// 	Value: []*armbatch.Pool{
		// 		{
		// 			Name: to.Ptr("testpool"),
		// 			Type: to.Ptr("Microsoft.Batch/batchAccounts/pools"),
		// 			Etag: to.Ptr("W/\"0x8D4EDFEBFADF4AB\""),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool"),
		// 			Properties: &armbatch.PoolProperties{
		// 				AllocationState: to.Ptr(armbatch.AllocationStateSteady),
		// 				AllocationStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
		// 				ApplicationLicenses: []*string{
		// 					to.Ptr("app-license0"),
		// 					to.Ptr("app-license1")},
		// 					ApplicationPackages: []*armbatch.ApplicationPackageReference{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
		// 							Version: to.Ptr("asdf"),
		// 					}},
		// 					Certificates: []*armbatch.CertificateReference{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567"),
		// 							StoreLocation: to.Ptr(armbatch.CertificateStoreLocationLocalMachine),
		// 							StoreName: to.Ptr("MY"),
		// 							Visibility: []*armbatch.CertificateVisibility{
		// 								to.Ptr(armbatch.CertificateVisibilityRemoteUser)},
		// 						}},
		// 						CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
		// 						CurrentDedicatedNodes: to.Ptr[int32](0),
		// 						CurrentLowPriorityNodes: to.Ptr[int32](0),
		// 						DeploymentConfiguration: &armbatch.DeploymentConfiguration{
		// 							CloudServiceConfiguration: &armbatch.CloudServiceConfiguration{
		// 								OSFamily: to.Ptr("4"),
		// 								OSVersion: to.Ptr("WA-GUEST-OS-4.45_201708-01"),
		// 							},
		// 						},
		// 						InterNodeCommunication: to.Ptr(armbatch.InterNodeCommunicationStateEnabled),
		// 						LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
		// 						Metadata: []*armbatch.MetadataItem{
		// 							{
		// 								Name: to.Ptr("metadata-1"),
		// 								Value: to.Ptr("value-1"),
		// 							},
		// 							{
		// 								Name: to.Ptr("metadata-2"),
		// 								Value: to.Ptr("value-2"),
		// 						}},
		// 						NetworkConfiguration: &armbatch.NetworkConfiguration{
		// 							EndpointConfiguration: &armbatch.PoolEndpointConfiguration{
		// 								InboundNatPools: []*armbatch.InboundNatPool{
		// 									{
		// 										Name: to.Ptr("testnat"),
		// 										BackendPort: to.Ptr[int32](12001),
		// 										FrontendPortRangeEnd: to.Ptr[int32](15100),
		// 										FrontendPortRangeStart: to.Ptr[int32](15000),
		// 										NetworkSecurityGroupRules: []*armbatch.NetworkSecurityGroupRule{
		// 											{
		// 												Access: to.Ptr(armbatch.NetworkSecurityGroupRuleAccessAllow),
		// 												Priority: to.Ptr[int32](150),
		// 												SourceAddressPrefix: to.Ptr("192.100.12.45"),
		// 												SourcePortRanges: []*string{
		// 													to.Ptr("*")},
		// 												},
		// 												{
		// 													Access: to.Ptr(armbatch.NetworkSecurityGroupRuleAccessDeny),
		// 													Priority: to.Ptr[int32](3500),
		// 													SourceAddressPrefix: to.Ptr("*"),
		// 													SourcePortRanges: []*string{
		// 														to.Ptr("*")},
		// 												}},
		// 												Protocol: to.Ptr(armbatch.InboundEndpointProtocolTCP),
		// 										}},
		// 									},
		// 									SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123"),
		// 								},
		// 								ProvisioningState: to.Ptr(armbatch.PoolProvisioningStateSucceeded),
		// 								ProvisioningStateTransitionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
		// 								ResizeOperationStatus: &armbatch.ResizeOperationStatus{
		// 									Errors: []*armbatch.ResizeError{
		// 										{
		// 											Code: to.Ptr("AllocationTimedout"),
		// 											Message: to.Ptr("Desired number of dedicated nodes could not be allocated as the resize timeout was reached"),
		// 									}},
		// 									NodeDeallocationOption: to.Ptr(armbatch.ComputeNodeDeallocationOptionTaskCompletion),
		// 									ResizeTimeout: to.Ptr("PT8M"),
		// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-28T10:22:55.940Z"); return t}()),
		// 									TargetDedicatedNodes: to.Ptr[int32](6),
		// 									TargetLowPriorityNodes: to.Ptr[int32](28),
		// 								},
		// 								ScaleSettings: &armbatch.ScaleSettings{
		// 									FixedScale: &armbatch.FixedScaleSettings{
		// 										ResizeTimeout: to.Ptr("PT8M"),
		// 										TargetDedicatedNodes: to.Ptr[int32](6),
		// 										TargetLowPriorityNodes: to.Ptr[int32](28),
		// 									},
		// 								},
		// 								StartTask: &armbatch.StartTask{
		// 									CommandLine: to.Ptr("cmd /c SET"),
		// 									EnvironmentSettings: []*armbatch.EnvironmentSetting{
		// 										{
		// 											Name: to.Ptr("MYSET"),
		// 											Value: to.Ptr("1234"),
		// 									}},
		// 									MaxTaskRetryCount: to.Ptr[int32](6),
		// 									ResourceFiles: []*armbatch.ResourceFile{
		// 										{
		// 											FileMode: to.Ptr("777"),
		// 											FilePath: to.Ptr("c:\\temp\\gohere"),
		// 											HTTPURL: to.Ptr("https://testaccount.blob.core.windows.net/example-blob-file"),
		// 									}},
		// 									UserIdentity: &armbatch.UserIdentity{
		// 										AutoUser: &armbatch.AutoUserSpecification{
		// 											ElevationLevel: to.Ptr(armbatch.ElevationLevelAdmin),
		// 											Scope: to.Ptr(armbatch.AutoUserScopePool),
		// 										},
		// 									},
		// 									WaitForSuccess: to.Ptr(true),
		// 								},
		// 								TaskSchedulingPolicy: &armbatch.TaskSchedulingPolicy{
		// 									NodeFillType: to.Ptr(armbatch.ComputeNodeFillTypePack),
		// 								},
		// 								TaskSlotsPerNode: to.Ptr[int32](13),
		// 								UserAccounts: []*armbatch.UserAccount{
		// 									{
		// 										Name: to.Ptr("username1"),
		// 										ElevationLevel: to.Ptr(armbatch.ElevationLevelAdmin),
		// 										LinuxUserConfiguration: &armbatch.LinuxUserConfiguration{
		// 											Gid: to.Ptr[int32](4567),
		// 											UID: to.Ptr[int32](1234),
		// 										},
		// 								}},
		// 								VMSize: to.Ptr("STANDARD_D4"),
		// 							},
		// 					}},
		// 				}
	}
}
