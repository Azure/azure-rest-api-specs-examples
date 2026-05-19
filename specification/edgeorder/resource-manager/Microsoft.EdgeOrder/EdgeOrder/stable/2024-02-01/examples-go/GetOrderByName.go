package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder/v2"
)

// Generated from example definition: 2024-02-01/GetOrderByName.json
func ExampleOrdersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("eb5dc900-6186-49d8-b7d7-febd866fdc1d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrdersClient().Get(ctx, "YourResourceGroupName", "eastus", "TestOrderName3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armedgeorder.OrdersClientGetResponse{
	// 	OrderResource: &armedgeorder.OrderResource{
	// 		Name: to.Ptr("TestOrderName3"),
	// 		Type: to.Ptr("Microsoft.EdgeOrder/orders"),
	// 		ID: to.Ptr("/subscriptions/eb5dc900-6186-49d8-b7d7-febd866fdc1d/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/locations/eastus/orders/TestOrderName3"),
	// 		Properties: &armedgeorder.OrderProperties{
	// 			CurrentStage: &armedgeorder.StageDetails{
	// 				StageName: to.Ptr(armedgeorder.StageNameInReview),
	// 				StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T11:35:20.8521455+05:30"); return t}()),
	// 			},
	// 			OrderItemIDs: []*string{
	// 				to.Ptr("/subscriptions/eb5dc900-6186-49d8-b7d7-febd866fdc1d/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/orderItems/TestOrderItemName3"),
	// 			},
	// 			OrderStageHistory: []*armedgeorder.StageDetails{
	// 				{
	// 					StageName: to.Ptr(armedgeorder.StageNamePlaced),
	// 					StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T11:30:30.2717243+05:30"); return t}()),
	// 				},
	// 				{
	// 					StageName: to.Ptr(armedgeorder.StageNameInReview),
	// 					StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T11:35:20.8521455+05:30"); return t}()),
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armedgeorder.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
	// 		},
	// 	},
	// }
}
