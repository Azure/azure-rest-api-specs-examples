package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0a25ea9680cf080b7d34e8c5f35f564425c6b1f7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-08-01/examples/AgentPoolsCreate_WindowsOSSKU.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithWindowsOssku() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolsClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "wnp2", armcontainerservice.AgentPool{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPool = armcontainerservice.AgentPool{
	// 	Name: to.Ptr("wnp2"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/wnp2"),
	// 	Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
	// 		Count: to.Ptr[int32](3),
	// 		CurrentOrchestratorVersion: to.Ptr("1.23.3"),
	// 		MaxPods: to.Ptr[int32](110),
	// 		OrchestratorVersion: to.Ptr("1.23.3"),
	// 		OSSKU: to.Ptr(armcontainerservice.OSSKUWindows2022),
	// 		OSType: to.Ptr(armcontainerservice.OSTypeWindows),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		VMSize: to.Ptr("Standard_D4s_v3"),
	// 	},
	// }
}
