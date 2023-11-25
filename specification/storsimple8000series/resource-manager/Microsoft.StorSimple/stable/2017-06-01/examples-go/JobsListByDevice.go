package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/JobsListByDevice.json
func ExampleJobsClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListByDevicePager("Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", &armstorsimple8000series.JobsClientListByDeviceOptions{Filter: to.Ptr("jobType%20eq%20'ManualBackup'")})
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
		// 			Name: to.Ptr("880e1774-94a8-4f3e-85e6-a61e6b94a8b7"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/jobs/880e1774-94a8-4f3e-85e6-a61e6b94a8b7"),
		// 			Kind: to.Ptr("Series8000"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T22:02:39.005Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple8000series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
		// 				DataStats: &armstorsimple8000series.DataStatistics{
		// 					CloudData: to.Ptr[int64](138),
		// 					ProcessedData: to.Ptr[int64](21474836480),
		// 					Throughput: to.Ptr[int64](76),
		// 					TotalData: to.Ptr[int64](21474836480),
		// 				},
		// 				DeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 				EntityLabel: to.Ptr("BkUpPolicy01ForSDKTest"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies"),
		// 				IsCancellable: to.Ptr(true),
		// 				JobType: to.Ptr(armstorsimple8000series.JobTypeManualBackup),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T22:01:19.708Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("69829a17-ab98-40c3-a6d3-fe8d315f52af"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/jobs/69829a17-ab98-40c3-a6d3-fe8d315f52af"),
		// 			Kind: to.Ptr("Series8000"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T21:30:37.397Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple8000series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
		// 				DataStats: &armstorsimple8000series.DataStatistics{
		// 					CloudData: to.Ptr[int64](152),
		// 					ProcessedData: to.Ptr[int64](26843545600),
		// 					Throughput: to.Ptr[int64](656),
		// 					TotalData: to.Ptr[int64](26843545600),
		// 				},
		// 				DeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 				EntityLabel: to.Ptr("BkUpPolicy01ForSDKTest"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies"),
		// 				IsCancellable: to.Ptr(true),
		// 				JobType: to.Ptr(armstorsimple8000series.JobTypeManualBackup),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T21:29:15.914Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 		},
		// 		{
		// 			Name: to.Ptr("0826eda8-3d17-4cb9-b2af-d18ecf6ab819"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/jobs/0826eda8-3d17-4cb9-b2af-d18ecf6ab819"),
		// 			Kind: to.Ptr("Series8000"),
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T21:23:21.028Z"); return t}()),
		// 			PercentComplete: to.Ptr[int32](100),
		// 			Properties: &armstorsimple8000series.JobProperties{
		// 				BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
		// 				BackupType: to.Ptr(armstorsimple8000series.BackupTypeCloudSnapshot),
		// 				DataStats: &armstorsimple8000series.DataStatistics{
		// 					CloudData: to.Ptr[int64](152),
		// 					ProcessedData: to.Ptr[int64](26843545600),
		// 					Throughput: to.Ptr[int64](85),
		// 					TotalData: to.Ptr[int64](26843545600),
		// 				},
		// 				DeviceID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest"),
		// 				EntityLabel: to.Ptr("BkUpPolicyForSDKTest1032280949"),
		// 				EntityType: to.Ptr("Microsoft.StorSimple/managers/devices/backupPolicies"),
		// 				IsCancellable: to.Ptr(true),
		// 				JobType: to.Ptr(armstorsimple8000series.JobTypeManualBackup),
		// 			},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-23T21:22:01.372Z"); return t}()),
		// 			Status: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
		// 	}},
		// }
	}
}
