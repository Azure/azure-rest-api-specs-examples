package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoClustersListResourceSkus.json
func ExampleClustersClient_NewListSKUsByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkusto.NewClustersClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListSKUsByResourcePager("kustorptest", "kustoCluster", nil)
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
		// page.ListResourceSKUsResult = armkusto.ListResourceSKUsResult{
		// 	Value: []*armkusto.AzureResourceSKU{
		// 		{
		// 			Capacity: &armkusto.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armkusto.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("clusters"),
		// 			SKU: &armkusto.AzureSKU{
		// 				Name: to.Ptr(armkusto.AzureSKUNameStandardL8SV3),
		// 				Tier: to.Ptr(armkusto.AzureSKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armkusto.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armkusto.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("clusters"),
		// 			SKU: &armkusto.AzureSKU{
		// 				Name: to.Ptr(armkusto.AzureSKUNameStandardL8AsV3),
		// 				Tier: to.Ptr(armkusto.AzureSKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armkusto.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armkusto.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("clusters"),
		// 			SKU: &armkusto.AzureSKU{
		// 				Name: to.Ptr(armkusto.AzureSKUNameStandardL16AsV3),
		// 				Tier: to.Ptr(armkusto.AzureSKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armkusto.AzureCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](100),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr(armkusto.AzureScaleTypeAutomatic),
		// 			},
		// 			ResourceType: to.Ptr("clusters"),
		// 			SKU: &armkusto.AzureSKU{
		// 				Name: to.Ptr(armkusto.AzureSKUNameStandardL16SV3),
		// 				Tier: to.Ptr(armkusto.AzureSKUTierStandard),
		// 			},
		// 	}},
		// }
	}
}
