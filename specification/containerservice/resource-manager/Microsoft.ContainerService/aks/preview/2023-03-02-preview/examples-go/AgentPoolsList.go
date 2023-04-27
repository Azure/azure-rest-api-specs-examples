package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17b0396c5ce59d801b9eb38fd6c47d149f6db2a2/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-03-02-preview/examples/AgentPoolsList.json
func ExampleAgentPoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAgentPoolsClient().NewListPager("rg1", "clustername1", nil)
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
		// page.AgentPoolListResult = armcontainerservice.AgentPoolListResult{
		// 	Value: []*armcontainerservice.AgentPool{
		// 		{
		// 			Name: to.Ptr("agentpool1"),
		// 			ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1"),
		// 			Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
		// 				Count: to.Ptr[int32](3),
		// 				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
		// 				MaxPods: to.Ptr[int32](110),
		// 				NodeImageVersion: to.Ptr("AKSUbuntu:1604:2020.03.11"),
		// 				OrchestratorVersion: to.Ptr("1.9.6"),
		// 				OSType: to.Ptr(armcontainerservice.OSTypeLinux),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				VMSize: to.Ptr("Standard_DS1_v2"),
		// 			},
		// 	}},
		// }
	}
}
