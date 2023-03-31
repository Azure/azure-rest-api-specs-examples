package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetPrivateStores.json
func ExamplePrivateStoreClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmarketplace.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateStoreClient().NewListPager(&armmarketplace.PrivateStoreClientListOptions{UseCache: nil})
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
		// page.PrivateStoreList = armmarketplace.PrivateStoreList{
		// 	Value: []*armmarketplace.PrivateStore{
		// 		{
		// 			Name: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
		// 			Type: to.Ptr("Microsoft.Marketplace/privateStores"),
		// 			ID: to.Ptr("/providers/Microsoft.Marketplace/privateStores/a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
		// 			Properties: &armmarketplace.PrivateStoreProperties{
		// 				Availability: to.Ptr(armmarketplace.AvailabilityEnabled),
		// 				Branding: map[string]*string{
		// 					"color": to.Ptr("blue"),
		// 					"iconUrl": to.Ptr("https://some-images.someDomail.com/image/stroeIcon.12345678-4321"),
		// 				},
		// 				CollectionIDs: []*string{
		// 					to.Ptr("ab80ed4-c5de-4593-b3c1-c46d4a3a56c7"),
		// 					to.Ptr("a09107d4-c585-4594-b3c6-c46d4a3a56c8")},
		// 					ETag: to.Ptr("\"9301f4fd-0000-0100-0000-5e248b350332\""),
		// 					IsGov: to.Ptr(false),
		// 					NotificationsSettings: &armmarketplace.NotificationsSettingsProperties{
		// 						Recipients: []*armmarketplace.Recipient{
		// 							{
		// 								DisplayName: to.Ptr("John Doe"),
		// 								EmailAddress: to.Ptr("john_doe@microsoft.com"),
		// 								PrincipalID: to.Ptr("6d583005-403b-407a-8ac0-c4af72b47ce9"),
		// 							},
		// 							{
		// 								DisplayName: to.Ptr("Jane Doe"),
		// 								EmailAddress: to.Ptr("jane_doe@microsoft.com"),
		// 								PrincipalID: to.Ptr("c5b680d4-aac2-4940-9e1c-399454056ff2"),
		// 						}},
		// 						SendToAllMarketplaceAdmins: to.Ptr(true),
		// 					},
		// 					PrivateStoreID: to.Ptr("a0e28e55-90c4-41d8-8e34-bb7ef7775406"),
		// 					TenantID: to.Ptr("f686d426-123a-42db-81b7-ab578e110ccd"),
		// 				},
		// 		}},
		// 	}
	}
}
