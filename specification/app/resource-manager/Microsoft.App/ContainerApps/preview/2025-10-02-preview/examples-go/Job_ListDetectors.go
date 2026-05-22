package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/Job_ListDetectors.json
func ExampleJobsClient_NewListDetectorsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("f07f3711-b45e-40fe-a941-4e6d93f851e6", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListDetectorsPager("mikono-workerapp-test-rg", "mikonojob1", nil)
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
		// page = armappcontainers.JobsClientListDetectorsResponse{
		// 	DiagnosticsCollection: armappcontainers.DiagnosticsCollection{
		// 		Value: []*armappcontainers.Diagnostics{
		// 			{
		// 				ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/jobs/mikonojob1/detectors/cappjobContainerAppAvailabilityMetrics"),
		// 				Name: to.Ptr("cappjobContainerAppAvailabilityMetrics"),
		// 				Type: to.Ptr("Microsoft.App/containerapps/detectors"),
		// 				Properties: &armappcontainers.DiagnosticsProperties{
		// 					Metadata: &armappcontainers.DiagnosticsDefinition{
		// 						ID: to.Ptr("cappjobContainerAppAvailabilityMetrics"),
		// 						Name: to.Ptr("Availability Metrics for Container App Jobs"),
		// 						Author: to.Ptr(""),
		// 						Category: to.Ptr("Availability and Performance"),
		// 						SupportTopicList: []*armappcontainers.DiagnosticSupportTopic{
		// 						},
		// 						Type: to.Ptr("Analysis"),
		// 						Score: to.Ptr[float32](0),
		// 					},
		// 					Dataset: []*armappcontainers.DiagnosticsDataAPIResponse{
		// 					},
		// 					Status: &armappcontainers.DiagnosticsStatus{
		// 						StatusID: to.Ptr[int32](4),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
