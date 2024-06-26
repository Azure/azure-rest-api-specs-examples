package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/JobsListByDevice.json
func ExampleJobsClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListByDevicePager("HSDK-ARCSX4MVKZ", "ResourceGroupForSDKTest", "hAzureSDKOperations", &armstorsimple1200series.JobsClientListByDeviceOptions{Filter: to.Ptr("jobType%20eq%20'Backup'")})
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
		// page.JobList = armstorsimple1200series.JobList{
		// 	Value: []*armstorsimple1200series.Job{
		// 		{
		// 			Name: to.Ptr("06c7ee19-35a2-4248-bf1b-408009b31b63"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/jobs/06c7ee19-35a2-4248-bf1b-408009b31b63"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:14:39.969Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple1200series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple1200series.JobStage{
		// 				},
		// 				JobType: to.Ptr(armstorsimple1200series.JobTypeBackup),
		// 				Stats: &armstorsimple1200series.JobStats{
		// 					CompletedWorkItemCount: to.Ptr[int32](0),
		// 					EstimatedTimeRemaining: to.Ptr[int32](0),
		// 					TotalWorkItemCount: to.Ptr[int32](0),
		// 				},
		// 				TargetType: to.Ptr(armstorsimple1200series.TargetTypeFileServer),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:14:03.804Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple1200series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("1d2de9da-a07f-4d73-a05f-01dd5a173128"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/jobs/1d2de9da-a07f-4d73-a05f-01dd5a173128"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T12:01:27.773Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple1200series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple1200series.JobStage{
		// 				},
		// 				JobType: to.Ptr(armstorsimple1200series.JobTypeBackup),
		// 				Stats: &armstorsimple1200series.JobStats{
		// 					CompletedWorkItemCount: to.Ptr[int32](0),
		// 					EstimatedTimeRemaining: to.Ptr[int32](0),
		// 					TotalWorkItemCount: to.Ptr[int32](0),
		// 				},
		// 				TargetType: to.Ptr(armstorsimple1200series.TargetTypeFileServer),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T12:00:48.719Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple1200series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("285ee145-913e-4885-bc01-6c904b1621be"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/jobs/285ee145-913e-4885-bc01-6c904b1621be"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T13:08:30.952Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple1200series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple1200series.JobStage{
		// 				},
		// 				JobType: to.Ptr(armstorsimple1200series.JobTypeBackup),
		// 				Stats: &armstorsimple1200series.JobStats{
		// 					CompletedWorkItemCount: to.Ptr[int32](0),
		// 					EstimatedTimeRemaining: to.Ptr[int32](0),
		// 					TotalWorkItemCount: to.Ptr[int32](0),
		// 				},
		// 				TargetType: to.Ptr(armstorsimple1200series.TargetTypeFileServer),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T13:07:51.054Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple1200series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("4886495a-9c0f-41ad-af0e-6a590b077be1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/jobs/4886495a-9c0f-41ad-af0e-6a590b077be1"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:51:26.333Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple1200series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple1200series.JobStage{
		// 				},
		// 				JobType: to.Ptr(armstorsimple1200series.JobTypeBackup),
		// 				Stats: &armstorsimple1200series.JobStats{
		// 					CompletedWorkItemCount: to.Ptr[int32](0),
		// 					EstimatedTimeRemaining: to.Ptr[int32](0),
		// 					TotalWorkItemCount: to.Ptr[int32](0),
		// 				},
		// 				TargetType: to.Ptr(armstorsimple1200series.TargetTypeFileServer),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:50:49.839Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple1200series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("57d1a3de-0174-47b6-8c8e-cbbb778316cd"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/jobs/57d1a3de-0174-47b6-8c8e-cbbb778316cd"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:53:26.667Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple1200series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple1200series.JobStage{
		// 				},
		// 				JobType: to.Ptr(armstorsimple1200series.JobTypeBackup),
		// 				Stats: &armstorsimple1200series.JobStats{
		// 					CompletedWorkItemCount: to.Ptr[int32](0),
		// 					EstimatedTimeRemaining: to.Ptr[int32](0),
		// 					TotalWorkItemCount: to.Ptr[int32](0),
		// 				},
		// 				TargetType: to.Ptr(armstorsimple1200series.TargetTypeFileServer),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:52:50.181Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple1200series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("6a645e88-9cf4-4e9b-8125-b5fdf71e8bee"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/jobs/6a645e88-9cf4-4e9b-8125-b5fdf71e8bee"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:48:07.972Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple1200series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple1200series.JobStage{
		// 				},
		// 				JobType: to.Ptr(armstorsimple1200series.JobTypeBackup),
		// 				Stats: &armstorsimple1200series.JobStats{
		// 					CompletedWorkItemCount: to.Ptr[int32](0),
		// 					EstimatedTimeRemaining: to.Ptr[int32](0),
		// 					TotalWorkItemCount: to.Ptr[int32](0),
		// 				},
		// 				TargetType: to.Ptr(armstorsimple1200series.TargetTypeFileServer),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:47:31.902Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple1200series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("70bddfae-689e-4ca8-9e32-8937fcf680e5"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/jobs/70bddfae-689e-4ca8-9e32-8937fcf680e5"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:45:52.808Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple1200series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple1200series.JobStage{
		// 				},
		// 				JobType: to.Ptr(armstorsimple1200series.JobTypeBackup),
		// 				Stats: &armstorsimple1200series.JobStats{
		// 					CompletedWorkItemCount: to.Ptr[int32](0),
		// 					EstimatedTimeRemaining: to.Ptr[int32](0),
		// 					TotalWorkItemCount: to.Ptr[int32](0),
		// 				},
		// 				TargetType: to.Ptr(armstorsimple1200series.TargetTypeFileServer),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:45:17.229Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple1200series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("c8053f2d-a5dc-4ecb-b2d0-be8f2db988e3"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/jobs/c8053f2d-a5dc-4ecb-b2d0-be8f2db988e3"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:07:57.731Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple1200series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
		// 				IsCancellable: to.Ptr(false),
		// 				JobStages: []*armstorsimple1200series.JobStage{
		// 				},
		// 				JobType: to.Ptr(armstorsimple1200series.JobTypeBackup),
		// 				Stats: &armstorsimple1200series.JobStats{
		// 					CompletedWorkItemCount: to.Ptr[int32](0),
		// 					EstimatedTimeRemaining: to.Ptr[int32](0),
		// 					TotalWorkItemCount: to.Ptr[int32](0),
		// 				},
		// 				TargetType: to.Ptr(armstorsimple1200series.TargetTypeFileServer),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:07:23.195Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple1200series.JobStatusSucceeded),
		// 	}},
		// }
	}
}
