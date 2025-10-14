package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: 2025-07-01/Agents_CreateOrUpdate_MinimumSet.json
func ExampleAgentsClient_CreateOrUpdate_agentsCreateOrUpdateMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("60bcfc77-6589-4da2-b7fd-f9ec9322cf95", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentsClient().CreateOrUpdate(ctx, "examples-rg", "examples-storageMoverName", "examples-agentName", armstoragemover.Agent{
		Properties: &armstoragemover.AgentProperties{
			ArcResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"),
			ArcVMUUID:     to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragemover.AgentsClientCreateOrUpdateResponse{
	// 	Agent: &armstoragemover.Agent{
	// 		Name: to.Ptr("examples-agentName"),
	// 		Type: to.Ptr("Microsoft.StorageMover/storageMovers/agents"),
	// 		ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/examples-agentName"),
	// 		Properties: &armstoragemover.AgentProperties{
	// 			AgentStatus: to.Ptr(armstoragemover.AgentStatusOnline),
	// 			AgentVersion: to.Ptr("1.0.0"),
	// 			ArcResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName"),
	// 			ArcVMUUID: to.Ptr("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9"),
	// 			LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:21:01.1075056Z"); return t}()),
	// 			LocalIPAddress: to.Ptr("192.168.0.0"),
	// 			MemoryInMB: to.Ptr[int64](4096),
	// 			NumberOfCores: to.Ptr[int64](8),
	// 			ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
	// 			TimeZone: to.Ptr("Eastern Standard Time"),
	// 			UptimeInSeconds: to.Ptr[int64](522),
	// 		},
	// 		SystemData: &armstoragemover.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T01:01:01.1075056Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:01:01.1075056Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
