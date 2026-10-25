package armedgeactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeactions/armedgeactions"
)

// Generated from example definition: 2026-10-01/EdgeActionExecutionFilters_ListByEdgeAction.json
func ExampleEdgeActionExecutionFiltersClient_NewListByEdgeActionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEdgeActionExecutionFiltersClient().NewListByEdgeActionPager("testrg", "edgeAction1", nil)
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
		// page = armedgeactions.EdgeActionExecutionFiltersClientListByEdgeActionResponse{
		// 	EdgeActionExecutionFilterListResult: armedgeactions.EdgeActionExecutionFilterListResult{
		// 		Value: []*armedgeactions.EdgeActionExecutionFilter{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1/executionFilters/executionFilter1"),
		// 				Name: to.Ptr("executionFilter1"),
		// 				Type: to.Ptr("Microsoft.Cdn/edgeActions/executionFilters"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armedgeactions.EdgeActionExecutionFilterProperties{
		// 					VersionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1/versions/version1"),
		// 					ExecutionFilterIdentifierHeaderName: to.Ptr("header-key"),
		// 					ExecutionFilterIdentifierHeaderValue: to.Ptr("header-value"),
		// 					ProvisioningState: to.Ptr(armedgeactions.ProvisioningStateSucceeded),
		// 					LastUpdateTime: to.Ptr(time.Date(2024, time.April, 25, 15, 19, 23, 0, time.UTC)),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1/executionFilters/executionFilter2"),
		// 				Name: to.Ptr("executionFilter2"),
		// 				Type: to.Ptr("Microsoft.Cdn/edgeActions/executionFilters"),
		// 				Location: to.Ptr("global"),
		// 				Properties: &armedgeactions.EdgeActionExecutionFilterProperties{
		// 					VersionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Cdn/edgeActions/edgeAction1/versions/version2"),
		// 					ExecutionFilterIdentifierHeaderName: to.Ptr("header-key2"),
		// 					ExecutionFilterIdentifierHeaderValue: to.Ptr("header-value2"),
		// 					ProvisioningState: to.Ptr(armedgeactions.ProvisioningStateSucceeded),
		// 					LastUpdateTime: to.Ptr(time.Date(2024, time.April, 25, 15, 19, 23, 0, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
