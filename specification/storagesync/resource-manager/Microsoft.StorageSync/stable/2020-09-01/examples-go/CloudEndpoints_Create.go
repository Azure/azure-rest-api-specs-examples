package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/CloudEndpoints_Create.json
func ExampleCloudEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudEndpointsClient().BeginCreate(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleCloudEndpoint_1", armstoragesync.CloudEndpointCreateParameters{
		Properties: &armstoragesync.CloudEndpointCreateParametersProperties{
			AzureFileShareName:       to.Ptr("cvcloud-afscv-0719-058-a94a1354-a1fd-4e9a-9a50-919fad8c4ba4"),
			FriendlyName:             to.Ptr("ankushbsubscriptionmgmtmab"),
			StorageAccountResourceID: to.Ptr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/tminienv59svc/providers/Microsoft.Storage/storageAccounts/tminienv59storage"),
			StorageAccountTenantID:   to.Ptr("\"72f988bf-86f1-41af-91ab-2d7cd011db47\""),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CloudEndpoint = armstoragesync.CloudEndpoint{
	// 	Name: to.Ptr("SampleCloudEndpoint_1"),
	// 	Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/cloudEndpoints"),
	// 	ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1/cloudEndpoints/SampleCloudEndpoint_1"),
	// 	Properties: &armstoragesync.CloudEndpointProperties{
	// 		AzureFileShareName: to.Ptr("cvcloud-afscv-0719-058-a94a1354-a1fd-4e9a-9a50-919fad8c4ba4"),
	// 		BackupEnabled: to.Ptr("false"),
	// 		FriendlyName: to.Ptr("ankushbsubscriptionmgmtmab"),
	// 		LastOperationName: to.Ptr("ICreateCloudEndpointWorkflow"),
	// 		LastWorkflowID: to.Ptr("storageSyncServices/GATest/workflows/24ba0c4a-348e-419b-8f7a-091d0d9f07a4"),
	// 		PartnershipID: to.Ptr("1|U0VSVkVSQVNTWU5DQ0xJRU5USEZTVjJ8RjhDODcwQTItMkFGNi00NDUyLTgzMDgtRjJCQTZEQjI3RkEwfEdFTkVSSUN8NTJCOERBMkYtNjFFMC00QTFGLThEREUtMzM2OTExRjM2N0ZC"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		StorageAccountResourceID: to.Ptr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/tminienv59svc/providers/Microsoft.Storage/storageAccounts/tminienv59storage"),
	// 		StorageAccountTenantID: to.Ptr("\"72f988bf-86f1-41af-91ab-2d7cd011db47\""),
	// 	},
	// }
}
