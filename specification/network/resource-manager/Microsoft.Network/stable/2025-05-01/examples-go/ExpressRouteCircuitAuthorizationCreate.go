package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ExpressRouteCircuitAuthorizationCreate.json
func ExampleExpressRouteCircuitAuthorizationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteCircuitAuthorizationsClient().BeginCreateOrUpdate(ctx, "rg1", "circuitName", "authorizatinName", armnetwork.ExpressRouteCircuitAuthorization{
		Properties: &armnetwork.AuthorizationPropertiesFormat{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteCircuitAuthorization = armnetwork.ExpressRouteCircuitAuthorization{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName/authorizations/authorizationName"),
	// 	Name: to.Ptr("authorizationName"),
	// 	Etag: to.Ptr("W/\"e22dd4b2-4c24-44cf-b702-70a472b62914\""),
	// 	Properties: &armnetwork.AuthorizationPropertiesFormat{
	// 		AuthorizationUseStatus: to.Ptr(armnetwork.AuthorizationUseStatusAvailable),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
