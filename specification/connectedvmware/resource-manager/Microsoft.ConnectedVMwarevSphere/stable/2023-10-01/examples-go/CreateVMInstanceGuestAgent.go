package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/CreateVMInstanceGuestAgent.json
func ExampleVMInstanceGuestAgentsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVMInstanceGuestAgentsClient().BeginCreate(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM", armconnectedvmware.GuestAgent{
		Properties: &armconnectedvmware.GuestAgentProperties{
			Credentials: &armconnectedvmware.GuestCredential{
				Password: to.Ptr("<password>"),
				Username: to.Ptr("tempuser"),
			},
			HTTPProxyConfig: &armconnectedvmware.HTTPProxyConfiguration{
				HTTPSProxy: to.Ptr("http://192.1.2.3:8080"),
			},
			PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
			ProvisioningAction:         to.Ptr(armconnectedvmware.ProvisioningActionInstall),
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
	// res.GuestAgent = armconnectedvmware.GuestAgent{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VirtualMachineInstances/guestAgents"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineInstances/default/guestAgents/default"),
	// 	Properties: &armconnectedvmware.GuestAgentProperties{
	// 		ProvisioningAction: to.Ptr(armconnectedvmware.ProvisioningActionInstall),
	// 		ProvisioningState: to.Ptr(armconnectedvmware.ProvisioningStateSucceeded),
	// 		Status: to.Ptr("connected"),
	// 	},
	// }
}
