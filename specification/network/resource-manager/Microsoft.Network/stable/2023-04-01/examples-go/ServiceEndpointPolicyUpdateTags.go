package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/baac183ffa684d94f697f0fc6f480e02cfb00f3d/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/ServiceEndpointPolicyUpdateTags.json
func ExampleServiceEndpointPoliciesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceEndpointPoliciesClient().UpdateTags(ctx, "rg1", "testServiceEndpointPolicy", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceEndpointPolicy = armnetwork.ServiceEndpointPolicy{
	// 	Name: to.Ptr("testServiceEndpointPolicy"),
	// 	Type: to.Ptr("Microsoft.Network/serviceEndpointPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testServiceEndpointPolicy"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
	// 			{
	// 				Name: to.Ptr("StorageServiceEndpointPolicyDefinition"),
	// 				Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
	// 					Description: to.Ptr("Storage Service EndpointPolicy Definition"),
	// 					Service: to.Ptr("Microsoft.Storage"),
	// 					ServiceResources: []*string{
	// 						to.Ptr("/subscriptions/subid1"),
	// 						to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
	// 						to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
	// 					},
	// 			}},
	// 			Subnets: []*armnetwork.Subnet{
	// 			},
	// 		},
	// 	}
}
