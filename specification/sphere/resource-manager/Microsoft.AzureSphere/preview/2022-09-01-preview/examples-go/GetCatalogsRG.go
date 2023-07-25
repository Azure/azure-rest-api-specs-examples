package armsphere_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ed9bde6a3db71b84fdba076ba0546213bcce56ee/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetCatalogsRG.json
func ExampleCatalogsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogsClient().NewListByResourceGroupPager("MyResourceGroup1", nil)
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
		// page.CatalogListResult = armsphere.CatalogListResult{
		// 	Value: []*armsphere.Catalog{
		// 		{
		// 			Name: to.Ptr("MyCatalog1"),
		// 			Type: to.Ptr("Microsoft.AzureSphere/catalogs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1"),
		// 			Location: to.Ptr("global"),
		// 		},
		// 		{
		// 			Name: to.Ptr("MyCatalog2"),
		// 			Type: to.Ptr("Microsoft.AzureSphere/catalogs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog2"),
		// 			Location: to.Ptr("global"),
		// 	}},
		// }
	}
}
