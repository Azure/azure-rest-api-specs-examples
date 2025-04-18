package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/97ee23a6db6078abcbec7b75bf9af8c503e9bb8b/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/StorageAccountDeleteObjectReplicationPolicy.json
func ExampleObjectReplicationPoliciesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewObjectReplicationPoliciesClient().Delete(ctx, "res6977", "sto2527", "{objectReplicationPolicy-Id}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
