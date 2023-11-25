package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/JobsGet.json
func ExampleJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().Get(ctx, "HSDK-ARCSX4MVKZ", "06c7ee19-35a2-4248-bf1b-408009b31b63", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armstorsimple1200series.Job{
	// 	Name: to.Ptr("06c7ee19-35a2-4248-bf1b-408009b31b63"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ/jobs/06c7ee19-35a2-4248-bf1b-408009b31b63"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:14:39.969Z"); return t}()),
	// 	PercentComplete: to.Ptr[int32](100),
	// 	Properties: &armstorsimple1200series.JobProperties{
	// 		BackupPointInTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
	// 		DeviceID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
	// 		EntityID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-ARCSX4MVKZ"),
	// 		EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
	// 		IsCancellable: to.Ptr(false),
	// 		JobStages: []*armstorsimple1200series.JobStage{
	// 		},
	// 		JobType: to.Ptr(armstorsimple1200series.JobTypeBackup),
	// 		Stats: &armstorsimple1200series.JobStats{
	// 			CompletedWorkItemCount: to.Ptr[int32](0),
	// 			EstimatedTimeRemaining: to.Ptr[int32](0),
	// 			TotalWorkItemCount: to.Ptr[int32](0),
	// 		},
	// 		TargetType: to.Ptr(armstorsimple1200series.TargetTypeFileServer),
	// 	},
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T11:14:03.804Z"); return t}()),
	// 	Status: to.Ptr(armstorsimple1200series.JobStatusSucceeded),
	// }
}
