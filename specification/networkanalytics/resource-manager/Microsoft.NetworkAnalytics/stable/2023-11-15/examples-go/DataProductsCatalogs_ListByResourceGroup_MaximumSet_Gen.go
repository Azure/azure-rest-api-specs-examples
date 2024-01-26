package armnetworkanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21a8d55d74e4425e96d76e5835f52cfc9eb95a22/specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProductsCatalogs_ListByResourceGroup_MaximumSet_Gen.json
func ExampleDataProductsCatalogsClient_NewListByResourceGroupPager_dataProductsCatalogsListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataProductsCatalogsClient().NewListByResourceGroupPager("aoiresourceGroupName", nil)
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
		// page.DataProductsCatalogListResult = armnetworkanalytics.DataProductsCatalogListResult{
		// 	Value: []*armnetworkanalytics.DataProductsCatalog{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.NetworkAnalytics/dataProductsCatalogs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/aoiresourceGroupName/providers/Microsoft.NetworkAnalytics/dataProductsCatalogs/default"),
		// 			SystemData: &armnetworkanalytics.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-04T08:26:27.150Z"); return t}()),
		// 				CreatedBy: to.Ptr("abc@micros.com"),
		// 				CreatedByType: to.Ptr(armnetworkanalytics.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-04T08:26:27.150Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("abc@micros.com"),
		// 				LastModifiedByType: to.Ptr(armnetworkanalytics.CreatedByTypeUser),
		// 			},
		// 			Properties: &armnetworkanalytics.DataProductsCatalogProperties{
		// 				ProvisioningState: to.Ptr(armnetworkanalytics.ProvisioningStateSucceeded),
		// 				Publishers: []*armnetworkanalytics.PublisherInformation{
		// 					{
		// 						DataProducts: []*armnetworkanalytics.DataProductInformation{
		// 							{
		// 								Description: to.Ptr("Official data product for Mobile Content Cloud."),
		// 								DataProductName: to.Ptr("MCC"),
		// 								DataProductVersions: []*armnetworkanalytics.DataProductVersion{
		// 									{
		// 										Version: to.Ptr("1.0.0"),
		// 								}},
		// 						}},
		// 						PublisherName: to.Ptr("Microsoft"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
