package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/ServiceEndpointPolicyCreateWithDefinition.json
func ExampleServiceEndpointPoliciesClient_BeginCreateOrUpdate_createServiceEndpointPolicyWithDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceEndpointPoliciesClient().BeginCreateOrUpdate(ctx, "rg1", "testPolicy", armnetwork.ServiceEndpointPolicy{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
			ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
				{
					Name: to.Ptr("StorageServiceEndpointPolicyDefinition"),
					Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
						Description: to.Ptr("Storage Service EndpointPolicy Definition"),
						Service:     to.Ptr("Microsoft.Storage"),
						ServiceResources: []*string{
							to.Ptr("/subscriptions/subid1"),
							to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
							to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
					},
				}},
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
	// res.ServiceEndpointPolicy = armnetwork.ServiceEndpointPolicy{
	// 	Name: to.Ptr("testnsg"),
	// 	Type: to.Ptr("Microsoft.Network/ServiceEndpointPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ServiceEndpointPolicies/testpolicy"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
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
