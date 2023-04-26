package armcontainerinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e60df62e9e0d88462e6abba81a76d94eab000f0d/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2023-05-01/examples/CachedImagesList.json
func ExampleLocationClient_NewListCachedImagesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationClient().NewListCachedImagesPager("westcentralus", nil)
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
		// page.CachedImagesListResult = armcontainerinstance.CachedImagesListResult{
		// 	Value: []*armcontainerinstance.CachedImages{
		// 		{
		// 			Image: to.Ptr("ubuntu:16.04"),
		// 			OSType: to.Ptr("Linux"),
		// 		},
		// 		{
		// 			Image: to.Ptr("alpine:3.6"),
		// 			OSType: to.Ptr("Linux"),
		// 		},
		// 		{
		// 			Image: to.Ptr("microsoft/nanoserver:10.0.14393.2485"),
		// 			OSType: to.Ptr("Windows"),
		// 	}},
		// }
	}
}
