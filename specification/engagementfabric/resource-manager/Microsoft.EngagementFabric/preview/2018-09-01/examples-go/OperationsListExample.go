package armengagementfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/engagementfabric/resource-manager/Microsoft.EngagementFabric/preview/2018-09-01/examples/OperationsListExample.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armengagementfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationList = armengagementfabric.OperationList{
		// 	Value: []*armengagementfabric.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/Accounts/read"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("List or get the EngagementFabric account"),
		// 				Operation: to.Ptr("List or get the EngagementFabric account"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/Accounts/write"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("Create or update the EngagementFabric account"),
		// 				Operation: to.Ptr("Create or update the EngagementFabric account"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/Accounts/delete"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("Delete the EngagementFabric account"),
		// 				Operation: to.Ptr("Delete the EngagementFabric account"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/Accounts/ListKeys/action"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("Get all keys of the EngagementFabric account"),
		// 				Operation: to.Ptr("Get all keys of the EngagementFabric account"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/Accounts/RegenerateKey/action"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("Regenerate the EngagementFabric account key"),
		// 				Operation: to.Ptr("Regenerate the EngagementFabric account key"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/Accounts/ListChannelTypes/action"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("List available EngagementFabric channel types and functions"),
		// 				Operation: to.Ptr("List available EngagementFabric channel types and functions"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/Accounts/Channels/read"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("List or get the EngagementFabric channel"),
		// 				Operation: to.Ptr("List or get the EngagementFabric channel"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Channels"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/Accounts/Channels/write"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("Create or update the EngagementFabric channel"),
		// 				Operation: to.Ptr("Create or update the EngagementFabric channel"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Channels"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/Accounts/Channels/delete"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("Delete the EngagementFabric channel"),
		// 				Operation: to.Ptr("Delete the EngagementFabric channel"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Channels"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/checkNameAvailability/action"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("Check name availability"),
		// 				Operation: to.Ptr("Check name availability"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Accounts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/operations/read"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("List available operations"),
		// 				Operation: to.Ptr("List available operations"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.EngagementFabric/skus/read"),
		// 			Display: &armengagementfabric.OperationDisplay{
		// 				Description: to.Ptr("List available SKUs"),
		// 				Operation: to.Ptr("List available SKUs"),
		// 				Provider: to.Ptr("Microsoft Customer Engagement Fabric"),
		// 				Resource: to.Ptr("Accounts"),
		// 			},
		// 	}},
		// }
	}
}
