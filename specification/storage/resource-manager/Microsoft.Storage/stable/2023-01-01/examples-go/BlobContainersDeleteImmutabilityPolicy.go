package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/BlobContainersDeleteImmutabilityPolicy.json
func ExampleBlobContainersClient_DeleteImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobContainersClient().DeleteImmutabilityPolicy(ctx, "res1581", "sto9621", "container4910", "8d59f81a7fa7be0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ImmutabilityPolicy = armstorage.ImmutabilityPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers/immutabilityPolicies"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res1581/providers/Microsoft.Storage/storageAccounts/sto9621/blobServices/default/containers/container4910/immutabilityPolicies/default"),
	// 	Etag: to.Ptr("\"8d59f81a87b40c0\""),
	// 	Properties: &armstorage.ImmutabilityPolicyProperty{
	// 		ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](0),
	// 		State: to.Ptr(armstorage.ImmutabilityPolicyStateUnlocked),
	// 	},
	// }
}
