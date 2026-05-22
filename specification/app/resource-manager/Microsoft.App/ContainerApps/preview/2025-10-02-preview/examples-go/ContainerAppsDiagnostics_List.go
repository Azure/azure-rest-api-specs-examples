package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/ContainerAppsDiagnostics_List.json
func ExampleContainerAppsDiagnosticsClient_NewListDetectorsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("f07f3711-b45e-40fe-a941-4e6d93f851e6", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsDiagnosticsClient().NewListDetectorsPager("mikono-workerapp-test-rg", "mikono-capp-stage1", nil)
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
		// page = armappcontainers.ContainerAppsDiagnosticsClientListDetectorsResponse{
		// 	DiagnosticsCollection: armappcontainers.DiagnosticsCollection{
		// 		Value: []*armappcontainers.Diagnostics{
		// 			{
		// 				ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/containerApps/mikono-capp-stage1/detectors/cappContainerAppAvailabilityMetrics"),
		// 				Name: to.Ptr("cappContainerAppAvailabilityMetrics"),
		// 				Type: to.Ptr("Microsoft.App/containerapps/detectors"),
		// 				Properties: &armappcontainers.DiagnosticsProperties{
		// 					Metadata: &armappcontainers.DiagnosticsDefinition{
		// 						ID: to.Ptr("cappContainerAppAvailabilityMetrics"),
		// 						Name: to.Ptr("Availability Metrics for Container Apps"),
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
