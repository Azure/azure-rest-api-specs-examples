package armstorageimportexport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListLocations.json
func ExampleLocationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageimportexport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLocationsClient().NewListPager(&armstorageimportexport.LocationsClientListOptions{AcceptLanguage: nil})
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
		// page.LocationsResponse = armstorageimportexport.LocationsResponse{
		// 	Value: []*armstorageimportexport.Location{
		// 		{
		// 			Name: to.Ptr("Australia East"),
		// 			Type: to.Ptr("Microsoft.ImportExport/locations"),
		// 			ID: to.Ptr("/providers/Microsoft.ImportExport/locations/australiaeast"),
		// 			Properties: &armstorageimportexport.LocationProperties{
		// 				AdditionalShippingInformation: to.Ptr(""),
		// 				AlternateLocations: []*string{
		// 					to.Ptr("/providers/Microsoft.ImportExport/locations/australiaeast")},
		// 					City: to.Ptr("Macquarie Park"),
		// 					CountryOrRegion: to.Ptr("Australia"),
		// 					Phone: to.Ptr("612 0000 0000"),
		// 					PostalCode: to.Ptr("2113"),
		// 					RecipientName: to.Ptr("Windows Azure Import/Export Service"),
		// 					StateOrProvince: to.Ptr("NSW"),
		// 					StreetAddress1: to.Ptr("Customer-A c/o NEXTDC Ltd, Delivery code ABCDEFG, 4 Eden Park Drive"),
		// 					StreetAddress2: to.Ptr(""),
		// 					SupportedCarriers: []*string{
		// 						to.Ptr("DHL")},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("Australia Southeast"),
		// 					Type: to.Ptr("Microsoft.ImportExport/locations"),
		// 					ID: to.Ptr("/providers/Microsoft.ImportExport/locations/australiasoutheast"),
		// 					Properties: &armstorageimportexport.LocationProperties{
		// 						AdditionalShippingInformation: to.Ptr(""),
		// 						AlternateLocations: []*string{
		// 							to.Ptr("/providers/Microsoft.ImportExport/locations/australiasoutheast")},
		// 							City: to.Ptr("Melbourne"),
		// 							CountryOrRegion: to.Ptr("Australia"),
		// 							Phone: to.Ptr("61 0 0000 0000"),
		// 							PostalCode: to.Ptr("3207"),
		// 							RecipientName: to.Ptr("Microsoft Azure Import/Export Service"),
		// 							StateOrProvince: to.Ptr("Melbourne"),
		// 							StreetAddress1: to.Ptr("Microsoft, c/o NEXTDC Ltd, Delivery code ABCDEFG, 826-830 Lorimer St, Port"),
		// 							StreetAddress2: to.Ptr(""),
		// 							SupportedCarriers: []*string{
		// 								to.Ptr("DHL")},
		// 							},
		// 					}},
		// 				}
	}
}
