package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/AvailableSkusPost.json
func ExampleServiceClient_NewListAvailableSKUsByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceClient().NewListAvailableSKUsByResourceGroupPager("YourResourceGroupName", "westus", armdatabox.AvailableSKURequest{
		Country:      to.Ptr("XX"),
		Location:     to.Ptr("westus"),
		TransferType: to.Ptr(armdatabox.TransferTypeImportToAzure),
	}, nil)
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
		// page.AvailableSKUsResult = armdatabox.AvailableSKUsResult{
		// 	Value: []*armdatabox.SKUInformation{
		// 		{
		// 			Enabled: to.Ptr(true),
		// 			Properties: &armdatabox.SKUProperties{
		// 				APIVersions: []*string{
		// 					to.Ptr("2018-01-01")},
		// 					Capacity: &armdatabox.SKUCapacity{
		// 						Maximum: to.Ptr("100"),
		// 						Usable: to.Ptr("80"),
		// 					},
		// 					Costs: []*armdatabox.SKUCost{
		// 						{
		// 							MeterID: to.Ptr("0cf23ffc-0b64-49e6-9bdd-1db885349042"),
		// 							MeterType: to.Ptr("DataBoxServiceFee"),
		// 							Multiplier: to.Ptr[float64](1),
		// 						},
		// 						{
		// 							MeterID: to.Ptr("a701f058-119b-4713-a923-bed7da4b7801"),
		// 							MeterType: to.Ptr("DataBoxShippingFee"),
		// 							Multiplier: to.Ptr[float64](1),
		// 						},
		// 						{
		// 							MeterID: to.Ptr("049fa331-0a48-4a81-9991-a6bef3c79fed"),
		// 							MeterType: to.Ptr("DataBoxExtraDayFee"),
		// 							Multiplier: to.Ptr[float64](1),
		// 						},
		// 						{
		// 							MeterID: to.Ptr("c3ea978d-6a0a-4632-b094-4fffcafcb057"),
		// 							MeterType: to.Ptr("DataBoxLostOrDamagedDeviceFee"),
		// 							Multiplier: to.Ptr[float64](1),
		// 					}},
		// 					CountriesWithinCommerceBoundary: []*string{
		// 						to.Ptr("XX")},
		// 						DataLocationToServiceLocationMap: []*armdatabox.DataLocationToServiceLocationMap{
		// 							{
		// 								DataLocation: to.Ptr("westus"),
		// 								ServiceLocation: to.Ptr("westus"),
		// 							},
		// 							{
		// 								DataLocation: to.Ptr("centralus"),
		// 								ServiceLocation: to.Ptr("westus"),
		// 							},
		// 							{
		// 								DataLocation: to.Ptr("eastus"),
		// 								ServiceLocation: to.Ptr("westus"),
		// 							},
		// 							{
		// 								DataLocation: to.Ptr("eastus2"),
		// 								ServiceLocation: to.Ptr("westus"),
		// 							},
		// 							{
		// 								DataLocation: to.Ptr("northcentralus"),
		// 								ServiceLocation: to.Ptr("westus"),
		// 							},
		// 							{
		// 								DataLocation: to.Ptr("southcentralus"),
		// 								ServiceLocation: to.Ptr("westus"),
		// 							},
		// 							{
		// 								DataLocation: to.Ptr("westcentralus"),
		// 								ServiceLocation: to.Ptr("westus"),
		// 							},
		// 							{
		// 								DataLocation: to.Ptr("westus2"),
		// 								ServiceLocation: to.Ptr("westus"),
		// 							},
		// 							{
		// 								DataLocation: to.Ptr("centraluseuap"),
		// 								ServiceLocation: to.Ptr("centraluseuap"),
		// 							},
		// 							{
		// 								DataLocation: to.Ptr("eastus2euap"),
		// 								ServiceLocation: to.Ptr("eastus2euap"),
		// 						}},
		// 						DisabledReason: to.Ptr(armdatabox.SKUDisabledReasonNone),
		// 					},
		// 					SKU: &armdatabox.SKU{
		// 						Name: to.Ptr(armdatabox.SKUNameDataBox),
		// 					},
		// 				},
		// 				{
		// 					Enabled: to.Ptr(true),
		// 					Properties: &armdatabox.SKUProperties{
		// 						APIVersions: []*string{
		// 							to.Ptr("2018-01-01")},
		// 							Capacity: &armdatabox.SKUCapacity{
		// 								Maximum: to.Ptr("40"),
		// 								Usable: to.Ptr("35"),
		// 							},
		// 							Costs: []*armdatabox.SKUCost{
		// 								{
		// 									MeterID: to.Ptr("d95cd8b5-b6f1-4cd9-ae86-a016d1945d6f"),
		// 									MeterType: to.Ptr("DataBoxDiskServiceFee"),
		// 									Multiplier: to.Ptr[float64](1),
		// 								},
		// 								{
		// 									MeterID: to.Ptr("4b8cf572-cb04-4ef3-9528-2cda4e9b544e"),
		// 									MeterType: to.Ptr("DataBoxDiskShippingFee"),
		// 									Multiplier: to.Ptr[float64](1),
		// 								},
		// 								{
		// 									MeterID: to.Ptr("b6ae9bbf-815d-49dd-bb2b-77c497b72ba4"),
		// 									MeterType: to.Ptr("DataBoxDiskDailyUsageFee"),
		// 									Multiplier: to.Ptr[float64](1),
		// 								},
		// 								{
		// 									MeterID: to.Ptr("08bc0ea1-6c82-421b-b953-2a7a65810d2e"),
		// 									MeterType: to.Ptr("DataBoxDiskLostDeviceFee"),
		// 									Multiplier: to.Ptr[float64](1),
		// 							}},
		// 							CountriesWithinCommerceBoundary: []*string{
		// 								to.Ptr("XX")},
		// 								DataLocationToServiceLocationMap: []*armdatabox.DataLocationToServiceLocationMap{
		// 									{
		// 										DataLocation: to.Ptr("westus"),
		// 										ServiceLocation: to.Ptr("westus"),
		// 									},
		// 									{
		// 										DataLocation: to.Ptr("centralus"),
		// 										ServiceLocation: to.Ptr("westus"),
		// 									},
		// 									{
		// 										DataLocation: to.Ptr("eastus"),
		// 										ServiceLocation: to.Ptr("westus"),
		// 									},
		// 									{
		// 										DataLocation: to.Ptr("eastus2"),
		// 										ServiceLocation: to.Ptr("westus"),
		// 									},
		// 									{
		// 										DataLocation: to.Ptr("northcentralus"),
		// 										ServiceLocation: to.Ptr("westus"),
		// 									},
		// 									{
		// 										DataLocation: to.Ptr("southcentralus"),
		// 										ServiceLocation: to.Ptr("westus"),
		// 									},
		// 									{
		// 										DataLocation: to.Ptr("westcentralus"),
		// 										ServiceLocation: to.Ptr("westus"),
		// 									},
		// 									{
		// 										DataLocation: to.Ptr("westus2"),
		// 										ServiceLocation: to.Ptr("westus"),
		// 									},
		// 									{
		// 										DataLocation: to.Ptr("centraluseuap"),
		// 										ServiceLocation: to.Ptr("centraluseuap"),
		// 									},
		// 									{
		// 										DataLocation: to.Ptr("eastus2euap"),
		// 										ServiceLocation: to.Ptr("eastus2euap"),
		// 								}},
		// 								DisabledReason: to.Ptr(armdatabox.SKUDisabledReasonNone),
		// 							},
		// 							SKU: &armdatabox.SKU{
		// 								Name: to.Ptr(armdatabox.SKUNameDataBoxDisk),
		// 							},
		// 						},
		// 						{
		// 							Enabled: to.Ptr(true),
		// 							Properties: &armdatabox.SKUProperties{
		// 								APIVersions: []*string{
		// 									to.Ptr("2018-01-01")},
		// 									Capacity: &armdatabox.SKUCapacity{
		// 										Maximum: to.Ptr("1000"),
		// 										Usable: to.Ptr("800"),
		// 									},
		// 									Costs: []*armdatabox.SKUCost{
		// 										{
		// 											MeterID: to.Ptr("d0dccaaf-3de9-4c7a-ba97-f83551b90126"),
		// 											MeterType: to.Ptr("DataBoxHeavyServiceFee"),
		// 											Multiplier: to.Ptr[float64](1),
		// 										},
		// 										{
		// 											MeterID: to.Ptr("7b49d11f-d4f7-4029-a197-04998fd282f9"),
		// 											MeterType: to.Ptr("DataBoxHeavyShippingFee"),
		// 											Multiplier: to.Ptr[float64](1),
		// 										},
		// 										{
		// 											MeterID: to.Ptr("c2c66d53-11b4-4f11-9642-43c7c336f0b7"),
		// 											MeterType: to.Ptr("DataBoxHeavyExtraDayFee"),
		// 											Multiplier: to.Ptr[float64](1),
		// 										},
		// 										{
		// 											MeterID: to.Ptr("188dcd7e-fbd7-4a41-aa42-162b81b0510f"),
		// 											MeterType: to.Ptr("DataBoxHeavyLostOrDamagedDeviceFee"),
		// 											Multiplier: to.Ptr[float64](1),
		// 									}},
		// 									CountriesWithinCommerceBoundary: []*string{
		// 										to.Ptr("XX")},
		// 										DataLocationToServiceLocationMap: []*armdatabox.DataLocationToServiceLocationMap{
		// 											{
		// 												DataLocation: to.Ptr("westus"),
		// 												ServiceLocation: to.Ptr("westus"),
		// 											},
		// 											{
		// 												DataLocation: to.Ptr("centralus"),
		// 												ServiceLocation: to.Ptr("westus"),
		// 											},
		// 											{
		// 												DataLocation: to.Ptr("eastus"),
		// 												ServiceLocation: to.Ptr("westus"),
		// 											},
		// 											{
		// 												DataLocation: to.Ptr("eastus2"),
		// 												ServiceLocation: to.Ptr("westus"),
		// 											},
		// 											{
		// 												DataLocation: to.Ptr("northcentralus"),
		// 												ServiceLocation: to.Ptr("westus"),
		// 											},
		// 											{
		// 												DataLocation: to.Ptr("southcentralus"),
		// 												ServiceLocation: to.Ptr("westus"),
		// 											},
		// 											{
		// 												DataLocation: to.Ptr("westcentralus"),
		// 												ServiceLocation: to.Ptr("westus"),
		// 											},
		// 											{
		// 												DataLocation: to.Ptr("westus2"),
		// 												ServiceLocation: to.Ptr("westus"),
		// 											},
		// 											{
		// 												DataLocation: to.Ptr("centraluseuap"),
		// 												ServiceLocation: to.Ptr("centraluseuap"),
		// 											},
		// 											{
		// 												DataLocation: to.Ptr("eastus2euap"),
		// 												ServiceLocation: to.Ptr("eastus2euap"),
		// 										}},
		// 										DisabledReason: to.Ptr(armdatabox.SKUDisabledReasonNone),
		// 										RequiredFeature: to.Ptr("HeavyCreateAccess"),
		// 									},
		// 									SKU: &armdatabox.SKU{
		// 										Name: to.Ptr(armdatabox.SKUNameDataBoxHeavy),
		// 									},
		// 							}},
		// 						}
	}
}
