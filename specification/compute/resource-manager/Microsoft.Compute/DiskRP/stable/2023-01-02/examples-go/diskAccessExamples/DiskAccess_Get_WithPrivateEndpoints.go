package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-01-02/examples/diskAccessExamples/DiskAccess_Get_WithPrivateEndpoints.json
func ExampleDiskAccessesClient_Get_getInformationAboutADiskAccessResourceWithPrivateEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.DiskAccess = armcompute.DiskAccess{
	// 	Name: to.Ptr("myDiskAccess"),
	// 	Type: to.Ptr("Microsoft.Compute/diskAccesses"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("Development"),
	// 		"project": to.Ptr("PrivateEndpoints"),
	// 	},
	// 	Properties: &armcompute.DiskAccessProperties{
	// 		PrivateEndpointConnections: []*armcompute.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("myDiskAccess.d4914cfa-6bc2-4049-a57c-3d1f622d8eef"),
	// 				Type: to.Ptr("Microsoft.Compute/diskAccesses/PrivateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess/privateEndpoinConnections/myDiskAccess.d4914cfa-6bc2-4049-a57c-3d1f622d8eef"),
	// 				Properties: &armcompute.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armcompute.PrivateEndpoint{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armcompute.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("Auto-Approved"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr(armcompute.PrivateEndpointServiceConnectionStatusApproved),
	// 					},
	// 					ProvisioningState: to.Ptr(armcompute.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T04:41:35.079872+00:00"); return t}()),
	// 	},
	// }
}
