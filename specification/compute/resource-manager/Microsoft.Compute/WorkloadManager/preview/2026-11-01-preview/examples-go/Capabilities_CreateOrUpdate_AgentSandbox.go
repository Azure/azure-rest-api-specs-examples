package armworkloadmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armworkloadmanager"
)

// Generated from example definition: 2026-11-01-preview/Capabilities_CreateOrUpdate_AgentSandbox.json
func ExampleCapabilitiesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadmanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCapabilitiesClient().BeginCreateOrUpdate(ctx, "rg-workload", "managed-agents-prod", "agentSandbox", armworkloadmanager.Capability{
		Location: to.Ptr("eastus2"),
		Kind:     to.Ptr(armworkloadmanager.CapabilityKindAgentSandbox),
		Properties: &armworkloadmanager.CapabilityProperties{
			VersionPolicy: to.Ptr(armworkloadmanager.VersionPolicyServiceManaged),
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
	// res = armworkloadmanager.CapabilitiesClientCreateOrUpdateResponse{
	// 	Capability: armworkloadmanager.Capability{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-workload/providers/Microsoft.Compute/workloadSpaces/managed-agents-prod/capabilities/agentSandbox"),
	// 		Name: to.Ptr("agentSandbox"),
	// 		Type: to.Ptr("Microsoft.Compute/workloadSpaces/capabilities"),
	// 		Location: to.Ptr("eastus2"),
	// 		Kind: to.Ptr(armworkloadmanager.CapabilityKindAgentSandbox),
	// 		Properties: &armworkloadmanager.CapabilityProperties{
	// 			VersionPolicy: to.Ptr(armworkloadmanager.VersionPolicyServiceManaged),
	// 			EffectiveVersion: to.Ptr("service-selected"),
	// 			ProvisioningState: to.Ptr(armworkloadmanager.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
