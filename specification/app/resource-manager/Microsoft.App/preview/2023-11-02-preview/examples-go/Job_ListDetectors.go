package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d74afb775446d7f0bc1810fdc5a128c56289e854/specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Job_ListDetectors.json
func ExampleJobsClient_ListDetectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobsClient().ListDetectors(ctx, "mikono-workerapp-test-rg", "mikonojob1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticsCollection = armappcontainers.DiagnosticsCollection{
	// 	Value: []*armappcontainers.Diagnostics{
	// 		{
	// 			Name: to.Ptr("cappjobContainerAppAvailabilityMetrics"),
	// 			Type: to.Ptr("Microsoft.App/containerapps/detectors"),
	// 			ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/jobs/mikonojob1/detectors/cappjobContainerAppAvailabilityMetrics"),
	// 			Properties: &armappcontainers.DiagnosticsProperties{
	// 				Dataset: []*armappcontainers.DiagnosticsDataAPIResponse{
	// 				},
	// 				Metadata: &armappcontainers.DiagnosticsDefinition{
	// 					Name: to.Ptr("Availability Metrics for Container App Jobs"),
	// 					Type: to.Ptr("Analysis"),
	// 					Author: to.Ptr(""),
	// 					Category: to.Ptr("Availability and Performance"),
	// 					ID: to.Ptr("cappjobContainerAppAvailabilityMetrics"),
	// 					Score: to.Ptr[float32](0),
	// 					SupportTopicList: []*armappcontainers.DiagnosticSupportTopic{
	// 					},
	// 				},
	// 				Status: &armappcontainers.DiagnosticsStatus{
	// 					StatusID: to.Ptr[int32](4),
	// 				},
	// 			},
	// 	}},
	// }
}
