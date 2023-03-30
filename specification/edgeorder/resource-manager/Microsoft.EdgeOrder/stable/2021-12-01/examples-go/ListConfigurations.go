package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListConfigurations.json
func ExampleManagementClient_NewListConfigurationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewListConfigurationsPager(armedgeorder.ConfigurationsRequest{
		ConfigurationFilters: []*armedgeorder.ConfigurationFilters{
			{
				FilterableProperty: []*armedgeorder.FilterableProperty{
					{
						Type: to.Ptr(armedgeorder.SupportedFilterTypesShipToCountries),
						SupportedValues: []*string{
							to.Ptr("US")},
					}},
				HierarchyInformation: &armedgeorder.HierarchyInformation{
					ProductFamilyName: to.Ptr("azurestackedge"),
					ProductLineName:   to.Ptr("azurestackedge"),
					ProductName:       to.Ptr("azurestackedgegpu"),
				},
			}},
	}, &armedgeorder.ManagementClientListConfigurationsOptions{SkipToken: nil})
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
		// page.Configurations = armedgeorder.Configurations{
		// 	Value: []*armedgeorder.Configuration{
		// 		{
		// 			Properties: &armedgeorder.ConfigurationProperties{
		// 				Description: &armedgeorder.Description{
		// 					Attributes: []*string{
		// 					},
		// 					DescriptionType: to.Ptr(armedgeorder.DescriptionTypeBase),
		// 					Keywords: []*string{
		// 						to.Ptr("GPU")},
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
		// 					DisplayName: to.Ptr("Azure Stack Edge Pro - 1 GPU"),
		// 					HierarchyInformation: &armedgeorder.HierarchyInformation{
		// 						ConfigurationName: to.Ptr("edgep_base"),
		// 						ProductFamilyName: to.Ptr("azurestackedge"),
		// 						ProductLineName: to.Ptr("azurestackedge"),
		// 						ProductName: to.Ptr("azurestackedgegpu"),
		// 					},
		// 					ImageInformation: []*armedgeorder.ImageInformation{
		// 					},
		// 					FilterableProperties: []*armedgeorder.FilterableProperty{
		// 						{
		// 							Type: to.Ptr(armedgeorder.SupportedFilterTypesShipToCountries),
		// 							SupportedValues: []*string{
		// 								to.Ptr("US")},
		// 						}},
		// 						Dimensions: &armedgeorder.Dimensions{
		// 							Depth: to.Ptr[float64](2),
		// 							Height: to.Ptr[float64](15),
		// 							Length: to.Ptr[float64](50),
		// 							LengthHeightUnit: to.Ptr(armedgeorder.LengthHeightUnitIN),
		// 							Weight: to.Ptr[float64](50),
		// 							WeightUnit: to.Ptr(armedgeorder.WeightMeasurementUnitLBS),
		// 							Width: to.Ptr[float64](5),
		// 						},
		// 						Specifications: []*armedgeorder.Specification{
		// 							{
		// 								Name: to.Ptr("Usable compute"),
		// 								Value: to.Ptr("40 vCPU"),
		// 							},
		// 							{
		// 								Name: to.Ptr("Usable memory"),
		// 								Value: to.Ptr("102 GB"),
		// 							},
		// 							{
		// 								Name: to.Ptr("Usable storage"),
		// 								Value: to.Ptr("4.2 TB"),
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
