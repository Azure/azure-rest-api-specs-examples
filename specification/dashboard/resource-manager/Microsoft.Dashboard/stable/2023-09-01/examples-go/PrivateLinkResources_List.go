package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/78eac0bd58633028293cb1ec1709baa200bed9e2/specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/PrivateLinkResources_List.json
func ExamplePrivateLinkResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListPager("myResourceGroup", "myWorkspace", nil)
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
		// page.PrivateLinkResourceListResult = armdashboard.PrivateLinkResourceListResult{
		// 	Value: []*armdashboard.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("grafana"),
		// 			Type: to.Ptr("Microsoft.Dashboard/grafana/PrivateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/Microsoft.Dashboard/grafana/myWorkspace/privateLinkResources/grafana"),
		// 			Properties: &armdashboard.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("grafana"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("grafana")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("grafana-test.azure.com")},
		// 					},
		// 			}},
		// 		}
	}
}
