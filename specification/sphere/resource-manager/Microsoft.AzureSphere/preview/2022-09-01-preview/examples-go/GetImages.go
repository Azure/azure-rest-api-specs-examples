package armsphere_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ed9bde6a3db71b84fdba076ba0546213bcce56ee/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetImages.json
func ExampleImagesClient_NewListByCatalogPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewImagesClient().NewListByCatalogPager("MyResourceGroup1", "MyCatalog1", &armsphere.ImagesClientListByCatalogOptions{Filter: nil,
		Top:         nil,
		Skip:        nil,
		Maxpagesize: nil,
	})
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
		// page.ImageListResult = armsphere.ImageListResult{
		// 	Value: []*armsphere.Image{
		// 		{
		// 			Name: to.Ptr("MyImage1"),
		// 			ID: to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 		},
		// 		{
		// 			Name: to.Ptr("MyImage2"),
		// 			ID: to.Ptr("00000000-0000-0000-0000-000000000002"),
		// 	}},
		// }
	}
}
