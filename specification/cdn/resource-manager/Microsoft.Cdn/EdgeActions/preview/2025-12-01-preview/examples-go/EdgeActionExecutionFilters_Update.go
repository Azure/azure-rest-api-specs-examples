package armedgeactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeactions/armedgeactions"
)

// Generated from example definition: 2025-12-01-preview/EdgeActionExecutionFilters_Update.json
func ExampleEdgeActionExecutionFiltersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEdgeActionExecutionFiltersClient().BeginUpdate(ctx, "testrg", "edgeAction1", "executionFilter1", armedgeactions.EdgeActionExecutionFilterUpdate{
		Properties: &armedgeactions.EdgeActionExecutionFilterUpdateProperties{
			ExecutionFilterIdentifierHeaderValue: to.Ptr("header-value2"),
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
	// res = armedgeactions.EdgeActionExecutionFiltersClientUpdateResponse{
	// 	EdgeActionExecutionFilter: &armedgeactions.EdgeActionExecutionFilter{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1/slots/slot1"),
	// 		Name: to.Ptr("executionFilter1"),
	// 		Type: to.Ptr("Microsoft.Cdn/edgeActions/executionFilters"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armedgeactions.EdgeActionExecutionFilterProperties{
	// 			VersionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/testrg/providers/Microsoft.Cdn/EdgeActions/edgeAction1/versions/version1"),
	// 			ExecutionFilterIdentifierHeaderName: to.Ptr("header-key"),
	// 			ExecutionFilterIdentifierHeaderValue: to.Ptr("header-value2"),
	// 			ProvisioningState: to.Ptr(armedgeactions.ProvisioningStateSucceeded),
	// 			LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-25T15:19:23Z"); return t}()),
	// 		},
	// 	},
	// }
}
