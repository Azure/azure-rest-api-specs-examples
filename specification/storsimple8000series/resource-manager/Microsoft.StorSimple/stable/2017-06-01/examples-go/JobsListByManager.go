package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/JobsListByManager.json
func ExampleJobsClient_NewListByManagerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListByManagerPager("ResourceGroupForSDKTest", "ManagerForSDKTest1", &armstorsimple8000series.JobsClientListByManagerOptions{Filter: to.Ptr("jobType%20eq%20'FailoverVolumeContainers'")})
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
		// page.JobList = armstorsimple8000series.JobList{
		// 	Value: []*armstorsimple8000series.Job{
		// 		{
		// 			Name: to.Ptr("07103ea6-0092-4bee-853c-72a98256421e"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/jemdeviceforsdk/jobs/07103ea6-0092-4bee-853c-72a98256421e"),
		// 			Kind: to.Ptr("Series8000"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-22T09:19:06.885Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple8000series.JobProperties{
		// 				DataStats: &armstorsimple8000series.DataStatistics{
		// 					CloudData: to.Ptr[int64](0),
		// 					ProcessedData: to.Ptr[int64](0),
		// 					Throughput: to.Ptr[int64](0),
		// 					TotalData: to.Ptr[int64](0),
		// 				},
		// 				DeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/jemdeviceforsdk"),
		// 				EntityLabel: to.Ptr("Device05ForSDKTest"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple8000series.JobStage{
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepCreateSAC"),
		// 						Message: to.Ptr("Creation of storage account credentials"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepTakeOwnership"),
		// 						Message: to.Ptr("Transfer of volume containers and backups"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepRestoreVolumes"),
		// 						Message: to.Ptr("Restoration of volumes"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepCreateACR"),
		// 						Message: to.Ptr("Creation of ACRs"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepCreateSchedules"),
		// 						Message: to.Ptr("Creation of backup schedules"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 				}},
		// 				JobType: to.Ptr(armstorsimple8000series.JobTypeFailoverVolumeContainers),
		// 				SourceDeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-22T09:17:48.254Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("55d9f42a-b130-4bbe-8f40-985174a37ce8"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/jobs/55d9f42a-b130-4bbe-8f40-985174a37ce8"),
		// 			Kind: to.Ptr("Series8000"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-19T09:04:37.689Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple8000series.JobProperties{
		// 				DataStats: &armstorsimple8000series.DataStatistics{
		// 					CloudData: to.Ptr[int64](0),
		// 					ProcessedData: to.Ptr[int64](0),
		// 					Throughput: to.Ptr[int64](0),
		// 					TotalData: to.Ptr[int64](0),
		// 				},
		// 				DeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 				EntityLabel: to.Ptr("jemdeviceforsdk"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple8000series.JobStage{
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepCreateSAC"),
		// 						Message: to.Ptr("Creation of storage account credentials"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepTakeOwnership"),
		// 						Message: to.Ptr("Transfer of volume containers and backups"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepRestoreVolumes"),
		// 						Message: to.Ptr("Restoration of volumes"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepCreateACR"),
		// 						Message: to.Ptr("Creation of ACRs"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepCreateSchedules"),
		// 						Message: to.Ptr("Creation of backup schedules"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 				}},
		// 				JobType: to.Ptr(armstorsimple8000series.JobTypeFailoverVolumeContainers),
		// 				SourceDeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/jemdeviceforsdk"),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-19T09:02:57.435Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("a8418e26-99e8-4b11-883f-c08ca74db2b0"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/jemdeviceforsdk/jobs/a8418e26-99e8-4b11-883f-c08ca74db2b0"),
		// 			Kind: to.Ptr("Series8000"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-13T10:51:59.006Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple8000series.JobProperties{
		// 				DataStats: &armstorsimple8000series.DataStatistics{
		// 					CloudData: to.Ptr[int64](0),
		// 					ProcessedData: to.Ptr[int64](0),
		// 					Throughput: to.Ptr[int64](0),
		// 					TotalData: to.Ptr[int64](0),
		// 				},
		// 				DeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/jemdeviceforsdk"),
		// 				EntityLabel: to.Ptr("Device05ForSDKTest"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple8000series.JobStage{
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepCreateSAC"),
		// 						Message: to.Ptr("Creation of storage account credentials"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepTakeOwnership"),
		// 						Message: to.Ptr("Transfer of volume containers and backups"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepRestoreVolumes"),
		// 						Message: to.Ptr("Restoration of volumes"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepCreateACR"),
		// 						Message: to.Ptr("Creation of ACRs"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 					},
		// 					{
		// 						Detail: to.Ptr(""),
		// 						ErrorCode: to.Ptr("CiSDRJobStepCreateSchedules"),
		// 						Message: to.Ptr("Creation of backup schedules"),
		// 						StageStatus: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 				}},
		// 				JobType: to.Ptr(armstorsimple8000series.JobTypeFailoverVolumeContainers),
		// 				SourceDeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-13T10:50:46.764Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 	}},
		// }
	}
}
