package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Catalogs_List.json
func ExampleCatalogsClient_NewListByDevCenterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogsClient().NewListByDevCenterPager("rg1", "Contoso", &armdevcenter.CatalogsClientListByDevCenterOptions{Top: nil})
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
		// page.CatalogListResult = armdevcenter.CatalogListResult{
		// 	Value: []*armdevcenter.Catalog{
		// 		{
		// 			Name: to.Ptr("CentralCatalog"),
		// 			Type: to.Ptr("Microsoft.DevCenter/devcenters/catalogs"),
		// 			ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/catalogs/CentralCatalog"),
		// 			SystemData: &armdevcenter.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 				CreatedBy: to.Ptr("User1"),
		// 				CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("User1"),
		// 				LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 			},
		// 			Properties: &armdevcenter.CatalogProperties{
		// 				GitHub: &armdevcenter.GitCatalog{
		// 					Path: to.Ptr("/templates"),
		// 					Branch: to.Ptr("main"),
		// 					SecretIdentifier: to.Ptr("https://contosokv.vault.azure.net/secrets/CentralRepoPat"),
		// 					URI: to.Ptr("https://github.com/Contoso/centralrepo-fake.git"),
		// 				},
		// 				LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:28:00.314Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
		// 				SyncState: to.Ptr(armdevcenter.CatalogSyncStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
