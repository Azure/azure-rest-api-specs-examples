package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder/v2"
)

// Generated from example definition: 2024-02-01/ListConfigurations.json
func ExampleProductsAndConfigurationsClient_NewListConfigurationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("eb5dc900-6186-49d8-b7d7-febd866fdc1d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProductsAndConfigurationsClient().NewListConfigurationsPager(armedgeorder.ConfigurationsRequest{
		ConfigurationFilter: &armedgeorder.ConfigurationFilter{
			FilterableProperty: []*armedgeorder.FilterableProperty{
				{
					Type: to.Ptr(armedgeorder.SupportedFilterTypesShipToCountries),
					SupportedValues: []*string{
						to.Ptr("US"),
					},
				},
			},
			HierarchyInformation: &armedgeorder.HierarchyInformation{
				ProductFamilyName: to.Ptr("azurestackedge"),
				ProductLineName:   to.Ptr("azurestackedge"),
				ProductName:       to.Ptr("azurestackedgegpu"),
			},
		},
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
		// page = armedgeorder.ProductsAndConfigurationsClientListConfigurationsResponse{
		// 	Configurations: armedgeorder.Configurations{
		// 		Value: []*armedgeorder.Configuration{
		// 			{
		// 				Properties: &armedgeorder.ConfigurationProperties{
		// 					Description: &armedgeorder.Description{
		// 						Attributes: []*string{
		// 						},
		// 						DescriptionType: to.Ptr(armedgeorder.DescriptionTypeBase),
		// 						Keywords: []*string{
		// 							to.Ptr("GPU"),
		// 						},
		// 						Links: []*armedgeorder.Link{
		// 						},
		// 						LongDescription: to.Ptr(""),
		// 						ShortDescription: to.Ptr(""),
		// 					},
		// 					AvailabilityInformation: &armedgeorder.AvailabilityInformation{
		// 						AvailabilityStage: to.Ptr(armedgeorder.AvailabilityStageAvailable),
		// 						DisabledReason: to.Ptr(armedgeorder.DisabledReasonNone),
		// 					},
		// 					CostInformation: &armedgeorder.CostInformation{
		// 						BillingMeterDetails: []*armedgeorder.BillingMeterDetails{
		// 						},
		// 					},
		// 					Dimensions: &armedgeorder.Dimensions{
		// 						Depth: to.Ptr[float64](2),
		// 						Height: to.Ptr[float64](15),
		// 						Length: to.Ptr[float64](50),
		// 						LengthHeightUnit: to.Ptr(armedgeorder.LengthHeightUnitIN),
		// 						Weight: to.Ptr[float64](50),
		// 						WeightUnit: to.Ptr(armedgeorder.WeightMeasurementUnitLBS),
		// 						Width: to.Ptr[float64](5),
		// 					},
		// 					DisplayName: to.Ptr("Azure Stack Edge Pro - 1 GPU"),
		// 					FilterableProperties: []*armedgeorder.FilterableProperty{
		// 						{
		// 							Type: to.Ptr(armedgeorder.SupportedFilterTypesShipToCountries),
		// 							SupportedValues: []*string{
		// 								to.Ptr("US"),
		// 							},
		// 						},
		// 					},
		// 					HierarchyInformation: &armedgeorder.HierarchyInformation{
		// 						ConfigurationName: to.Ptr("edgep_base"),
		// 						ProductFamilyName: to.Ptr("azurestackedge"),
		// 						ProductLineName: to.Ptr("azurestackedge"),
		// 						ProductName: to.Ptr("azurestackedgegpu"),
		// 					},
		// 					ImageInformation: []*armedgeorder.ImageInformation{
		// 					},
		// 					Specifications: []*armedgeorder.Specification{
		// 						{
		// 							Name: to.Ptr("Usable compute"),
		// 							Value: to.Ptr("40 vCPU"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Usable memory"),
		// 							Value: to.Ptr("102 GB"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Usable storage"),
		// 							Value: to.Ptr("4.2 TB"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
