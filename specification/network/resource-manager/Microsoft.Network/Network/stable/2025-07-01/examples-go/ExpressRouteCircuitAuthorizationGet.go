package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteCircuitAuthorizationGet.json
func ExampleExpressRouteCircuitAuthorizationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteCircuitAuthorizationsClient().Get(ctx, "rg1", "circuitName", "authorizationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteCircuitAuthorizationsClientGetResponse{
	// 	ExpressRouteCircuitAuthorization: armnetwork.ExpressRouteCircuitAuthorization{
	// 		Name: to.Ptr("MyAuthorization1"),
	// 		Type: to.Ptr("Microsoft.Network/expressRouteCircuits/authorizations"),
	// 		Etag: to.Ptr("W/\"e33c875f-48df-4a91-b7d3-eb95b5ddbb89\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName/authorizations/MyAuthorization1"),
	// 		Properties: &armnetwork.AuthorizationPropertiesFormat{
	// 			AuthorizationKey: to.Ptr("authKey"),
	// 			AuthorizationUseStatus: to.Ptr(armnetwork.AuthorizationUseStatusAvailable),
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
