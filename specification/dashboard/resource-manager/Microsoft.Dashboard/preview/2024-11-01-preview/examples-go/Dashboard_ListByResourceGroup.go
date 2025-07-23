package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2024-11-01-preview/Dashboard_ListByResourceGroup.json
func ExampleManagedDashboardsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDashboardsClient().NewListPager("myResourceGroup", nil)
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
		// page = armdashboard.ManagedDashboardsClientListResponse{
		// 	ManagedDashboardListResponse: armdashboard.ManagedDashboardListResponse{
		// 		Value: []*armdashboard.ManagedDashboard{
		// 			{
		// 				Name: to.Ptr("myDashboard"),
		// 				Type: to.Ptr("Microsoft.Dashboard/dashboards"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Dashboard/dashboards/myDashboard"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armdashboard.ManagedDashboardProperties{
		// 					ProvisioningState: to.Ptr(armdashboard.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Environment": to.Ptr("Dev"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
