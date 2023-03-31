package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderAtResourceGroupLevel.json
func ExampleManagementClient_NewListOrderAtResourceGroupLevelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewListOrderAtResourceGroupLevelPager("YourResourceGroupName", &armedgeorder.ManagementClientListOrderAtResourceGroupLevelOptions{SkipToken: nil})
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
		// page.OrderResourceList = armedgeorder.OrderResourceList{
		// 	Value: []*armedgeorder.OrderResource{
		// 		{
		// 			Name: to.Ptr("TestOrderItemName1"),
		// 			Type: to.Ptr("Microsoft.EdgeOrder/orders"),
		// 			ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/locations/eastus/orders/TestOrderName1"),
		// 			Properties: &armedgeorder.OrderProperties{
		// 				CurrentStage: &armedgeorder.StageDetails{
		// 					StageName: to.Ptr(armedgeorder.StageNamePlaced),
		// 					StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T10:55:46.9437439+05:30"); return t}()),
		// 				},
		// 				OrderItemIDs: []*string{
		// 					to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/orderItems/TestOrderItemName1")},
		// 					OrderStageHistory: []*armedgeorder.StageDetails{
		// 						{
		// 							StageName: to.Ptr(armedgeorder.StageNamePlaced),
		// 							StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T10:55:46.9437439+05:30"); return t}()),
		// 						},
		// 						{
		// 							StageName: to.Ptr(armedgeorder.StageNameInReview),
		// 							StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 					}},
		// 				},
		// 				SystemData: &armedgeorder.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("TestOrderItemName2"),
		// 				Type: to.Ptr("Microsoft.EdgeOrder/orders"),
		// 				ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/locations/eastus/orders/TestOrderName2"),
		// 				Properties: &armedgeorder.OrderProperties{
		// 					CurrentStage: &armedgeorder.StageDetails{
		// 						StageName: to.Ptr(armedgeorder.StageNamePlaced),
		// 						StageStatus: to.Ptr(armedgeorder.StageStatusInProgress),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T10:58:46.5241979+05:30"); return t}()),
		// 					},
		// 					OrderItemIDs: []*string{
		// 						to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/orderItems/TestOrderItemName2")},
		// 						OrderStageHistory: []*armedgeorder.StageDetails{
		// 							{
		// 								StageName: to.Ptr(armedgeorder.StageNamePlaced),
		// 								StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 							},
		// 							{
		// 								StageName: to.Ptr(armedgeorder.StageNameInReview),
		// 								StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 						}},
		// 					},
		// 					SystemData: &armedgeorder.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 					},
		// 			}},
		// 		}
	}
}
