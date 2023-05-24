package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Catalogs_CreateAdo.json
func ExampleCatalogsClient_BeginCreateOrUpdate_catalogsCreateOrUpdateAdo() {
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
			AdoGit: &armdevcenter.GitCatalog{
				Path:             to.Ptr("/templates"),
				Branch:           to.Ptr("main"),
				SecretIdentifier: to.Ptr("https://contosokv.vault.azure.net/secrets/CentralRepoPat"),
				URI:              to.Ptr("https://contoso@dev.azure.com/contoso/contosoOrg/_git/centralrepo-fakecontoso"),
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
