package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/ManagedEnvironmentDiagnostics_List.json
func ExampleManagedEnvironmentDiagnosticsClient_ListDetectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("f07f3711-b45e-40fe-a941-4e6d93f851e6", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedEnvironmentDiagnosticsClient().ListDetectors(ctx, "mikono-workerapp-test-rg", "mikonokubeenv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.ManagedEnvironmentDiagnosticsClientListDetectorsResponse{
	// 	DiagnosticsCollection: armappcontainers.DiagnosticsCollection{
	// 		Value: []*armappcontainers.Diagnostics{
	// 			{
	// 				ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/managedEnvironments/mikonokubeenv/detectors/ManagedEnvAvailabilityMetrics"),
	// 				Name: to.Ptr("ManagedEnvAvailabilityMetrics"),
	// 				Type: to.Ptr("Microsoft.App/managedEnvironments/detectors"),
	// 				Properties: &armappcontainers.DiagnosticsProperties{
	// 					Metadata: &armappcontainers.DiagnosticsDefinition{
	// 						ID: to.Ptr("ManagedEnvAvailabilityMetrics"),
	// 						Name: to.Ptr("Availability Metrics for Managed Environments"),
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
