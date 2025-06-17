package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/extension/ExtensionMetadata_Get.json
func ExampleExtensionMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtensionMetadataClient().Get(ctx, "EastUS", "microsoft.azure.monitor", "azuremonitorlinuxagent", "1.9.1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtensionValue = armhybridcompute.ExtensionValue{
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/Providers/Microsoft.HybridCompute/locations/eastus/publishers/microsoft.azure.monitor/extensionTypes/azuremonitorlinuxagent/versions/1.9.1"),
	// 	Properties: &armhybridcompute.ExtensionValueProperties{
	// 		ExtensionType: to.Ptr("azuremonitorlinuxagent"),
	// 		Publisher: to.Ptr("microsoft.azure.monitor"),
	// 		Version: to.Ptr("1.9.1"),
	// 	},
	// }
}
