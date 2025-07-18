package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/ListResourceHealthMetadataBySite.json
func ExampleResourceHealthMetadataClient_NewListBySiteSlotPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceHealthMetadataClient().NewListBySiteSlotPager("Default-Web-NorthCentralUS", "newsiteinnewASE-NCUS", "Production", nil)
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
		// page.ResourceHealthMetadataCollection = armappservice.ResourceHealthMetadataCollection{
		// 	Value: []*armappservice.ResourceHealthMetadata{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Web/sites/resourceHealthMetadata"),
		// 			ID: to.Ptr("/subscriptions/4adb32ad-8327-4cbb-b775-b84b4465bb38/resourceGroups/Default-Web-NorthCentralUS/providers/Microsoft.Web/sites/newsiteinnewASE-NCUS/resourceHealthMetadata/default"),
		// 			Properties: &armappservice.ResourceHealthMetadataProperties{
		// 				Category: to.Ptr("Shared"),
		// 				SignalAvailability: to.Ptr(true),
		// 			},
		// 	}},
		// }
	}
}
