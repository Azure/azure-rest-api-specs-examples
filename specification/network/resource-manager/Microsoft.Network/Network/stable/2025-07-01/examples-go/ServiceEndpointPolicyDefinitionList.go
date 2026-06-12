package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ServiceEndpointPolicyDefinitionList.json
func ExampleServiceEndpointPolicyDefinitionsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceEndpointPolicyDefinitionsClient().NewListByResourceGroupPager("rg1", "testPolicy", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armnetwork.ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse{
		// 	ServiceEndpointPolicyDefinitionListResult: armnetwork.ServiceEndpointPolicyDefinitionListResult{
		// 		Value: []*armnetwork.ServiceEndpointPolicyDefinition{
		// 			{
		// 				Name: to.Ptr("testDef"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testPolicy/serviceEndpointPolicyDefinitions/testDef"),
		// 				Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
		// 					Description: to.Ptr("Storage Service EndpointPolicy Definition"),
		// 					Service: to.Ptr("Microsoft.Storage"),
		// 					ServiceResources: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/storageRg"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
