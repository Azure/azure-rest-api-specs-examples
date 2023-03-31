package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListProductFamiliesMetadata.json
func ExampleManagementClient_NewListProductFamiliesMetadataPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewListProductFamiliesMetadataPager(&armedgeorder.ManagementClientListProductFamiliesMetadataOptions{SkipToken: nil})
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
		// page.ProductFamiliesMetadata = armedgeorder.ProductFamiliesMetadata{
		// 	Value: []*armedgeorder.ProductFamiliesMetadataDetails{
		// 		{
		// 			Properties: &armedgeorder.ProductFamilyProperties{
		// 				Description: &armedgeorder.Description{
		// 					Attributes: []*string{
		// 					},
		// 					DescriptionType: to.Ptr(armedgeorder.DescriptionTypeBase),
		// 					Keywords: []*string{
		// 					},
		// 					Links: []*armedgeorder.Link{
		// 					},
		// 					ShortDescription: to.Ptr("Azure managed physical edge compute device"),
		// 				},
		// 				AvailabilityInformation: &armedgeorder.AvailabilityInformation{
		// 					AvailabilityStage: to.Ptr(armedgeorder.AvailabilityStageAvailable),
		// 					DisabledReason: to.Ptr(armedgeorder.DisabledReasonNone),
		// 				},
		// 				DisplayName: to.Ptr("Azure Stack Edge"),
		// 				HierarchyInformation: &armedgeorder.HierarchyInformation{
		// 					ConfigurationName: to.Ptr(""),
		// 					ProductFamilyName: to.Ptr("azurestackedge"),
		// 					ProductLineName: to.Ptr(""),
		// 					ProductName: to.Ptr(""),
		// 				},
		// 				ImageInformation: []*armedgeorder.ImageInformation{
		// 				},
		// 				FilterableProperties: []*armedgeorder.FilterableProperty{
		// 					{
		// 						Type: to.Ptr(armedgeorder.SupportedFilterTypesShipToCountries),
		// 						SupportedValues: []*string{
		// 							to.Ptr("US")},
		// 						},
		// 						{
		// 							Type: to.Ptr(armedgeorder.SupportedFilterTypesDoubleEncryptionStatus),
		// 							SupportedValues: []*string{
		// 								to.Ptr("Enabled")},
		// 						}},
		// 						ProductLines: []*armedgeorder.ProductLine{
		// 						},
		// 						ResourceProviderDetails: []*armedgeorder.ResourceProviderDetails{
		// 							{
		// 								ResourceProviderNamespace: to.Ptr("Microsoft.DataBoxEdge"),
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
