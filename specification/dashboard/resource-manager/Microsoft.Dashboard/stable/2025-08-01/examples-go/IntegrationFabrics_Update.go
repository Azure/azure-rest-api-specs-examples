package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2025-08-01/IntegrationFabrics_Update.json
func ExampleIntegrationFabricsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationFabricsClient().BeginUpdate(ctx, "myResourceGroup", "myWorkspace", "sampleIntegration", armdashboard.IntegrationFabricUpdateParameters{
		Properties: &armdashboard.IntegrationFabricPropertiesUpdateParameters{
			Scenarios: []*string{
				to.Ptr("scenario1"),
			},
		},
		Tags: map[string]*string{
			"Environment": to.Ptr("Dev 2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdashboard.IntegrationFabricsClientUpdateResponse{
	// 	IntegrationFabric: &armdashboard.IntegrationFabric{
	// 		Name: to.Ptr("sampleIntegration"),
	// 		Type: to.Ptr("Microsoft.Dashboard/grafana/integrationFabrics"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Dashboard/grafana/myWorkspace/integrationFabrics/myIntegrationFabricName"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armdashboard.IntegrationFabricProperties{
	// 			DataSourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAmw"),
	// 			ProvisioningState: to.Ptr(armdashboard.ProvisioningStateSucceeded),
	// 			Scenarios: []*string{
	// 				to.Ptr("scenario1"),
	// 			},
	// 			TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myAks"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"Environment": to.Ptr("Dev 2"),
	// 		},
	// 	},
	// }
}
