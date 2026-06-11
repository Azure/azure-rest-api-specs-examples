package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ServiceEndpointPolicyDefinitionCreate.json
func ExampleServiceEndpointPolicyDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceEndpointPolicyDefinitionsClient().BeginCreateOrUpdate(ctx, "rg1", "testPolicy", "testDefinition", armnetwork.ServiceEndpointPolicyDefinition{
		Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
			Description: to.Ptr("Storage Service EndpointPolicy Definition"),
			Service:     to.Ptr("Microsoft.Storage"),
			ServiceResources: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/storageRg"),
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ServiceEndpointPolicyDefinitionsClientCreateOrUpdateResponse{
	// 	ServiceEndpointPolicyDefinition: armnetwork.ServiceEndpointPolicyDefinition{
	// 		Name: to.Ptr("testDefinition"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testPolicy/serviceEndpointPolicyDefinitions/testDefinition"),
	// 		Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
	// 			Description: to.Ptr("Storage Service EndpointPolicy Definition"),
	// 			Service: to.Ptr("Microsoft.Storage"),
	// 			ServiceResources: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/storageRg"),
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount"),
	// 			},
	// 		},
	// 	},
	// }
}
