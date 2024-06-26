package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/PrivateEndpointConnections/PrivateLinkResourcesGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "resourcegroup", "example-RelayNamespace-5849", "{PrivateLinkResource name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armrelay.PrivateLinkResource{
	// 	Name: to.Ptr("namespace"),
	// 	Type: to.Ptr("Microsoft.Relay/namespaces/privateLinkResources"),
	// 	ID: to.Ptr("subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/alitest/providers/Microsoft.Relay/namespaces/relay-private-endpoint-test/privateLinkResources/namespace"),
	// 	Properties: &armrelay.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("namespace"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("namespace")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.servicebus.windows.net")},
	// 			},
	// 		}
}
