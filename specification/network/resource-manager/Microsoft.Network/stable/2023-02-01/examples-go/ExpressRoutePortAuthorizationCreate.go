package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/ExpressRoutePortAuthorizationCreate.json
func ExampleExpressRoutePortAuthorizationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRoutePortAuthorizationsClient().BeginCreateOrUpdate(ctx, "rg1", "expressRoutePortName", "authorizatinName", armnetwork.ExpressRoutePortAuthorization{
		Properties: &armnetwork.ExpressRoutePortAuthorizationPropertiesFormat{},
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
	// res.ExpressRoutePortAuthorization = armnetwork.ExpressRoutePortAuthorization{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ExpressRoutePorts/expressRoutePortName/authorizations/authorizationName"),
	// 	Name: to.Ptr("authorizationName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRoutePorts/authorizations"),
	// 	Etag: to.Ptr("W/\"e22dd4b2-4c24-44cf-b702-70a472b62914\""),
	// 	Properties: &armnetwork.ExpressRoutePortAuthorizationPropertiesFormat{
	// 		AuthorizationUseStatus: to.Ptr(armnetwork.ExpressRoutePortAuthorizationUseStatusAvailable),
	// 		CircuitResourceURI: to.Ptr(""),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
