package armstorageimportexport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/GetLocation.json
func ExampleLocationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageimportexport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().Get(ctx, "West US", &armstorageimportexport.LocationsClientGetOptions{AcceptLanguage: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Location = armstorageimportexport.Location{
	// 	Name: to.Ptr("West US"),
	// 	Type: to.Ptr("Microsoft.ImportExport/locations"),
	// 	ID: to.Ptr("/providers/Microsoft.ImportExport/locations/westus"),
	// 	Properties: &armstorageimportexport.LocationProperties{
	// 		AdditionalShippingInformation: to.Ptr(""),
	// 		AlternateLocations: []*string{
	// 			to.Ptr("/providers/Microsoft.ImportExport/locations/westus")},
	// 			City: to.Ptr("Santa Clara"),
	// 			CountryOrRegion: to.Ptr("USA"),
	// 			Phone: to.Ptr("408 352 7600"),
	// 			PostalCode: to.Ptr("95050"),
	// 			RecipientName: to.Ptr("Microsoft Azure Import/Export Service"),
	// 			StateOrProvince: to.Ptr("CA"),
	// 			StreetAddress1: to.Ptr("2045 Lafayette Street"),
	// 			StreetAddress2: to.Ptr(""),
	// 			SupportedCarriers: []*string{
	// 				to.Ptr("FedEx")},
	// 			},
	// 		}
}
