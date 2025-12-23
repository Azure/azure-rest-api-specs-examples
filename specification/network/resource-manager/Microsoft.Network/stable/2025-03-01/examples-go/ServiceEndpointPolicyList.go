package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ServiceEndpointPolicyList.json
func ExampleServiceEndpointPoliciesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceEndpointPoliciesClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ServiceEndpointPolicyListResult = armnetwork.ServiceEndpointPolicyListResult{
		// 	Value: []*armnetwork.ServiceEndpointPolicy{
		// 		{
		// 			Name: to.Ptr("testServiceEndpointPolicy"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testServiceEndpointPolicy"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 				ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testServiceEndpointPolicy/serviceEndpointPolicyDefinitions/StorageServiceEndpointPolicyDefinition"),
		// 						Name: to.Ptr("StorageServiceEndpointPolicyDefinition"),
		// 						Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
		// 							Description: to.Ptr("Storage Service EndpointPolicy Definition"),
		// 							Service: to.Ptr("Microsoft.Storage"),
		// 							ServiceResources: []*string{
		// 								to.Ptr("/subscriptions/subid1"),
		// 								to.Ptr("/subscriptions/subid1resourceGroups/storageRg"),
		// 								to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
		// 							},
		// 					}},
		// 					Subnets: []*armnetwork.Subnet{
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testServiceEndpointPolicy1"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testServiceEndpointPolicy1"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 					ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testServiceEndpointPolicy1/serviceEndpointPolicyDefinitions/StorageServiceEndpointPolicyDefinition1"),
		// 							Name: to.Ptr("StorageServiceEndpointPolicyDefinition1"),
		// 							Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
		// 								Description: to.Ptr("Storage Service EndpointPolicy Definition"),
		// 								Service: to.Ptr("Microsoft.Storage"),
		// 								ServiceResources: []*string{
		// 									to.Ptr("/subscriptions/subid1"),
		// 									to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
		// 									to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
		// 								},
		// 						}},
		// 						Subnets: []*armnetwork.Subnet{
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
