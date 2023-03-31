package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d89e212da2cc34206fe711f44dfcb6f8fdece2a1/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-09-04/examples/SyncSets_Delete.json
func ExampleSyncSetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSyncSetsClient().Delete(ctx, "resourceGroup", "resourceName", "childResourceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
