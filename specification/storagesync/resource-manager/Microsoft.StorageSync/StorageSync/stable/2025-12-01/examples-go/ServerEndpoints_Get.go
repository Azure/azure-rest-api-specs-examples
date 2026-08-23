package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync/v2"
)

// Generated from example definition: 2025-12-01/ServerEndpoints_Get.json
func ExampleServerEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("52b8da2f-61e0-4a1f-8dde-336911f367fb", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerEndpointsClient().Get(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleServerEndpoint_1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragesync.ServerEndpointsClientGetResponse{
	// 	XMSCorrelationRequestID: to.Ptr("d166ca76-dad2-49df-b409-d2acfd42d730"),
	// 	XMSRequestID: to.Ptr("74e55a4d-1c6f-46de-9a8d-278e53a47403"),
	// 	ServerEndpoint: armstoragesync.ServerEndpoint{
	// 		Name: to.Ptr("SampleServerEndpoint_1"),
	// 		Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/serverEndpoints"),
	// 		ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1/serverEndpoints/SampleServerEndpoint_1"),
	// 		Properties: &armstoragesync.ServerEndpointProperties{
	// 			CloudTiering: to.Ptr(armstoragesync.FeatureStatusOn),
	// 			CloudTieringStatus: &armstoragesync.ServerEndpointCloudTieringStatus{
	// 				CachePerformance: &armstoragesync.CloudTieringCachePerformance{
	// 					CacheHitBytes: to.Ptr[int64](922337203685477600),
	// 					CacheHitBytesPercent: to.Ptr[int32](50),
	// 					CacheMissBytes: to.Ptr[int64](922337203685477600),
	// 					LastUpdatedTimestamp: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 				},
	// 				DatePolicyStatus: &armstoragesync.CloudTieringDatePolicyStatus{
	// 					LastUpdatedTimestamp: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 					TieredFilesMostRecentAccessTimestamp: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 				},
	// 				FilesNotTiering: &armstoragesync.CloudTieringFilesNotTiering{
	// 					Errors: []*armstoragesync.FilesNotTieringError{
	// 						{
	// 							ErrorCode: to.Ptr[int32](-2134347771),
	// 							FileCount: to.Ptr[int64](10),
	// 						},
	// 						{
	// 							ErrorCode: to.Ptr[int32](-2134347770),
	// 							FileCount: to.Ptr[int64](20),
	// 						},
	// 						{
	// 							ErrorCode: to.Ptr[int32](-2134347769),
	// 							FileCount: to.Ptr[int64](30),
	// 						},
	// 					},
	// 					LastUpdatedTimestamp: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 					TotalFileCount: to.Ptr[int64](60),
	// 				},
	// 				Health: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 				HealthLastUpdatedTimestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 32, 51, 105791500, time.UTC)),
	// 				LastCloudTieringResult: to.Ptr[int32](-2134347771),
	// 				LastSuccessTimestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 32, 51, 105791500, time.UTC)),
	// 				LastUpdatedTimestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 32, 51, 105791500, time.UTC)),
	// 				LowDiskMode: &armstoragesync.CloudTieringLowDiskMode{
	// 					LastUpdatedTimestamp: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 					State: to.Ptr(armstoragesync.CloudTieringLowDiskModeStateDisabled),
	// 				},
	// 				SpaceSavings: &armstoragesync.CloudTieringSpaceSavings{
	// 					CachedSizeBytes: to.Ptr[int64](50000000),
	// 					LastUpdatedTimestamp: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 					SpaceSavingsBytes: to.Ptr[int64](50000000),
	// 					SpaceSavingsPercent: to.Ptr[int32](50),
	// 					TotalSizeCloudBytes: to.Ptr[int64](100000000),
	// 					VolumeSizeBytes: to.Ptr[int64](922337203685477600),
	// 				},
	// 				VolumeFreeSpacePolicyStatus: &armstoragesync.CloudTieringVolumeFreeSpacePolicyStatus{
	// 					CurrentVolumeFreeSpacePercent: to.Ptr[int32](5),
	// 					EffectiveVolumeFreeSpacePolicy: to.Ptr[int32](30),
	// 					LastUpdatedTimestamp: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 				},
	// 			},
	// 			FriendlyName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
	// 			InitialDownloadPolicy: to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
	// 			InitialUploadPolicy: to.Ptr(armstoragesync.InitialUploadPolicyMerge),
	// 			LastOperationName: to.Ptr("ICreateServerEndpointWorkflow"),
	// 			LastWorkflowID: to.Ptr("storageSyncServices/healthDemo1/workflows/569afb5c-7172-4cf8-a8ea-d987f727f11a"),
	// 			LocalCacheMode: to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
	// 			OfflineDataTransfer: to.Ptr(armstoragesync.FeatureStatusOn),
	// 			OfflineDataTransferShareName: to.Ptr("myfileshare"),
	// 			OfflineDataTransferStorageAccountResourceID: to.Ptr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/myRG/providers/Microsoft.Storage/storageAccounts/mysa"),
	// 			OfflineDataTransferStorageAccountTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			RecallStatus: &armstoragesync.ServerEndpointRecallStatus{
	// 				LastUpdatedTimestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 32, 51, 105791500, time.UTC)),
	// 				RecallErrors: []*armstoragesync.ServerEndpointRecallError{
	// 					{
	// 						Count: to.Ptr[int64](5),
	// 						ErrorCode: to.Ptr[int32](-2134347775),
	// 					},
	// 					{
	// 						Count: to.Ptr[int64](10),
	// 						ErrorCode: to.Ptr[int32](-2134347774),
	// 					},
	// 					{
	// 						Count: to.Ptr[int64](15),
	// 						ErrorCode: to.Ptr[int32](-2134347773),
	// 					},
	// 				},
	// 				TotalRecallErrorsCount: to.Ptr[int64](30),
	// 			},
	// 			ServerEndpointProvisioningStatus: &armstoragesync.ServerEndpointProvisioningStatus{
	// 				ProvisioningStatus: to.Ptr(armstoragesync.ServerProvisioningStatusReadySyncFunctional),
	// 				ProvisioningStepStatuses: []*armstoragesync.ServerEndpointProvisioningStepStatus{
	// 					{
	// 						Name: to.Ptr("ManifestGeneration"),
	// 						AdditionalInformation: map[string]*string{
	// 							"ItemsProcessed": to.Ptr("1001"),
	// 							"ItemsRemaining": to.Ptr("0"),
	// 						},
	// 						EndTime: to.Ptr(time.Date(2019, time.April, 17, 19, 6, 59, 195922700, time.UTC)),
	// 						ErrorCode: to.Ptr[int32](0),
	// 						MinutesLeft: to.Ptr[int32](0),
	// 						ProgressPercentage: to.Ptr[int32](100),
	// 						StartTime: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 						Status: to.Ptr("Completed"),
	// 					},
	// 					{
	// 						Name: to.Ptr("ManifestConsumption"),
	// 						AdditionalInformation: map[string]*string{
	// 							"ItemsProcessed": to.Ptr("1001"),
	// 							"ItemsRemaining": to.Ptr("0"),
	// 						},
	// 						EndTime: to.Ptr(time.Date(2019, time.April, 17, 19, 6, 59, 195922700, time.UTC)),
	// 						ErrorCode: to.Ptr[int32](0),
	// 						MinutesLeft: to.Ptr[int32](0),
	// 						ProgressPercentage: to.Ptr[int32](100),
	// 						StartTime: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 						Status: to.Ptr("Completed"),
	// 					},
	// 					{
	// 						Name: to.Ptr("NamespaceApply"),
	// 						AdditionalInformation: map[string]*string{
	// 							"ItemsProcessed": to.Ptr("1001"),
	// 							"ItemsRemaining": to.Ptr("0"),
	// 						},
	// 						EndTime: to.Ptr(time.Date(2019, time.April, 17, 19, 6, 59, 195922700, time.UTC)),
	// 						ErrorCode: to.Ptr[int32](0),
	// 						MinutesLeft: to.Ptr[int32](0),
	// 						ProgressPercentage: to.Ptr[int32](100),
	// 						StartTime: to.Ptr(time.Date(2019, time.April, 17, 19, 4, 59, 195922700, time.UTC)),
	// 						Status: to.Ptr("Completed"),
	// 					},
	// 					{
	// 						Name: to.Ptr("PrepareForSync"),
	// 						AdditionalInformation: map[string]*string{
	// 							"ItemsProcessed": to.Ptr("1001"),
	// 							"ItemsRemaining": to.Ptr("0"),
	// 						},
	// 						EndTime: to.Ptr(time.Date(2019, time.May, 17, 19, 8, 59, 195922700, time.UTC)),
	// 						ErrorCode: to.Ptr[int32](0),
	// 						MinutesLeft: to.Ptr[int32](0),
	// 						ProgressPercentage: to.Ptr[int32](100),
	// 						StartTime: to.Ptr(time.Date(2019, time.May, 17, 19, 6, 59, 195922700, time.UTC)),
	// 						Status: to.Ptr("Completed"),
	// 					},
	// 				},
	// 				ProvisioningType: to.Ptr("FastDRv2"),
	// 			},
	// 			ServerLocalPath: to.Ptr("D:\\SampleServerEndpoint_1"),
	// 			ServerName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
	// 			ServerResourceID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
	// 			SyncStatus: &armstoragesync.ServerEndpointSyncStatus{
	// 				CombinedHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 				DownloadActivity: &armstoragesync.ServerEndpointSyncActivityStatus{
	// 					AppliedBytes: to.Ptr[int64](94805587),
	// 					AppliedItemCount: to.Ptr[int64](100),
	// 					PerItemErrorCount: to.Ptr[int64](0),
	// 					Timestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 32, 51, 105791500, time.UTC)),
	// 					TotalBytes: to.Ptr[int64](19583674),
	// 					TotalItemCount: to.Ptr[int64](305),
	// 					RemainingFileCount: to.Ptr[int64](200),
	// 					RemainingDirectoryCount: to.Ptr[int64](5),
	// 					RemainingDeleteCount: to.Ptr[int64](0),
	// 					RemainingLogicalSizeBytes: to.Ptr[int64](19583674),
	// 					IsRemainingFinal: to.Ptr(true),
	// 					RecentItemsPerSecond: to.Ptr[float64](12.5),
	// 					RecentMegabytesPerSecond: to.Ptr[float64](1.2),
	// 					Warning: to.Ptr(armstoragesync.ServerEndpointSyncSessionWarningTypeNoWarning),
	// 				},
	// 				DownloadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateHealthy),
	// 				DownloadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
	// 					LastSyncPerItemErrorCount: to.Ptr[int64](0),
	// 					LastSyncResult: to.Ptr[int32](0),
	// 					LastSyncSuccessTimestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 28, 33, 921733400, time.UTC)),
	// 					LastSyncTimestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 28, 33, 921733400, time.UTC)),
	// 				},
	// 				LastUpdatedTimestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 32, 51, 105791500, time.UTC)),
	// 				OfflineDataTransferStatus: to.Ptr(armstoragesync.ServerEndpointOfflineDataTransferStateInProgress),
	// 				SyncActivity: to.Ptr(armstoragesync.ServerEndpointSyncActivityStateUploadAndDownload),
	// 				UploadActivity: &armstoragesync.ServerEndpointSyncActivityStatus{
	// 					AppliedBytes: to.Ptr[int64](57348983),
	// 					AppliedItemCount: to.Ptr[int64](1000),
	// 					PerItemErrorCount: to.Ptr[int64](300),
	// 					Timestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 32, 51, 105791500, time.UTC)),
	// 					TotalBytes: to.Ptr[int64](1958367412),
	// 					TotalItemCount: to.Ptr[int64](2310),
	// 					RemainingFileCount: to.Ptr[int64](1290),
	// 					RemainingDirectoryCount: to.Ptr[int64](10),
	// 					RemainingDeleteCount: to.Ptr[int64](0),
	// 					RemainingLogicalSizeBytes: to.Ptr[int64](1901018429),
	// 					IsRemainingFinal: to.Ptr(true),
	// 					RecentItemsPerSecond: to.Ptr[float64](125.3),
	// 					RecentMegabytesPerSecond: to.Ptr[float64](5.8),
	// 					InProgressLargeFilePath: to.Ptr("dir/largefile.zip"),
	// 					InProgressLargeFileSizeBytes: to.Ptr[int64](1073741824),
	// 					InProgressLargeFilePercentComplete: to.Ptr[int32](42),
	// 					Warning: to.Ptr(armstoragesync.ServerEndpointSyncSessionWarningTypeNoWarning),
	// 				},
	// 				UploadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 				UploadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
	// 					LastSyncPerItemErrorCount: to.Ptr[int64](1000),
	// 					LastSyncResult: to.Ptr[int32](-2134351810),
	// 					LastSyncTimestamp: to.Ptr(time.Date(2018, time.June, 11, 23, 32, 51, 105791500, time.UTC)),
	// 				},
	// 			},
	// 			TierFilesOlderThanDays: to.Ptr[int32](0),
	// 			VolumeFreeSpacePercent: to.Ptr[int32](100),
	// 		},
	// 	},
	// }
}
