package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/ProjectCatalogs_List.json
func ExampleProjectCatalogsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectCatalogsClient().NewListPager("rg1", "DevProject", nil)
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
		// page = armdevcenter.ProjectCatalogsClientListResponse{
		// 	CatalogListResult: armdevcenter.CatalogListResult{
		// 		Value: []*armdevcenter.Catalog{
		// 			{
		// 				Name: to.Ptr("CentralCatalog"),
		// 				Type: to.Ptr("Microsoft.DevCenter/projects/catalogs"),
		// 				ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/Contoso/catalogs"),
		// 				Properties: &armdevcenter.CatalogProperties{
		// 					ConnectionState: to.Ptr(armdevcenter.CatalogConnectionStateConnected),
		// 					GitHub: &armdevcenter.GitCatalog{
		// 						Path: to.Ptr("/templates"),
		// 						Branch: to.Ptr("main"),
		// 						SecretIdentifier: to.Ptr("https://contosokv.vault.azure.net/secrets/CentralRepoPat"),
		// 						URI: to.Ptr("https://github.com/Contoso/centralrepo-fake.git"),
		// 					},
		// 					LastConnectionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:28:00.314Z"); return t}()),
		// 					LastSyncStats: &armdevcenter.SyncStats{
		// 						Added: to.Ptr[int32](1),
		// 						Removed: to.Ptr[int32](1),
		// 						SynchronizationErrors: to.Ptr[int32](1),
		// 						Unchanged: to.Ptr[int32](1),
		// 						Updated: to.Ptr[int32](1),
		// 						ValidationErrors: to.Ptr[int32](1),
		// 					},
		// 					LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:28:00.314Z"); return t}()),
		// 					ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
		// 					SyncState: to.Ptr(armdevcenter.CatalogSyncStateSucceeded),
		// 				},
		// 				SystemData: &armdevcenter.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 					CreatedBy: to.Ptr("User1"),
		// 					CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("User1"),
		// 					LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
