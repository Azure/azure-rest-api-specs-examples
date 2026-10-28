package armdatatransfer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datatransfer/armdatatransfer"
)

// Generated from example definition: 2026-02-06-preview/listFlowsByPipeline.json
func ExamplePipelinesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatatransfer.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPipelinesClient().List(ctx, "testRG", "testPipeline", &armdatatransfer.PipelinesClientListOptions{
		Connections: &armdatatransfer.ConnectionIDList{
			Value: []*string{
				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection"),
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatatransfer.PipelinesClientListResponse{
	// 	ListFlowsByPipelineResult: armdatatransfer.ListFlowsByPipelineResult{
	// 		Value: []*armdatatransfer.ListFlowsByPipelineConnection{
	// 			{
	// 				Flows: []*armdatatransfer.Flow{
	// 					{
	// 						Type: to.Ptr("Microsoft.AzureDataTransfer/connections/flows"),
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection/flows/testFlow"),
	// 						Location: to.Ptr("East US"),
	// 						Properties: &armdatatransfer.FlowProperties{
	// 							FlowProfile: &armdatatransfer.FlowProfileMetadata{
	// 								Pipeline: to.Ptr("testPipeline"),
	// 								Name: to.Ptr("flowProfileName"),
	// 								FlowProfileID: to.Ptr("25f2c41a-6e29-4c40-a033-b4de81d4a681"),
	// 								ReplicationScenario: to.Ptr(armdatatransfer.DataClassTypeFiles),
	// 								Status: to.Ptr(armdatatransfer.FlowProfileStatusEnabled),
	// 								Description: to.Ptr("Testing 123"),
	// 							},
	// 							StorageAccountName: to.Ptr("testsa"),
	// 							StorageContainerName: to.Ptr("testcontainer"),
	// 						},
	// 					},
	// 					{
	// 						Type: to.Ptr("Microsoft.AzureDataTransfer/connections/flows"),
	// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection/flows/testFlow2"),
	// 						Location: to.Ptr("East US"),
	// 						Properties: &armdatatransfer.FlowProperties{
	// 							FlowProfile: &armdatatransfer.FlowProfileMetadata{
	// 								Pipeline: to.Ptr("testPipeline"),
	// 								Name: to.Ptr("flowProfileName"),
	// 								FlowProfileID: to.Ptr("25f2c41a-6e29-4c40-a033-b4de81d4a681"),
	// 								ReplicationScenario: to.Ptr(armdatatransfer.DataClassTypeFiles),
	// 								Status: to.Ptr(armdatatransfer.FlowProfileStatusEnabled),
	// 								Description: to.Ptr("Testing 123"),
	// 							},
	// 							StorageAccountName: to.Ptr("testsa"),
	// 							StorageContainerName: to.Ptr("testcontainer"),
	// 						},
	// 					},
	// 				},
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection"),
	// 			},
	// 		},
	// 	},
	// }
}
