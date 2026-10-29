package armdatatransfer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datatransfer/armdatatransfer"
)

// Generated from example definition: 2026-02-06-preview/listFlowsByConnection.json
func ExampleFlowsClient_NewListByConnectionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatatransfer.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFlowsClient().NewListByConnectionPager("testRG", "testConnection", nil)
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
		// page = armdatatransfer.FlowsClientListByConnectionResponse{
		// 	FlowListResult: armdatatransfer.FlowListResult{
		// 		Value: []*armdatatransfer.Flow{
		// 			{
		// 				Name: to.Ptr("testConnection"),
		// 				Type: to.Ptr("Microsoft.AzureDataTransfer/connections/flows"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection/flows/testFlow"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armdatatransfer.FlowProperties{
		// 					Connection: &armdatatransfer.SelectedResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testFlow"),
		// 					},
		// 					FlowProfile: &armdatatransfer.FlowProfileMetadata{
		// 						Pipeline: to.Ptr("pipelinea"),
		// 						Name: to.Ptr("flowProfileNamea"),
		// 						FlowProfileID: to.Ptr("25f2c41a-6e29-4c40-a033-b4de81d4a681"),
		// 						ReplicationScenario: to.Ptr(armdatatransfer.DataClassTypeFiles),
		// 						Status: to.Ptr(armdatatransfer.FlowProfileStatusEnabled),
		// 						Description: to.Ptr("Testing 123"),
		// 					},
		// 					StorageAccountName: to.Ptr("testsa"),
		// 					StorageContainerName: to.Ptr("testcontainer"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testConnection2"),
		// 				Type: to.Ptr("Microsoft.AzureDataTransfer/connections/flows"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection2/flows/testFlow"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armdatatransfer.FlowProperties{
		// 					Connection: &armdatatransfer.SelectedResource{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testFlow"),
		// 					},
		// 					FlowProfile: &armdatatransfer.FlowProfileMetadata{
		// 						Pipeline: to.Ptr("pipelinea"),
		// 						Name: to.Ptr("flowProfileNamea"),
		// 						FlowProfileID: to.Ptr("25f2c41a-6e29-4c40-a033-b4de81d4a681"),
		// 						ReplicationScenario: to.Ptr(armdatatransfer.DataClassTypeFiles),
		// 						Status: to.Ptr(armdatatransfer.FlowProfileStatusEnabled),
		// 						Description: to.Ptr("Testing 123"),
		// 					},
		// 					StorageAccountName: to.Ptr("testsa"),
		// 					StorageContainerName: to.Ptr("testcontainer"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
