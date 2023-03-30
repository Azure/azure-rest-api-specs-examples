package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c583b05741fadfdca116be3b9ccb1c4be8a73258/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Catalogs_Patch.json
func ExampleCatalogsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCatalogsClient().BeginUpdate(ctx, "rg1", "Contoso", "CentralCatalog", armdevcenter.CatalogUpdate{
		Properties: &armdevcenter.CatalogUpdateProperties{
			GitHub: &armdevcenter.GitCatalog{
				Path: to.Ptr("/environments"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Catalog = armdevcenter.Catalog{
	// 	Name: to.Ptr("CentralCatalog"),
	// 	Type: to.Ptr("Microsoft.DevCenter/devcenters/catalogs"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/catalogs/CentralCatalog"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 		CreatedBy: to.Ptr("User1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Properties: &armdevcenter.CatalogProperties{
	// 		GitHub: &armdevcenter.GitCatalog{
	// 			Path: to.Ptr("/environments"),
	// 			Branch: to.Ptr("main"),
	// 			SecretIdentifier: to.Ptr("https://contosokv.vault.azure.net/secrets/CentralRepoPat"),
	// 			URI: to.Ptr("https://github.com/Contoso/centralrepo-fake.git"),
	// 		},
	// 		LastSyncTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:28:00.314Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		SyncState: to.Ptr(armdevcenter.CatalogSyncStateSucceeded),
	// 	},
	// }
}
