package armlabservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4c2cdccf6ca3281dd50ed8788ce1de2e0d480973/specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Skus/listSkus.json
func ExampleSKUsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlabservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListPager(&armlabservices.SKUsClientListOptions{Filter: nil})
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
		// page.PagedSKUInfos = armlabservices.PagedSKUInfos{
		// 	Value: []*armlabservices.SKUInfo{
		// 		{
		// 			Name: to.Ptr("Standard_Fv2_2_4GB_64_S_SSD"),
		// 			Capabilities: []*armlabservices.SKUCapabilities{
		// 				{
		// 					Name: to.Ptr("vCPUs"),
		// 					Value: to.Ptr("2"),
		// 				},
		// 				{
		// 					Name: to.Ptr("MemoryGB"),
		// 					Value: to.Ptr("4"),
		// 				},
		// 				{
		// 					Name: to.Ptr("StorageGB"),
		// 					Value: to.Ptr("64"),
		// 				},
		// 				{
		// 					Name: to.Ptr("StorageType"),
		// 					Value: to.Ptr("StandardSSD"),
		// 				},
		// 				{
		// 					Name: to.Ptr("HyperVGenerations"),
		// 					Value: to.Ptr("V1,V2"),
		// 				},
		// 				{
		// 					Name: to.Ptr("IsGpu"),
		// 					Value: to.Ptr("False"),
		// 			}},
		// 			Capacity: &armlabservices.SKUCapacity{
		// 				Default: to.Ptr[int64](1),
		// 				Maximum: to.Ptr[int64](400),
		// 				Minimum: to.Ptr[int64](0),
		// 				ScaleType: to.Ptr(armlabservices.ScaleTypeAutomatic),
		// 			},
		// 			Costs: []*armlabservices.SKUCost{
		// 				{
		// 					ExtendedUnit: to.Ptr("WindowsHourly"),
		// 					MeterID: to.Ptr(""),
		// 					Quantity: to.Ptr[float32](20.4),
		// 				},
		// 				{
		// 					ExtendedUnit: to.Ptr("HybridBenefitHourly"),
		// 					MeterID: to.Ptr(""),
		// 					Quantity: to.Ptr[float32](20.4),
		// 				},
		// 				{
		// 					ExtendedUnit: to.Ptr("InactiveHourly"),
		// 					MeterID: to.Ptr(""),
		// 					Quantity: to.Ptr[float32](0.7),
		// 				},
		// 				{
		// 					ExtendedUnit: to.Ptr("LinuxHourly"),
		// 					MeterID: to.Ptr(""),
		// 					Quantity: to.Ptr[float32](20.4),
		// 			}},
		// 			Family: to.Ptr("Fv2"),
		// 			Locations: []*string{
		// 				to.Ptr("eastus2")},
		// 				ResourceType: to.Ptr("labs"),
		// 				Size: to.Ptr("Fv2_2_4GB_64_S_SSD"),
		// 				Tier: to.Ptr(armlabservices.LabServicesSKUTierStandard),
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard_Fv2_2_4GB_256_S_SSD"),
		// 				Capabilities: []*armlabservices.SKUCapabilities{
		// 					{
		// 						Name: to.Ptr("vCPUs"),
		// 						Value: to.Ptr("2"),
		// 					},
		// 					{
		// 						Name: to.Ptr("MemoryGB"),
		// 						Value: to.Ptr("4"),
		// 					},
		// 					{
		// 						Name: to.Ptr("StorageGB"),
		// 						Value: to.Ptr("256"),
		// 					},
		// 					{
		// 						Name: to.Ptr("StorageType"),
		// 						Value: to.Ptr("StandardSSD"),
		// 					},
		// 					{
		// 						Name: to.Ptr("HyperVGenerations"),
		// 						Value: to.Ptr("V1,V2"),
		// 					},
		// 					{
		// 						Name: to.Ptr("IsGpu"),
		// 						Value: to.Ptr("False"),
		// 				}},
		// 				Capacity: &armlabservices.SKUCapacity{
		// 					Default: to.Ptr[int64](1),
		// 					Maximum: to.Ptr[int64](400),
		// 					Minimum: to.Ptr[int64](0),
		// 					ScaleType: to.Ptr(armlabservices.ScaleTypeAutomatic),
		// 				},
		// 				Costs: []*armlabservices.SKUCost{
		// 					{
		// 						ExtendedUnit: to.Ptr("WindowsHourly"),
		// 						MeterID: to.Ptr(""),
		// 						Quantity: to.Ptr[float32](22.8),
		// 					},
		// 					{
		// 						ExtendedUnit: to.Ptr("HybridBenefitHourly"),
		// 						MeterID: to.Ptr(""),
		// 						Quantity: to.Ptr[float32](22.8),
		// 					},
		// 					{
		// 						ExtendedUnit: to.Ptr("InactiveHourly"),
		// 						MeterID: to.Ptr(""),
		// 						Quantity: to.Ptr[float32](2.7),
		// 					},
		// 					{
		// 						ExtendedUnit: to.Ptr("LinuxHourly"),
		// 						MeterID: to.Ptr(""),
		// 						Quantity: to.Ptr[float32](22.8),
		// 				}},
		// 				Family: to.Ptr("Fv2"),
		// 				Locations: []*string{
		// 					to.Ptr("eastus2")},
		// 					ResourceType: to.Ptr("labs"),
		// 					Size: to.Ptr("Fv2_2_4GB_256_S_SSD"),
		// 					Tier: to.Ptr(armlabservices.LabServicesSKUTierStandard),
		// 				},
		// 				{
		// 					Name: to.Ptr("Standard_Dv4_2_8GB_128_S_SSD"),
		// 					Capabilities: []*armlabservices.SKUCapabilities{
		// 						{
		// 							Name: to.Ptr("vCPUs"),
		// 							Value: to.Ptr("2"),
		// 						},
		// 						{
		// 							Name: to.Ptr("MemoryGB"),
		// 							Value: to.Ptr("8"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageGB"),
		// 							Value: to.Ptr("128"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageType"),
		// 							Value: to.Ptr("StandardSSD"),
		// 						},
		// 						{
		// 							Name: to.Ptr("HyperVGenerations"),
		// 							Value: to.Ptr("V1,V2"),
		// 						},
		// 						{
		// 							Name: to.Ptr("IsGpu"),
		// 							Value: to.Ptr("False"),
		// 					}},
		// 					Capacity: &armlabservices.SKUCapacity{
		// 						Default: to.Ptr[int64](1),
		// 						Maximum: to.Ptr[int64](400),
		// 						Minimum: to.Ptr[int64](0),
		// 						ScaleType: to.Ptr(armlabservices.ScaleTypeAutomatic),
		// 					},
		// 					Costs: []*armlabservices.SKUCost{
		// 						{
		// 							ExtendedUnit: to.Ptr("WindowsHourly"),
		// 							MeterID: to.Ptr(""),
		// 							Quantity: to.Ptr[float32](24.2),
		// 						},
		// 						{
		// 							ExtendedUnit: to.Ptr("HybridBenefitHourly"),
		// 							MeterID: to.Ptr(""),
		// 							Quantity: to.Ptr[float32](24.2),
		// 						},
		// 						{
		// 							ExtendedUnit: to.Ptr("InactiveHourly"),
		// 							MeterID: to.Ptr(""),
		// 							Quantity: to.Ptr[float32](1.4),
		// 						},
		// 						{
		// 							ExtendedUnit: to.Ptr("LinuxHourly"),
		// 							MeterID: to.Ptr(""),
		// 							Quantity: to.Ptr[float32](24.2),
		// 					}},
		// 					Family: to.Ptr("Dv4"),
		// 					Locations: []*string{
		// 						to.Ptr("eastus2")},
		// 						ResourceType: to.Ptr("labs"),
		// 						Size: to.Ptr("Dv4_2_8GB_128_S_SSD"),
		// 						Tier: to.Ptr(armlabservices.LabServicesSKUTierStandard),
		// 					},
		// 					{
		// 						Name: to.Ptr("Standard_Dv4_2_8GB_256_S_SSD"),
		// 						Capabilities: []*armlabservices.SKUCapabilities{
		// 							{
		// 								Name: to.Ptr("vCPUs"),
		// 								Value: to.Ptr("2"),
		// 							},
		// 							{
		// 								Name: to.Ptr("MemoryGB"),
		// 								Value: to.Ptr("8"),
		// 							},
		// 							{
		// 								Name: to.Ptr("StorageGB"),
		// 								Value: to.Ptr("256"),
		// 							},
		// 							{
		// 								Name: to.Ptr("StorageType"),
		// 								Value: to.Ptr("StandardSSD"),
		// 							},
		// 							{
		// 								Name: to.Ptr("HyperVGenerations"),
		// 								Value: to.Ptr("V1,V2"),
		// 							},
		// 							{
		// 								Name: to.Ptr("IsGpu"),
		// 								Value: to.Ptr("False"),
		// 						}},
		// 						Capacity: &armlabservices.SKUCapacity{
		// 							Default: to.Ptr[int64](1),
		// 							Maximum: to.Ptr[int64](400),
		// 							Minimum: to.Ptr[int64](0),
		// 							ScaleType: to.Ptr(armlabservices.ScaleTypeAutomatic),
		// 						},
		// 						Costs: []*armlabservices.SKUCost{
		// 							{
		// 								ExtendedUnit: to.Ptr("WindowsHourly"),
		// 								MeterID: to.Ptr(""),
		// 								Quantity: to.Ptr[float32](25.8),
		// 							},
		// 							{
		// 								ExtendedUnit: to.Ptr("HybridBenefitHourly"),
		// 								MeterID: to.Ptr(""),
		// 								Quantity: to.Ptr[float32](25.8),
		// 							},
		// 							{
		// 								ExtendedUnit: to.Ptr("InactiveHourly"),
		// 								MeterID: to.Ptr(""),
		// 								Quantity: to.Ptr[float32](2.7),
		// 							},
		// 							{
		// 								ExtendedUnit: to.Ptr("LinuxHourly"),
		// 								MeterID: to.Ptr(""),
		// 								Quantity: to.Ptr[float32](25.8),
		// 						}},
		// 						Family: to.Ptr("Dv4"),
		// 						Locations: []*string{
		// 							to.Ptr("eastus2")},
		// 							ResourceType: to.Ptr("labs"),
		// 							Size: to.Ptr("Dv4_2_8GB_256_S_SSD"),
		// 							Tier: to.Ptr(armlabservices.LabServicesSKUTierStandard),
		// 					}},
		// 				}
	}
}
