package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f41d0c9332078cb2ef07b749081d94915255ada5/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationReconcile.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_BeginReconcileForPrivateLinkScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().BeginReconcileForPrivateLinkScope(ctx, "my-resource-group", "my-privatelinkscope", "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.myAssociation", nil)
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
	// res.NetworkSecurityPerimeterConfigurationReconcileResult = armhybridcompute.NetworkSecurityPerimeterConfigurationReconcileResult{
	// 	Location: to.Ptr("{callbackUrl}"),
	// }
}
