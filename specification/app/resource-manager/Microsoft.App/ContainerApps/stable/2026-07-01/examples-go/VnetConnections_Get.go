package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/VnetConnections_Get.json
func ExampleVnetConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVnetConnectionsClient().Get(ctx, "myRg", "testgroup", "myVnetConnection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.VnetConnectionsClientGetResponse{
	// 	VnetConnection: armappcontainers.VnetConnection{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRg/providers/Microsoft.App/sandboxGroups/testgroup/vnetConnections/myVnetConnection"),
	// 		Name: to.Ptr("myVnetConnection"),
	// 		Type: to.Ptr("Microsoft.App/sandboxGroups/vnetConnections"),
	// 		Properties: &armappcontainers.VnetConnectionProperties{
	// 			ProvisioningState: to.Ptr(armappcontainers.VnetConnectionProvisioningStateSucceeded),
	// 			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRg/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
	// 		},
	// 	},
	// }
}
