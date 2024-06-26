package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c583b05741fadfdca116be3b9ccb1c4be8a73258/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/Catalogs_CreateGitHub.json
func ExampleCatalogsClient_BeginCreateOrUpdate_catalogsCreateOrUpdateGitHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
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
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
