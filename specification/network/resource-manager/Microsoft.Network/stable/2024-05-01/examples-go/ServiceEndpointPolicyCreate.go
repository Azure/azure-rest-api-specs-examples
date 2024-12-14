package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ab04533261eff228f28e08900445d0edef3eb70c/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/ServiceEndpointPolicyCreate.json
func ExampleServiceEndpointPoliciesClient_BeginCreateOrUpdate_createServiceEndpointPolicy() {
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
	// 		},
	// 		Subnets: []*armnetwork.Subnet{
	// 		},
	// 	},
	// }
}
