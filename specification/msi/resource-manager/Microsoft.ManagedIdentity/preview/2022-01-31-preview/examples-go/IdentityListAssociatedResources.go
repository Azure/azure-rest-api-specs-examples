package armmsi_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/examples/IdentityListAssociatedResources.json
func ExampleUserAssignedIdentitiesClient_NewListAssociatedResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmsi.NewUserAssignedIdentitiesClient("1cscb752-d7c9-463f-9731-fd31edada74a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListAssociatedResourcesPager("testrg",
		"testid",
		&armmsi.UserAssignedIdentitiesClientListAssociatedResourcesOptions{Filter: to.Ptr("contains(name, 'test')"),
			Orderby:   to.Ptr("name asc"),
			Top:       to.Ptr[int32](10),
			Skip:      to.Ptr[int32](1),
			Skiptoken: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
