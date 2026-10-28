package armdatatransfer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datatransfer/armdatatransfer"
)

// Generated from example definition: 2026-02-06-preview/createOrUpdatePipeline.json
func ExamplePipelinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatatransfer.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPipelinesClient().BeginCreateOrUpdate(ctx, "testRG", "testPipeline", armdatatransfer.Pipeline{
		Location: to.Ptr("East US"),
		Properties: &armdatatransfer.PipelineProperties{
			RemoteCloud: to.Ptr("testdc"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatatransfer.PipelinesClientCreateOrUpdateResponse{
	// 	Pipeline: armdatatransfer.Pipeline{
	// 		Type: to.Ptr("Microsoft.AzureDataTransfer/pipelines"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/pipelines/testPipeline"),
	// 		Location: to.Ptr("East US"),
	// 	},
	// }
}
