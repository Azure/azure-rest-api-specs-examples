package armdashboard_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard/v2"
)

// Generated from example definition: 2024-11-01-preview/ManagedPrivateEndpoints_List.json
func ExampleManagedPrivateEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdashboard.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedPrivateEndpointsClient().NewListPager("myResourceGroup", "myWorkspace", nil)
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
		// page = armdashboard.ManagedPrivateEndpointsClientListResponse{
		// 	ManagedPrivateEndpointModelListResponse: armdashboard.ManagedPrivateEndpointModelListResponse{
		// 		Value: []*armdashboard.ManagedPrivateEndpointModel{
		// 			{
		// 				Name: to.Ptr("myMPEName"),
		// 				Type: to.Ptr("Microsoft.Dashboard/grafana/managedPrivateEndpoint"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/Microsoft.Dashboard/grafana/myWorkspace/managedPrivateEndpoints/myPrivateEndpointName"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armdashboard.ManagedPrivateEndpointModelProperties{
		// 					ConnectionState: &armdashboard.ManagedPrivateEndpointConnectionState{
		// 						Description: to.Ptr("Auto-Approved"),
		// 						Status: to.Ptr(armdashboard.ManagedPrivateEndpointConnectionStatusApproved),
		// 					},
		// 					GroupIDs: []*string{
		// 						to.Ptr("grafana"),
		// 					},
		// 					PrivateLinkResourceID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-000000000000/resourceGroups/xx-rg/providers/Microsoft.Kusto/Clusters/sampleKustoResource1"),
		// 					PrivateLinkResourceRegion: to.Ptr("West US"),
		// 					PrivateLinkServicePrivateIP: to.Ptr("10.0.0.5"),
		// 					PrivateLinkServiceURL: to.Ptr("my-self-hosted-influxdb.westus.mydomain.com"),
		// 					ProvisioningState: to.Ptr(armdashboard.ProvisioningStateSucceeded),
		// 					RequestMessage: to.Ptr("Example Request Message"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myMPEName2"),
		// 				Type: to.Ptr("Microsoft.Dashboard/grafana/managedPrivateEndpoint"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/Microsoft.Dashboard/grafana/myWorkspace/managedPrivateEndpoints/myPrivateEndpointName2"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armdashboard.ManagedPrivateEndpointModelProperties{
		// 					ConnectionState: &armdashboard.ManagedPrivateEndpointConnectionState{
		// 						Description: to.Ptr("Example Reject Reason"),
		// 						Status: to.Ptr(armdashboard.ManagedPrivateEndpointConnectionStatusRejected),
		// 					},
		// 					GroupIDs: []*string{
		// 						to.Ptr("grafana"),
		// 					},
		// 					PrivateLinkResourceID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-000000000000/resourceGroups/xx-rg/providers/Microsoft.Kusto/Clusters/sampleKustoResource2"),
		// 					PrivateLinkResourceRegion: to.Ptr("West US"),
		// 					ProvisioningState: to.Ptr(armdashboard.ProvisioningStateSucceeded),
		// 					RequestMessage: to.Ptr("Example Request Message 2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
