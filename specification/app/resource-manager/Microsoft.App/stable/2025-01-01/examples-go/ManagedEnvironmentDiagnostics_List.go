package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ManagedEnvironmentDiagnostics_List.json
func ExampleManagedEnvironmentDiagnosticsClient_ListDetectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.DiagnosticsCollection = armappcontainers.DiagnosticsCollection{
	// 	Value: []*armappcontainers.Diagnostics{
	// 		{
	// 			Name: to.Ptr("ManagedEnvAvailabilityMetrics"),
	// 			Type: to.Ptr("Microsoft.App/managedEnvironments/detectors"),
	// 			ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/managedEnvironments/mikonokubeenv/detectors/ManagedEnvAvailabilityMetrics"),
	// 			Properties: &armappcontainers.DiagnosticsProperties{
	// 				Dataset: []*armappcontainers.DiagnosticsDataAPIResponse{
	// 				},
	// 				Metadata: &armappcontainers.DiagnosticsDefinition{
	// 					Name: to.Ptr("Availability Metrics for Managed Environments"),
	// 					Type: to.Ptr("Analysis"),
	// 					Author: to.Ptr(""),
	// 					Category: to.Ptr("Availability and Performance"),
	// 					ID: to.Ptr("ManagedEnvAvailabilityMetrics"),
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
