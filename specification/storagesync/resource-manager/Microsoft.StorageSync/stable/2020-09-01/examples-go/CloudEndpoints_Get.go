package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/CloudEndpoints_Get.json
func ExampleCloudEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudEndpointsClient().Get(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleCloudEndpoint_1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CloudEndpoint = armstoragesync.CloudEndpoint{
	// 	Name: to.Ptr("SampleCloudEndpoint_1"),
	// 	Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/cloudEndpoints"),
	// 	ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SyncGroup_Restore_08-08_Test112/cloudEndpoints/CEP_Restore_08-08_Test112"),
	// 	Properties: &armstoragesync.CloudEndpointProperties{
	// 		AzureFileShareName: to.Ptr(""),
	// 		ChangeEnumerationStatus: &armstoragesync.CloudEndpointChangeEnumerationStatus{
	// 			Activity: &armstoragesync.CloudEndpointChangeEnumerationActivity{
	// 				DeletesProgressPercent: to.Ptr[int32](12),
	// 				LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:14:59.096Z"); return t}()),
	// 				MinutesRemaining: to.Ptr[int32](589),
	// 				OperationState: to.Ptr(armstoragesync.CloudEndpointChangeEnumerationActivityStateEnumerationInProgress),
	// 				ProcessedDirectoriesCount: to.Ptr[int64](364),
	// 				ProcessedFilesCount: to.Ptr[int64](6948),
	// 				ProgressPercent: to.Ptr[int32](67),
	// 				StartedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:14:59.096Z"); return t}()),
	// 				StatusCode: to.Ptr[int32](0),
	// 				TotalCountsState: to.Ptr(armstoragesync.CloudEndpointChangeEnumerationTotalCountsStateFinal),
	// 				TotalDirectoriesCount: to.Ptr[int64](694),
	// 				TotalFilesCount: to.Ptr[int64](12834),
	// 				TotalSizeBytes: to.Ptr[int64](5782301239408),
	// 			},
	// 			LastEnumerationStatus: &armstoragesync.CloudEndpointLastChangeEnumerationStatus{
	// 				CompletedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:14:59.096Z"); return t}()),
	// 				NamespaceDirectoriesCount: to.Ptr[int64](28),
	// 				NamespaceFilesCount: to.Ptr[int64](3489),
	// 				NamespaceSizeBytes: to.Ptr[int64](3248804),
	// 				NextRunTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:14:59.096Z"); return t}()),
	// 				StartedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:14:59.096Z"); return t}()),
	// 			},
	// 			LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:14:59.096Z"); return t}()),
	// 		},
	// 		FriendlyName: to.Ptr("SampleAzureFileShareName_1"),
	// 		LastOperationName: to.Ptr("ICreateCloudEndpointWorkflow"),
	// 		LastWorkflowID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/workflows/a377fdd5-949a-40ab-9629-06cd0e9852f9"),
	// 		PartnershipID: to.Ptr("1|U0VSVkVSQVNTWU5DQ0xJRU5USEZTVjJ8MTkxNjYwQ0QtNkExQS00RjhDLTk3ODctQTZCRUQyMDZBMUREfEdFTkVSSUN8M0EwNDgyODMtMzM4Ri00MDAyLUE5REQtQTUwRkRBRENCMzky"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		StorageAccountResourceID: to.Ptr(""),
	// 		StorageAccountTenantID: to.Ptr("\"a4d1b191-c1af-4cef-a14b-f670e0beea52\""),
	// 	},
	// }
}
