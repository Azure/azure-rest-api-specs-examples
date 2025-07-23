package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2024-11-01-preview/IntegrationFabrics_Get.json
func ExampleIntegrationFabricsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationFabricsClient().Get(ctx, "myResourceGroup", "myWorkspace", "sampleIntegration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdashboard.IntegrationFabricsClientGetResponse{
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
	// 				to.Ptr("scenario2"),
	// 			},
	// 			TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myAks"),
	// 		},
	// 	},
	// }
}
