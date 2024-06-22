package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/ServiceEndpointPolicyDefinitionCreate.json
func ExampleServiceEndpointPolicyDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceEndpointPolicyDefinitionsClient().BeginCreateOrUpdate(ctx, "rg1", "testPolicy", "testDefinition", armnetwork.ServiceEndpointPolicyDefinition{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceEndpointPolicyDefinition = armnetwork.ServiceEndpointPolicyDefinition{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testPolicy/serviceEndpointPolicyDefinitions/testDefinition"),
	// 	Name: to.Ptr("testDefinition"),
	// 	Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
	// 		Description: to.Ptr("Storage Service EndpointPolicy Definition"),
	// 		Service: to.Ptr("Microsoft.Storage"),
	// 		ServiceResources: []*string{
	// 			to.Ptr("/subscriptions/subid1"),
	// 			to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
	// 			to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
	// 		},
	// 	}
}
