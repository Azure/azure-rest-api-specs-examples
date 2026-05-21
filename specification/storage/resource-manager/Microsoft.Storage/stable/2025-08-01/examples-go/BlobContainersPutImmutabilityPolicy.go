package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2025-08-01/BlobContainersPutImmutabilityPolicy.json
func ExampleBlobContainersClient_CreateOrUpdateImmutabilityPolicy_createOrUpdateImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobContainersClient().CreateOrUpdateImmutabilityPolicy(ctx, "res1782", "sto7069", "container6397", armstorage.ImmutabilityPolicy{
		Properties: &armstorage.ImmutabilityPolicyProperty{
			AllowProtectedAppendWrites:            to.Ptr(true),
			ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](3),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.BlobContainersClientCreateOrUpdateImmutabilityPolicyResponse{
	// 	ImmutabilityPolicy: armstorage.ImmutabilityPolicy{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers/immutabilityPolicies"),
	// 		Etag: to.Ptr("\"8d59f830cb130e5\""),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res1782/providers/Microsoft.Storage/storageAccounts/sto7069/blobServices/default/containers/container6397/immutabilityPolicies/default"),
	// 		Properties: &armstorage.ImmutabilityPolicyProperty{
	// 			AllowProtectedAppendWrites: to.Ptr(true),
	// 			ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](3),
	// 			State: to.Ptr(armstorage.ImmutabilityPolicyStateUnlocked),
	// 		},
	// 	},
	// }
}
