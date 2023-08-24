package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2023-04-01/examples/SyncIdentityProviders_Delete.json
func ExampleSyncIdentityProvidersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSyncIdentityProvidersClient().Delete(ctx, "resourceGroup", "resourceName", "childResourceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
