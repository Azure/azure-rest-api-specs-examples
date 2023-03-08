package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/241e964afe675a7be98aa6a2e171a3c5f830816c/specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/Agents_CreateOrUpdate.json
func ExampleAgentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewAgentsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "examples-rg", "examples-storageMoverName", "examples-agentName", armstoragemover.Agent{
		Properties: &armstoragemover.AgentProperties{
			Description:   to.Ptr("Example Agent Description"),
			ArcResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"),
			ArcVMUUID:     to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Agent = armstoragemover.Agent{
	// 	Name: to.Ptr("examples-agentName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName"),
	// 	Properties: &armstoragemover.AgentProperties{
	// 		Description: to.Ptr("Example Agent Description"),
	// 		AgentStatus: to.Ptr(armstoragemover.AgentStatusRegistering),
	// 		AgentVersion: to.Ptr("1.0.0"),
	// 		ArcResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"),
	// 		ArcVMUUID: to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
	// 	},
	// }
}
