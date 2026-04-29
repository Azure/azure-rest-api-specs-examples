package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-01-02/diskAccessExamples/DiskAccess_Get_WithPrivateEndpoints.json
func ExampleDiskAccessesClient_Get_getInformationAboutADiskAccessResourceWithPrivateEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiskAccessesClient().Get(ctx, "myResourceGroup", "myDiskAccess", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DiskAccessesClientGetResponse{
	// 	DiskAccess: &armcompute.DiskAccess{
	// 		Properties: &armcompute.DiskAccessProperties{
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T04:41:35.079872+00:00"); return t}()),
	// 			PrivateEndpointConnections: []*armcompute.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("myDiskAccess.d4914cfa-6bc2-4049-a57c-3d1f622d8eef"),
	// 					Type: to.Ptr("Microsoft.Compute/diskAccesses/PrivateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess/privateEndpoinConnections/myDiskAccess.d4914cfa-6bc2-4049-a57c-3d1f622d8eef"),
	// 					Properties: &armcompute.PrivateEndpointConnectionProperties{
	// 						ProvisioningState: to.Ptr(armcompute.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 						PrivateEndpoint: &armcompute.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armcompute.PrivateLinkServiceConnectionState{
	// 							ActionsRequired: to.Ptr("None"),
	// 							Description: to.Ptr("Auto-Approved"),
	// 							Status: to.Ptr(armcompute.PrivateEndpointServiceConnectionStatusApproved),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Type: to.Ptr("Microsoft.Compute/diskAccesses"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("Development"),
	// 			"project": to.Ptr("PrivateEndpoints"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess"),
	// 		Name: to.Ptr("myDiskAccess"),
	// 	},
	// }
}
