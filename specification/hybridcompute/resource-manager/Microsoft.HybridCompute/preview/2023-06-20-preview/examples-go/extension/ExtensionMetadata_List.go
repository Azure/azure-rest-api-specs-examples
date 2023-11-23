package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/extension/ExtensionMetadata_List.json
func ExampleExtensionMetadataClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExtensionMetadataClient().NewListPager("EastUS", "microsoft.azure.monitor", "azuremonitorlinuxagent", nil)
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
		// page.ExtensionValueListResult = armhybridcompute.ExtensionValueListResult{
		// 	Value: []*armhybridcompute.ExtensionValue{
		// 		{
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/Providers/Microsoft.HybridCompute/locations/eastus/publishers/microsoft.azure.monitor/extensionTypes/azuremonitorlinuxagent/versions/1.9.1"),
		// 			Properties: &armhybridcompute.ExtensionValueProperties{
		// 				ExtensionType: to.Ptr("azuremonitorlinuxagent"),
		// 				Publisher: to.Ptr("microsoft.azure.monitor"),
		// 				Version: to.Ptr("1.9.1"),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/Providers/Microsoft.HybridCompute/locations/eastus/publishers/microsoft.azure.monitor/extensionTypes/azuremonitorlinuxagent/versions/1.9.2"),
		// 			Properties: &armhybridcompute.ExtensionValueProperties{
		// 				ExtensionType: to.Ptr("azuremonitorlinuxagent"),
		// 				Publisher: to.Ptr("microsoft.azure.monitor"),
		// 				Version: to.Ptr("1.9.2"),
		// 			},
		// 	}},
		// }
	}
}
