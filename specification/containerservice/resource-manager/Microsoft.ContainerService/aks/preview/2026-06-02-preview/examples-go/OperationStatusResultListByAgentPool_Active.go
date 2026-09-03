package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v9"
)

// Generated from example definition: 2026-06-02-preview/OperationStatusResultListByAgentPool_Active.json
func ExampleOperationStatusResultClient_NewListByAgentPoolPager_listActiveOperationsOnAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationStatusResultClient().NewListByAgentPoolPager("rg1", "clustername1", "agentpool1", &armcontainerservice.OperationStatusResultClientListByAgentPoolOptions{
		ActiveOnly: to.Ptr(true)})
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
		// page = armcontainerservice.OperationStatusResultClientListByAgentPoolResponse{
		// 	OperationStatusResultList: armcontainerservice.OperationStatusResultList{
		// 		Value: []*armcontainerservice.OperationStatusResult{
		// 			{
		// 				Name: to.Ptr("d11edb09-6e27-429f-9fe5-17baf773bc4a"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1/operations/d11edb09-6e27-429f-9fe5-17baf773bc4a"),
		// 				Status: to.Ptr("ScalingAgentPool: 3/5 nodes completed"),
		// 				PercentComplete: to.Ptr[float64](60),
		// 				StartTime: to.Ptr(time.Date(2026, time.May, 6, 10, 5, 0, 0, time.UTC)),
		// 				OperationType: to.Ptr("PatchAgentPool"),
		// 				SubOperationType: to.Ptr("Scaling"),
		// 			},
		// 			{
		// 				Name: to.Ptr("a22fdc10-7e38-530f-a6f6-28cbe884cd5b"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1/operations/a22fdc10-7e38-530f-a6f6-28cbe884cd5b"),
		// 				Status: to.Ptr("UpgradingNodes: 8/20 nodes completed"),
		// 				PercentComplete: to.Ptr[float64](40),
		// 				StartTime: to.Ptr(time.Date(2026, time.May, 6, 10, 0, 0, 0, time.UTC)),
		// 				OperationType: to.Ptr("PutAgentPool"),
		// 				SubOperationType: to.Ptr("Upgrading"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
