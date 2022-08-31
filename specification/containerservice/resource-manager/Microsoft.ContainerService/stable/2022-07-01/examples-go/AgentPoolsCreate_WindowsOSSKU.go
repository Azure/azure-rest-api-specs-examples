package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-07-01/examples/AgentPoolsCreate_WindowsOSSKU.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithWindowsOssku() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "wnp2", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			OrchestratorVersion: to.Ptr("1.23.3"),
			OSSKU:               to.Ptr(armcontainerservice.OSSKUWindows2022),
			OSType:              to.Ptr(armcontainerservice.OSTypeWindows),
			VMSize:              to.Ptr("Standard_D4s_v3"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
