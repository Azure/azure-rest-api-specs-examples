package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/GetOrderByName.json
func ExampleManagementClient_GetOrderByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementClient().GetOrderByName(ctx, "TestOrderName3", "YourResourceGroupName", "eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OrderResource = armedgeorder.OrderResource{
	// 	Name: to.Ptr("TestOrderName3"),
	// 	Type: to.Ptr("Microsoft.EdgeOrder/orders"),
	// 	ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/locations/eastus/orders/TestOrderName3"),
	// 	Properties: &armedgeorder.OrderProperties{
	// 		CurrentStage: &armedgeorder.StageDetails{
	// 			StageName: to.Ptr(armedgeorder.StageNameInReview),
	// 			StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T06:05:20.852Z"); return t}()),
	// 		},
	// 		OrderItemIDs: []*string{
	// 			to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/orderItems/TestOrderItemName3")},
	// 			OrderStageHistory: []*armedgeorder.StageDetails{
	// 				{
	// 					StageName: to.Ptr(armedgeorder.StageNamePlaced),
	// 					StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T06:00:30.271Z"); return t}()),
	// 				},
	// 				{
	// 					StageName: to.Ptr(armedgeorder.StageNameInReview),
	// 					StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T06:05:20.852Z"); return t}()),
	// 			}},
	// 		},
	// 		SystemData: &armedgeorder.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
	// 		},
	// 	}
}
