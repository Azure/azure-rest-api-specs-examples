package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/communication/resource-manager/Microsoft.Communication/preview/2023-06-01-preview/examples/suppressionLists/getSuppressionLists.json
func ExampleSuppressionListsClient_NewListByDomainPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSuppressionListsClient().NewListByDomainPager("contosoResourceGroup", "contosoEmailService", "contoso.com", nil)
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
		// page.SuppressionListResourceCollection = armcommunication.SuppressionListResourceCollection{
		// 	Value: []*armcommunication.SuppressionListResource{
		// 		{
		// 			Name: to.Ptr("aaaa1111-bbbb-2222-3333-aaaa11112222"),
		// 			Type: to.Ptr("Microsoft.Communication/EmailServices/Domains/SuppressionLists"),
		// 			ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/EmailServices/contosoEmailService/Domains/contoso.com/suppressionLists/aaaa1111-bbbb-2222-3333-aaaa11112222"),
		// 			Properties: &armcommunication.SuppressionListProperties{
		// 				CreatedTimeStamp: to.Ptr("2023-06-06T17:06:26.100Z"),
		// 				LastUpdatedTimeStamp: to.Ptr("2023-06-06T17:06:26.100Z"),
		// 				ListName: to.Ptr("contosoNewsAlerts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("aaaa1111-bbbb-2222-3333-aaaa11112222"),
		// 			Type: to.Ptr("Microsoft.Communication/EmailServices/Domains/SuppressionLists"),
		// 			ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/contosoResourceGroup/providers/Microsoft.Communication/EmailServices/contosoEmailService/Domains/contoso.com/suppressionLists/aaaa1111-bbbb-2222-3333-aaaa11112222"),
		// 			Properties: &armcommunication.SuppressionListProperties{
		// 				CreatedTimeStamp: to.Ptr("2023-06-06T17:06:26.100Z"),
		// 				LastUpdatedTimeStamp: to.Ptr("2023-06-06T17:06:26.100Z"),
		// 				ListName: to.Ptr("contosoShopping"),
		// 			},
		// 	}},
		// }
	}
}
