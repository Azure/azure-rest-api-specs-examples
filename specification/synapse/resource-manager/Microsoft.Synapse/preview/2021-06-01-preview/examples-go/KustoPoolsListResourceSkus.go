package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListResourceSkus.json
func ExampleKustoPoolsClient_NewListSKUsByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKustoPoolsClient().NewListSKUsByResourcePager("synapseWorkspaceName", "kustoclusterrptest4", "kustorptest", nil)
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
		// page.ListResourceSKUsResult = armsynapse.ListResourceSKUsResult{
		// 	Value: []*armsynapse.AzureResourceSKU{
		// 		{
		// 			Capacity: &armsynapse.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armsynapse.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("workspaces/kustoPools"),
		// 			SKU: &armsynapse.AzureSKU{
		// 				Name: to.Ptr(armsynapse.SKUNameComputeOptimized),
		// 				Size: to.Ptr(armsynapse.SKUSizeMedium),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armsynapse.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armsynapse.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("workspaces/kustoPools"),
		// 			SKU: &armsynapse.AzureSKU{
		// 				Name: to.Ptr(armsynapse.SKUNameComputeOptimized),
		// 				Size: to.Ptr(armsynapse.SKUSizeLarge),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armsynapse.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armsynapse.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("workspaces/kustoPools"),
		// 			SKU: &armsynapse.AzureSKU{
		// 				Name: to.Ptr(armsynapse.SKUNameStorageOptimized),
		// 				Size: to.Ptr(armsynapse.SKUSizeMedium),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armsynapse.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armsynapse.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("workspaces/kustoPools"),
		// 			SKU: &armsynapse.AzureSKU{
		// 				Name: to.Ptr(armsynapse.SKUNameStorageOptimized),
		// 				Size: to.Ptr(armsynapse.SKUSizeLarge),
		// 			},
		// 	}},
		// }
	}
}
