package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/preview/2022-05-01-preview/examples/EndpointsPutDefault.json
func ExampleEndpointsClient_CreateOrUpdate_hybridConnectivityEndpointsPutDefault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().CreateOrUpdate(ctx, "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine", "default", armhybridconnectivity.EndpointResource{
		Properties: &armhybridconnectivity.EndpointProperties{
			Type: to.Ptr(armhybridconnectivity.TypeDefault),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EndpointResource = armhybridconnectivity.EndpointResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.HybridConnectivity/endpoints"),
	// 	ID: to.Ptr("/subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default"),
	// 	Properties: &armhybridconnectivity.EndpointProperties{
	// 		Type: to.Ptr(armhybridconnectivity.TypeDefault),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
