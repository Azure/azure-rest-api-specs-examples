package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/ExpressRoutePortAuthorizationGet.json
func ExampleExpressRoutePortAuthorizationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRoutePortAuthorizationsClient().Get(ctx, "rg1", "expressRoutePortName", "authorizationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRoutePortAuthorization = armnetwork.ExpressRoutePortAuthorization{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ExpressRoutePorts/expressRoutePortName/authorizations/authorizationName"),
	// 	Name: to.Ptr("authorizationName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRoutePorts/authorizations"),
	// 	Etag: to.Ptr("W/\"e33c875f-48df-4a91-b7d3-eb95b5ddbb89\""),
	// 	Properties: &armnetwork.ExpressRoutePortAuthorizationPropertiesFormat{
	// 		AuthorizationKey: to.Ptr("authKey"),
	// 		AuthorizationUseStatus: to.Ptr(armnetwork.ExpressRoutePortAuthorizationUseStatusAvailable),
	// 		CircuitResourceURI: to.Ptr(""),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
