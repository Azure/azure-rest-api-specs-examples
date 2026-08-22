package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync/v2"
)

// Generated from example definition: 2025-12-01/CloudEndpoints_Update.json
func ExampleCloudEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("11071075-D90D-4F53-B814-AF8F9B5C39D2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudEndpointsClient().BeginUpdate(ctx, "rgstoragesync", "llg", "wwuoouzucgvfrsvjfgsobajg", "mjpalurfyrwkmqeygi", armstoragesync.CloudEndpointUpdateParameters{
		Properties: &armstoragesync.CloudEndpointUpdateProperties{
			ChangeEnumerationIntervalDays: to.Ptr[int32](14),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragesync.CloudEndpointsClientUpdateResponse{
	// 	CloudEndpoint: armstoragesync.CloudEndpoint{
	// 		Name: to.Ptr("mjpalurfyrwkmqeygi"),
	// 		Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/cloudEndpoints"),
	// 		ID: to.Ptr("/subscriptions/11071075-D90D-4F53-B814-AF8F9B5C39D2/resourceGroups/rgstoragesync/providers/Microsoft.StorageSync/storageSyncServices/llg/syncGroups/wwuoouzucgvfrsvjfgsobajg/cloudEndpoints/mjpalurfyrwkmqeygi"),
	// 		Properties: &armstoragesync.CloudEndpointProperties{
	// 			AzureFileShareName: to.Ptr("fileshare1"),
	// 			ChangeEnumerationStatus: &armstoragesync.CloudEndpointChangeEnumerationStatus{
	// 				Activity: &armstoragesync.CloudEndpointChangeEnumerationActivity{
	// 					DeletesProgressPercent: to.Ptr[int32](62),
	// 					LastUpdatedTimestamp: to.Ptr(time.Date(2025, time.December, 3, 20, 59, 48, 891000000, time.UTC)),
	// 					MinutesRemaining: to.Ptr[int32](0),
	// 					OperationState: to.Ptr(armstoragesync.CloudEndpointChangeEnumerationActivityStateInitialEnumerationInProgress),
	// 					ProcessedDirectoriesCount: to.Ptr[int64](0),
	// 					ProcessedFilesCount: to.Ptr[int64](0),
	// 					ProgressPercent: to.Ptr[int32](30),
	// 					StartedTimestamp: to.Ptr(time.Date(2025, time.December, 3, 20, 59, 48, 892000000, time.UTC)),
	// 					StatusCode: to.Ptr[int32](13),
	// 					TotalCountsState: to.Ptr(armstoragesync.CloudEndpointChangeEnumerationTotalCountsStateCalculating),
	// 					TotalDirectoriesCount: to.Ptr[int64](0),
	// 					TotalFilesCount: to.Ptr[int64](0),
	// 					TotalSizeBytes: to.Ptr[int64](0),
	// 				},
	// 				LastEnumerationStatus: &armstoragesync.CloudEndpointLastChangeEnumerationStatus{
	// 					CompletedTimestamp: to.Ptr(time.Date(2025, time.December, 3, 20, 59, 48, 891000000, time.UTC)),
	// 					NamespaceDirectoriesCount: to.Ptr[int64](0),
	// 					NamespaceFilesCount: to.Ptr[int64](0),
	// 					NamespaceSizeBytes: to.Ptr[int64](0),
	// 					NextRunTimestamp: to.Ptr(time.Date(2025, time.December, 3, 20, 59, 48, 891000000, time.UTC)),
	// 					StartedTimestamp: to.Ptr(time.Date(2025, time.December, 3, 20, 59, 48, 891000000, time.UTC)),
	// 				},
	// 				LastUpdatedTimestamp: to.Ptr(time.Date(2025, time.December, 3, 20, 59, 48, 891000000, time.UTC)),
	// 			},
	// 			FriendlyName: to.Ptr("Cloud Endpoint"),
	// 			LastOperationName: to.Ptr("Initialize"),
	// 			LastWorkflowID: to.Ptr("080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
	// 			PartnershipID: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			StorageAccountResourceID: to.Ptr("/subscriptions/11071075-D90D-4F53-B814-AF8F9B5C39D2/resourceGroups/rgstoragesync/providers/Microsoft.Storage/storageAccounts/storageaccount1"),
	// 			StorageAccountTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			BackupEnabled: to.Ptr("true"),
	// 			ChangeEnumerationIntervalDays: to.Ptr[int32](14),
	// 		},
	// 		SystemData: &armstoragesync.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armstoragesync.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2025, time.December, 3, 20, 59, 33, 899000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armstoragesync.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2025, time.December, 3, 20, 59, 33, 899000000, time.UTC)),
	// 		},
	// 	},
	// }
}
