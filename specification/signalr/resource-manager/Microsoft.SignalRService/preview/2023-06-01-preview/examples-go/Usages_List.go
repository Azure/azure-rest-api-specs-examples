package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/Usages_List.json
func ExampleUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("eastus", nil)
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
		// page.UsageList = armsignalr.UsageList{
		// 	Value: []*armsignalr.Usage{
		// 		{
		// 			Name: &armsignalr.UsageName{
		// 				LocalizedValue: to.Ptr("Usage1"),
		// 				Value: to.Ptr("Usage1"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.SignalRService/locations/eastus/usages/Usage1"),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr("Count"),
		// 		},
		// 		{
		// 			Name: &armsignalr.UsageName{
		// 				LocalizedValue: to.Ptr("Usage2"),
		// 				Value: to.Ptr("Usage2"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.SignalRService/locations/eastus/usages/Usage2"),
		// 			Limit: to.Ptr[int64](100),
		// 			Unit: to.Ptr("Count"),
		// 	}},
		// }
	}
}
