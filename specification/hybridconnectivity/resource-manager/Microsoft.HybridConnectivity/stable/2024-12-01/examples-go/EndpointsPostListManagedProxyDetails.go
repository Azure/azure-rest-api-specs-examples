package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/EndpointsPostListManagedProxyDetails.json
func ExampleEndpointsClient_ListManagedProxyDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().ListManagedProxyDetails(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/arcGroup/providers/Microsoft.Compute/virtualMachines/vm00006", "default", armhybridconnectivity.ManagedProxyRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.EndpointsClientListManagedProxyDetailsResponse{
	// 	ManagedProxyResource: &armhybridconnectivity.ManagedProxyResource{
	// 		ExpiresOn: to.Ptr[int64](1620000256),
	// 		Proxy: to.Ptr("uid.r.proxy.arc.com"),
	// 	},
	// }
}
