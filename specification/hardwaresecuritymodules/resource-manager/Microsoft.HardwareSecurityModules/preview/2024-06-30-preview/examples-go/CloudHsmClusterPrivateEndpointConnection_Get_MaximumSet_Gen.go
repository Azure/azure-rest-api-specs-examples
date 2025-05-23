package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e838027e88cca634c1545e744630de9262a6e72a/specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/CloudHsmClusterPrivateEndpointConnection_Get_MaximumSet_Gen.json
func ExampleCloudHsmClusterPrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhardwaresecuritymodules.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudHsmClusterPrivateEndpointConnectionsClient().Get(ctx, "rgcloudhsm", "chsm1", "sample-pec", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armhardwaresecuritymodules.PrivateEndpointConnection{
	// 	Name: to.Ptr("sample-pec"),
	// 	Type: to.Ptr("Microsoft.HardwareSecurityModules/cloudHsmClusters/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgcloudhsm/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/chsm1/privateEndpointConnections/sample-pec"),
	// 	SystemData: &armhardwaresecuritymodules.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("User1"),
	// 		CreatedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:00:00.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User2"),
	// 		LastModifiedByType: to.Ptr(armhardwaresecuritymodules.CreatedByTypeUser),
	// 	},
	// 	Properties: &armhardwaresecuritymodules.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armhardwaresecuritymodules.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pec"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armhardwaresecuritymodules.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armhardwaresecuritymodules.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armhardwaresecuritymodules.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}
