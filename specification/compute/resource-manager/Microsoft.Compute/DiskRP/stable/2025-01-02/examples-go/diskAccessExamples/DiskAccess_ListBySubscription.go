package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskAccessExamples/DiskAccess_ListBySubscription.json
func ExampleDiskAccessesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiskAccessesClient().NewListPager(nil)
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
		// page.DiskAccessList = armcompute.DiskAccessList{
		// 	Value: []*armcompute.DiskAccess{
		// 		{
		// 			Name: to.Ptr("myDiskAccess"),
		// 			Type: to.Ptr("Microsoft.Compute/diskAccesses"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("PrivateEndpoints"),
		// 			},
		// 			Properties: &armcompute.DiskAccessProperties{
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T04:41:35.079Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myDiskAccess2"),
		// 			Type: to.Ptr("Microsoft.Compute/diskAccesses"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/mySecondResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("PrivateEndpoints"),
		// 			},
		// 			Properties: &armcompute.DiskAccessProperties{
		// 				PrivateEndpointConnections: []*armcompute.PrivateEndpointConnection{
		// 					{
		// 						Name: to.Ptr("myDiskAccess.d4914cfa-6bc2-4049-a57c-3d1f622d8eef"),
		// 						Type: to.Ptr("Microsoft.Compute/diskAccesses/PrivateEndpointConnections"),
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/mySecondResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess2/privateEndpoinConnections/myDiskAccess2.d4914cfa-6bc2-4049-a57c-3d1f622d8eef"),
		// 						Properties: &armcompute.PrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armcompute.PrivateEndpoint{
		// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/mySecondResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint2"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armcompute.PrivateLinkServiceConnectionState{
		// 								Description: to.Ptr("Auto-Approved"),
		// 								ActionsRequired: to.Ptr("None"),
		// 								Status: to.Ptr(armcompute.PrivateEndpointServiceConnectionStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armcompute.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T04:41:35.079Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
