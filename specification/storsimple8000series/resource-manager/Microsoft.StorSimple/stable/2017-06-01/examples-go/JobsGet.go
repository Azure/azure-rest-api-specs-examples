package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/JobsGet.json
func ExampleJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().Get(ctx, "sca07forsdktest", "70a29339-de6d-48e8-b24f-e25ee6769a00", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armstorsimple8000series.Job{
	// 	Name: to.Ptr("70a29339-de6d-48e8-b24f-e25ee6769a00"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/jobs"),
	// 	ID: to.Ptr("/subscriptions/d3ebfe71-b7a9-4c57-92b9-68a2afde4de5/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/sca07forsdktest/jobs/70a29339-de6d-48e8-b24f-e25ee6769a00"),
	// 	Kind: to.Ptr("Series8000"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T13:08:19.818Z"); return t}()),
	// 	PercentComplete: to.Ptr[int32](100),
	// 	Properties: &armstorsimple8000series.JobProperties{
	// 		DeviceID: to.Ptr("/subscriptions/d3ebfe71-b7a9-4c57-92b9-68a2afde4de5/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/sca07forsdktest"),
	// 		EntityLabel: to.Ptr("sca07forsdktest"),
	// 		EntityType: to.Ptr("Microsoft.StorSimple/managers/devices"),
	// 		IsCancellable: to.Ptr(false),
	// 		JobType: to.Ptr(armstorsimple8000series.JobTypeCreateCloudAppliance),
	// 	},
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T12:58:14.364Z"); return t}()),
	// 	Status: to.Ptr(armstorsimple8000series.JobStatusSucceeded),
	// }
}
