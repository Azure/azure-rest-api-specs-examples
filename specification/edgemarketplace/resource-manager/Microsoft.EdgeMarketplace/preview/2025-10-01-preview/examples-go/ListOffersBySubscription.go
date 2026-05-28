package armedgemarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgemarketplace/armedgemarketplace"
)

// Generated from example definition: 2025-10-01-preview/ListOffersBySubscription.json
func ExampleOffersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgemarketplace.NewClientFactory("4bed37fd-19a1-4d31-8b44-40267555bec5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOffersClient().NewListBySubscriptionPager(nil)
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
		// page = armedgemarketplace.OffersClientListBySubscriptionResponse{
		// 	OfferListResult: armedgemarketplace.OfferListResult{
		// 		Value: []*armedgemarketplace.Offer{
		// 			{
		// 				Properties: &armedgemarketplace.OfferProperties{
		// 					ContentVersion: to.Ptr("1.0"),
		// 					ContentURL: to.Ptr("test"),
		// 					ProvisioningState: to.Ptr(armedgemarketplace.ResourceProvisioningStateSucceeded),
		// 					OfferContent: &armedgemarketplace.OfferContent{
		// 						DisplayName: to.Ptr("Ubuntu Pro 22.04 LTS"),
		// 						Summary: to.Ptr("Ubuntu Pro is providing additional coverage for production environments running in the cloud."),
		// 						LongSummary: to.Ptr("The official Ubuntu Linux, optimized for Azure with ten years of maintenance and additional security, compliance (e.g. FIPS, CIS, DISA) and management tools."),
		// 						Description: to.Ptr("<p>Ubuntu Pro is a cross-cloud OS optimized for Azure"),
		// 						OfferID: to.Ptr("0001-com-ubuntu-pro-jammy"),
		// 						Popularity: to.Ptr[int32](8),
		// 						OfferPublisher: &armedgemarketplace.OfferPublisher{
		// 							PublisherID: to.Ptr("canonical"),
		// 							PublisherDisplayName: to.Ptr("Canonical"),
		// 						},
		// 						OfferType: to.Ptr("VirtualMachine"),
		// 						Availability: to.Ptr(armedgemarketplace.OfferAvailability("Preview")),
		// 						ReleaseType: to.Ptr(armedgemarketplace.OfferReleaseTypePreview),
		// 						TermsAndConditions: &armedgemarketplace.TermsAndConditions{
		// 							LegalTermsURI: to.Ptr("https://query.prod.cms.rt.microsoft.com/cms/api/am/binary/RW14H4N"),
		// 							LegalTermsType: to.Ptr("None"),
		// 							PrivacyPolicyURI: to.Ptr("http://www.ubuntu.com/aboutus/privacypolicy"),
		// 						},
		// 						SupportURI: to.Ptr("https://ubuntu.com/azure/support"),
		// 						CategoryIDs: []*string{
		// 						},
		// 						IconFileUris: &armedgemarketplace.IconFileUris{
		// 							Small: to.Ptr("https://store-images.s-microsoft.com/image/apps.32072.2b66b35a-b926-4ddc-85fa-cc9d9cb874d9.77c42e2a-c945-4abb-8faa-365a38a108ed.42bcf340-fd72-405a-844d-7b0331419ec1"),
		// 							Medium: to.Ptr("https://store-images.s-microsoft.com/image/apps.58881.2b66b35a-b926-4ddc-85fa-cc9d9cb874d9.77c42e2a-c945-4abb-8faa-365a38a108ed.31dd70ea-0671-4ad3-992b-6ce1a3f7a49e"),
		// 							Wide: to.Ptr("https://store-images.s-microsoft.com/image/apps.17366.2b66b35a-b926-4ddc-85fa-cc9d9cb874d9.77c42e2a-c945-4abb-8faa-365a38a108ed.a2ff8d5b-de36-4d9c-be47-81e15b0c6f6a"),
		// 							Large: to.Ptr("https://store-images.s-microsoft.com/image/apps.60707.2b66b35a-b926-4ddc-85fa-cc9d9cb874d9.77c42e2a-c945-4abb-8faa-365a38a108ed.50b643a0-de63-4ac1-88a2-d2390e2123c2"),
		// 						},
		// 					},
		// 					MarketplaceSKUs: []*armedgemarketplace.MarketplaceSKU{
		// 						{
		// 							CatalogPlanID: to.Ptr("canonical.0001-com-ubuntu-pro-jammypro-22_04-lts"),
		// 							MarketplaceSKUID: to.Ptr("pro-22_04-lts"),
		// 							DisplayName: to.Ptr("pro-22_04-lts"),
		// 							Type: to.Ptr("test"),
		// 							OperatingSystem: &armedgemarketplace.SKUOperatingSystem{
		// 								Name: to.Ptr("linux"),
		// 								Type: to.Ptr("test"),
		// 								Family: to.Ptr("test"),
		// 							},
		// 							MarketplaceSKUVersions: []*armedgemarketplace.MarketplaceSKUVersion{
		// 								{
		// 									Name: to.Ptr("22.04.202204200"),
		// 									SizeOnDiskInMb: to.Ptr[int32](0),
		// 									MinimumDownloadSizeInMb: to.Ptr[int32](0),
		// 									StageName: to.Ptr("teststage"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/4bed37fd-19a1-4d31-8b44-40267555bec5/resourceGroups/edgemarketplace-rg/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/edgemarketplace-demo/providers/Microsoft.EdgeMarketplace/publishers/canonical/offers/0001-com-ubuntu-pro-jammy"),
		// 				Name: to.Ptr("0001-com-ubuntu-pro-jammy"),
		// 				Type: to.Ptr("Microsoft.EdgeMarketplace/publishers/offers"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
