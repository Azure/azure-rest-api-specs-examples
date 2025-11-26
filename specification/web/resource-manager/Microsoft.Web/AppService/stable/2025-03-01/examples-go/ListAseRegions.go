package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListAseRegions.json
func ExampleWebSiteManagementClient_NewListAseRegionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebSiteManagementClient().NewListAseRegionsPager(nil)
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
		// page.AseRegionCollection = armappservice.AseRegionCollection{
		// 	Value: []*armappservice.AseRegion{
		// 		{
		// 			Type: to.Ptr("Microsoft.Web/aseRegions"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/aseRegions"),
		// 			Properties: &armappservice.AseRegionProperties{
		// 				AvailableOS: []*string{
		// 					to.Ptr("Windows"),
		// 					to.Ptr("Linux"),
		// 					to.Ptr("HyperV")},
		// 					AvailableSKU: []*string{
		// 						to.Ptr("I1v2"),
		// 						to.Ptr("I2v2"),
		// 						to.Ptr("I3v2")},
		// 						DedicatedHost: to.Ptr(true),
		// 						DisplayName: to.Ptr("southcentralus"),
		// 						Standard: to.Ptr(true),
		// 						ZoneRedundant: to.Ptr(true),
		// 					},
		// 				},
		// 				{
		// 					Type: to.Ptr("Microsoft.Web/aseRegions"),
		// 					ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/aseRegions"),
		// 					Properties: &armappservice.AseRegionProperties{
		// 						AvailableOS: []*string{
		// 							to.Ptr("Windows"),
		// 							to.Ptr("Linux")},
		// 							AvailableSKU: []*string{
		// 								to.Ptr("I1v4"),
		// 								to.Ptr("I2v2"),
		// 								to.Ptr("I3v2")},
		// 								DedicatedHost: to.Ptr(true),
		// 								DisplayName: to.Ptr("northcentralus"),
		// 								Standard: to.Ptr(true),
		// 								ZoneRedundant: to.Ptr(true),
		// 							},
		// 					}},
		// 				}
	}
}
