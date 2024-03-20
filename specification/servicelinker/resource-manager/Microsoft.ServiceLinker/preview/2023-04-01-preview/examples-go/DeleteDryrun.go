package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c16ce913afbdaa073e8ca5e480f3b465db2de542/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/DeleteDryrun.json
func ExampleLinkersClient_DeleteDryrun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewLinkersClient().DeleteDryrun(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "dryrunName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
