package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/AllVirtualNetworkGatewayRadiusServerSecretsList.json
func ExampleVirtualNetworkGatewaysClient_ListRadiusSecrets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkGatewaysClient().ListRadiusSecrets(ctx, "rg1", "vpngw", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RadiusAuthServerListResult = armnetwork.RadiusAuthServerListResult{
	// 	Value: []*armnetwork.RadiusAuthServer{
	// 		{
	// 			RadiusServerAddress: to.Ptr("1.1.1.1"),
	// 			RadiusServerSecret: to.Ptr("abc123"),
	// 		},
	// 		{
	// 			RadiusServerAddress: to.Ptr("2.2.2.2"),
	// 			RadiusServerSecret: to.Ptr("mnop1234"),
	// 	}},
	// }
}
