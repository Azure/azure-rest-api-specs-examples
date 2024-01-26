package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/Cluster_ListStreamingJobs.json
func ExampleClustersClient_NewListStreamingJobsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListStreamingJobsPager("sjrg", "testcluster", nil)
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
		// page.ClusterJobListResult = armstreamanalytics.ClusterJobListResult{
		// 	Value: []*armstreamanalytics.ClusterJob{
		// 		{
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sjrg/providers/microsoft.streamAnalytics/streamingjobs/AFilterSample"),
		// 			JobState: to.Ptr(armstreamanalytics.JobStateRunning),
		// 			StreamingUnits: to.Ptr[int32](6),
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sjrg/providers/microsoft.streamAnalytics/streamingjobs/AnotherFilterSample"),
		// 			JobState: to.Ptr(armstreamanalytics.JobStateStopped),
		// 			StreamingUnits: to.Ptr[int32](1),
		// 	}},
		// }
	}
}
