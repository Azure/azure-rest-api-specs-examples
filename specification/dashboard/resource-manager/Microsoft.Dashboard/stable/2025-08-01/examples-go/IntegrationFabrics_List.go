package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2025-08-01/IntegrationFabrics_List.json
func ExampleIntegrationFabricsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationFabricsClient().NewListPager("myResourceGroup", "myWorkspace", nil)
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
		// page = armdashboard.IntegrationFabricsClientListResponse{
		// 	IntegrationFabricListResponse: armdashboard.IntegrationFabricListResponse{
		// 		Value: []*armdashboard.IntegrationFabric{
		// 			{
		// 				Name: to.Ptr("sampleIntegration"),
		// 				Type: to.Ptr("Microsoft.Dashboard/grafana/integrationFabrics"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Dashboard/grafana/myWorkspace/integrationFabrics/myIntegrationFabricName"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armdashboard.IntegrationFabricProperties{
		// 					DataSourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAmw1"),
		// 					ProvisioningState: to.Ptr(armdashboard.ProvisioningStateSucceeded),
		// 					Scenarios: []*string{
		// 						to.Ptr("scenario1"),
		// 						to.Ptr("scenario2"),
		// 					},
		// 					TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myAks1"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sampleIntegration"),
		// 				Type: to.Ptr("Microsoft.Dashboard/grafana/integrationFabrics"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Dashboard/grafana/myWorkspace/integrationFabrics/myIntegrationFabricName"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armdashboard.IntegrationFabricProperties{
		// 					DataSourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAmw2"),
		// 					ProvisioningState: to.Ptr(armdashboard.ProvisioningStateSucceeded),
		// 					Scenarios: []*string{
		// 						to.Ptr("scenario1"),
		// 						to.Ptr("scenario2"),
		// 					},
		// 					TargetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myAks2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
