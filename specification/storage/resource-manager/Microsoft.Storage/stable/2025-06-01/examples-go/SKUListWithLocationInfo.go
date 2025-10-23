package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4e9df3afd38a1cfa00a5d49419dce51bd014601f/specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/SKUListWithLocationInfo.json
func ExampleSKUsClient_NewListPager_skuListWithLocationInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListPager(nil)
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
		// page.SKUListResult = armstorage.SKUListResult{
		// 	Value: []*armstorage.SKUInformation{
		// 		{
		// 			Name: to.Ptr(armstorage.SKUNamePremiumLRS),
		// 			Capabilities: []*armstorage.SKUCapability{
		// 				{
		// 					Name: to.Ptr("supportschangenotification"),
		// 					Value: to.Ptr("true"),
		// 				},
		// 				{
		// 					Name: to.Ptr("supportsfileencryption"),
		// 					Value: to.Ptr("false"),
		// 				},
		// 				{
		// 					Name: to.Ptr("supportshoeboxcapacitymetrics"),
		// 					Value: to.Ptr("false"),
		// 				},
		// 				{
		// 					Name: to.Ptr("supportsnetworkacls"),
		// 					Value: to.Ptr("false"),
		// 			}},
		// 			Kind: to.Ptr(armstorage.KindFileStorage),
		// 			LocationInfo: []*armstorage.SKUInformationLocationInfoItem{
		// 				{
		// 					Location: to.Ptr("centraluseuap"),
		// 					Zones: []*string{
		// 						to.Ptr("1"),
		// 						to.Ptr("2"),
		// 						to.Ptr("3")},
		// 				}},
		// 				Locations: []*string{
		// 					to.Ptr("centraluseuap")},
		// 					ResourceType: to.Ptr("storageAccounts"),
		// 					Restrictions: []*armstorage.Restriction{
		// 					},
		// 					Tier: to.Ptr(armstorage.SKUTierPremium),
		// 			}},
		// 		}
	}
}
