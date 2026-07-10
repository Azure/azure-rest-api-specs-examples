package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/Catalogs_CreateGitHub.json
func ExampleCatalogsClient_BeginCreateOrUpdate_catalogsCreateOrUpdateGitHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCatalogsClient().BeginCreateOrUpdate(ctx, "rg1", "Contoso", "CentralCatalog", armdevcenter.Catalog{
		Properties: &armdevcenter.CatalogProperties{
			GitHub: &armdevcenter.GitCatalog{
				Path:             to.Ptr("/templates"),
				Branch:           to.Ptr("main"),
				SecretIdentifier: to.Ptr("https://contosokv.vault.azure.net/secrets/CentralRepoPat"),
				URI:              to.Ptr("https://github.com/Contoso/centralrepo-fake.git"),
			},
			SyncType:                   to.Ptr(armdevcenter.CatalogSyncTypeManual),
			AutoImageBuildEnableStatus: to.Ptr(armdevcenter.CatalogAutoImageBuildEnableStatusEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.CatalogsClientCreateOrUpdateResponse{
	// 	Catalog: armdevcenter.Catalog{
	// 		Name: to.Ptr("CentralCatalog"),
	// 		Type: to.Ptr("Microsoft.DevCenter/devcenters/catalogs"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/catalogs/CentralCatalog"),
	// 		Properties: &armdevcenter.CatalogProperties{
	// 			ConnectionState: to.Ptr(armdevcenter.CatalogConnectionStateConnected),
	// 			GitHub: &armdevcenter.GitCatalog{
	// 				Path: to.Ptr("/templates"),
	// 				Branch: to.Ptr("main"),
	// 				SecretIdentifier: to.Ptr("https://contosokv.vault.azure.net/secrets/CentralRepoPat"),
	// 				URI: to.Ptr("https://github.com/Contoso/centralrepo-fake.git"),
	// 			},
	// 			LastSyncStats: &armdevcenter.SyncStats{
	// 				Added: to.Ptr[int32](0),
	// 				Removed: to.Ptr[int32](0),
	// 				SynchronizationErrors: to.Ptr[int32](0),
	// 				Unchanged: to.Ptr[int32](0),
	// 				Updated: to.Ptr[int32](0),
	// 				ValidationErrors: to.Ptr[int32](0),
	// 			},
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateAccepted),
	// 			SyncState: to.Ptr(armdevcenter.CatalogSyncStateSucceeded),
	// 			SyncType: to.Ptr(armdevcenter.CatalogSyncTypeManual),
	// 			AutoImageBuildEnableStatus: to.Ptr(armdevcenter.CatalogAutoImageBuildEnableStatusEnabled),
	// 		},
	// 		SystemData: &armdevcenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 			CreatedBy: to.Ptr("User1"),
	// 			CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("User1"),
	// 			LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
