package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7678455846b1000fd31db27596e4ca3d299a872/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Authorizations_CreateOrUpdate.json
func ExampleAuthorizationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAuthorizationsClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", "authorization1", armavs.ExpressRouteAuthorization{}, nil)
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
	// res.ExpressRouteAuthorization = armavs.ExpressRouteAuthorization{
	// 	Name: to.Ptr("authorization1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/authorizations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/authorizations/authorization1"),
	// 	Properties: &armavs.ExpressRouteAuthorizationProperties{
	// 		ExpressRouteAuthorizationID: to.Ptr("/subscriptions/5206f269-120b-41ef-a95b-0dce7109de61/resourceGroups/tnt34-cust-mockp02-spearj2dev/providers/Microsoft.Network/expressroutecircuits/tnt34-cust-mockp02-spearj2dev-er/authorizations/myauth"),
	// 		ExpressRouteAuthorizationKey: to.Ptr("37b0db3b-3b17-4c7b-bf76-bf13b01bcadc"),
	// 		ExpressRouteID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2/providers/Microsoft.Network/expressroutecircuits/tnt13-41a90db2-9d5e-4bd5-a77a-5ce7b58213d6-eastus2-xconnect"),
	// 		ProvisioningState: to.Ptr(armavs.ExpressRouteAuthorizationProvisioningStateSucceeded),
	// 	},
	// }
}
