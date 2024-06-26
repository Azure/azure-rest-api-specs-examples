package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b9403296f0b0e112b0d8222ad05fd1d79ee10e03/specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/GetOperationsSubscription.json
func ExampleClient_NewListSubscriptionOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListSubscriptionOperationsPager(nil)
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
		// page.Operations = armmaps.Operations{
		// 	Value: []*armmaps.OperationDetail{
		// 		{
		// 			Name: to.Ptr("Microsoft.Maps/register/action"),
		// 			Display: &armmaps.OperationDisplay{
		// 				Description: to.Ptr("Register the provider"),
		// 				Operation: to.Ptr("Register the provider"),
		// 				Provider: to.Ptr("Microsoft Maps"),
		// 				Resource: to.Ptr("Maps Account"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Maps/accounts/write"),
		// 			Display: &armmaps.OperationDisplay{
		// 				Description: to.Ptr("Create or update a Maps Account."),
		// 				Operation: to.Ptr("Create or update a Maps Account."),
		// 				Provider: to.Ptr("Microsoft Maps"),
		// 				Resource: to.Ptr("Maps Account"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Maps/accounts/read"),
		// 			Display: &armmaps.OperationDisplay{
		// 				Description: to.Ptr("Get a Maps Account."),
		// 				Operation: to.Ptr("Get a Maps Account."),
		// 				Provider: to.Ptr("Microsoft Maps"),
		// 				Resource: to.Ptr("Maps Account"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Maps/accounts/delete"),
		// 			Display: &armmaps.OperationDisplay{
		// 				Description: to.Ptr("Delete a Maps Account."),
		// 				Operation: to.Ptr("Delete a Maps Account."),
		// 				Provider: to.Ptr("Microsoft Maps"),
		// 				Resource: to.Ptr("Maps Account"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Maps/accounts/listKeys/action"),
		// 			Display: &armmaps.OperationDisplay{
		// 				Description: to.Ptr("List Maps Account keys"),
		// 				Operation: to.Ptr("List keys"),
		// 				Provider: to.Ptr("Microsoft Maps"),
		// 				Resource: to.Ptr("Maps Account"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Maps/accounts/regenerateKey/action"),
		// 			Display: &armmaps.OperationDisplay{
		// 				Description: to.Ptr("Generate new Maps Account primary or secondary key"),
		// 				Operation: to.Ptr("Generate new primary or secondary key"),
		// 				Provider: to.Ptr("Microsoft Maps"),
		// 				Resource: to.Ptr("Maps Account"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Maps/accounts/creators/write"),
		// 			Display: &armmaps.OperationDisplay{
		// 				Description: to.Ptr("Create or update a Maps Creator."),
		// 				Operation: to.Ptr("Create or update a Maps Creator."),
		// 				Provider: to.Ptr("Microsoft Maps"),
		// 				Resource: to.Ptr("Creator"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Maps/accounts/creators/read"),
		// 			Display: &armmaps.OperationDisplay{
		// 				Description: to.Ptr("Get a Maps Creator."),
		// 				Operation: to.Ptr("Get a Maps Creator."),
		// 				Provider: to.Ptr("Microsoft Maps"),
		// 				Resource: to.Ptr("Creator"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Maps/accounts/creators/delete"),
		// 			Display: &armmaps.OperationDisplay{
		// 				Description: to.Ptr("Delete a Maps Creator."),
		// 				Operation: to.Ptr("Delete a Maps Creator."),
		// 				Provider: to.Ptr("Microsoft Maps"),
		// 				Resource: to.Ptr("Creator"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
