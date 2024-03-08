package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/InitiateJitNetworkAccessPolicy_example.json
func ExampleJitNetworkAccessPoliciesClient_Initiate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewJitNetworkAccessPoliciesClient().Initiate(ctx, "myRg1", "westeurope", "default", armsecurity.JitNetworkAccessPolicyInitiateRequest{
		Justification: to.Ptr("testing a new version of the product"),
		VirtualMachines: []*armsecurity.JitNetworkAccessPolicyInitiateVirtualMachine{
			{
				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1"),
				Ports: []*armsecurity.JitNetworkAccessPolicyInitiatePort{
					{
						AllowedSourceAddressPrefix: to.Ptr("192.127.0.2"),
						Number:                     to.Ptr[int32](3389),
					}},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
