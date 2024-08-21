package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Compute/listNodes.json
func ExampleComputeClient_NewListNodesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewComputeClient().NewListNodesPager("testrg123", "workspaces123", "compute123", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Nodes {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AmlComputeNodesInformation = armmachinelearning.AmlComputeNodesInformation{
		// 	Nodes: []*armmachinelearning.AmlComputeNodeInformation{
		// 		{
		// 			NodeID: to.Ptr("tvm-3601533753_1-20170719t162906z"),
		// 			NodeState: to.Ptr(armmachinelearning.NodeStateRunning),
		// 			Port: to.Ptr[int32](50000),
		// 			PrivateIPAddress: to.Ptr("13.84.190.124"),
		// 			PublicIPAddress: to.Ptr("13.84.190.134"),
		// 			RunID: to.Ptr("2f378a44-38f2-443a-9f0d-9909d0b47890"),
		// 		},
		// 		{
		// 			NodeID: to.Ptr("tvm-3601533753_2-20170719t162906z"),
		// 			NodeState: to.Ptr(armmachinelearning.NodeStateIdle),
		// 			Port: to.Ptr[int32](50001),
		// 			PrivateIPAddress: to.Ptr("13.84.190.124"),
		// 			PublicIPAddress: to.Ptr("13.84.190.134"),
		// 	}},
		// }
	}
}
