package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodeConfigurationCounts.json
func ExampleNodeCountInformationClient_Get_getNodesNodeConfigurationCounts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNodeCountInformationClient().Get(ctx, "rg", "myAutomationAccount33", armautomation.CountTypeNodeconfiguration, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NodeCounts = armautomation.NodeCounts{
	// 	TotalCount: to.Ptr[int32](16),
	// 	Value: []*armautomation.NodeCount{
	// 		{
	// 			Name: to.Ptr("client.localhost"),
	// 			Properties: &armautomation.NodeCountProperties{
	// 				Count: to.Ptr[int32](24),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("server.localhost"),
	// 			Properties: &armautomation.NodeCountProperties{
	// 				Count: to.Ptr[int32](3),
	// 			},
	// 	}},
	// }
}
