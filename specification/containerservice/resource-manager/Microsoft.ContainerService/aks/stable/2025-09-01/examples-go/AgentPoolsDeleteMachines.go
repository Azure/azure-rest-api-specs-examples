package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c86c7e36fb15171a9967d9fdc47784f2e4202ca6/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-09-01/examples/AgentPoolsDeleteMachines.json
func ExampleAgentPoolsClient_BeginDeleteMachines() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolsClient().BeginDeleteMachines(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPoolDeleteMachinesParameter{
		MachineNames: []*string{
			to.Ptr("aks-nodepool1-42263519-vmss00000a"),
			to.Ptr("aks-nodepool1-42263519-vmss00000b")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
