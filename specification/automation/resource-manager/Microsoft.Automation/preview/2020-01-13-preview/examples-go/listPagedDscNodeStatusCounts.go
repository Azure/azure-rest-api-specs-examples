package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodeStatusCounts.json
func ExampleNodeCountInformationClient_Get_getNodesStatusCounts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNodeCountInformationClient().Get(ctx, "rg", "myAutomationAccount33", armautomation.CountTypeStatus, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NodeCounts = armautomation.NodeCounts{
	// 	TotalCount: to.Ptr[int32](6),
	// 	Value: []*armautomation.NodeCount{
	// 		{
	// 			Name: to.Ptr("Compliant"),
	// 			Properties: &armautomation.NodeCountProperties{
	// 				Count: to.Ptr[int32](10),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Failed"),
	// 			Properties: &armautomation.NodeCountProperties{
	// 				Count: to.Ptr[int32](1),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("InProgress"),
	// 			Properties: &armautomation.NodeCountProperties{
	// 				Count: to.Ptr[int32](1),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("NotCompliant"),
	// 			Properties: &armautomation.NodeCountProperties{
	// 				Count: to.Ptr[int32](3),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Pending"),
	// 			Properties: &armautomation.NodeCountProperties{
	// 				Count: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Unresponsive"),
	// 			Properties: &armautomation.NodeCountProperties{
	// 				Count: to.Ptr[int32](4),
	// 			},
	// 	}},
	// }
}
