package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ServiceEndpointPolicyDefinitionCreate.json
func ExampleServiceEndpointPolicyDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewServiceEndpointPolicyDefinitionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "testPolicy", "testDefinition", armnetwork.ServiceEndpointPolicyDefinition{
		Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
			Description: to.Ptr("Storage Service EndpointPolicy Definition"),
			Service:     to.Ptr("Microsoft.Storage"),
			ServiceResources: []*string{
				to.Ptr("/subscriptions/subid1"),
				to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
				to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
