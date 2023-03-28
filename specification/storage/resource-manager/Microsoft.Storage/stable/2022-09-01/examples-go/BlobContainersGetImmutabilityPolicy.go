package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersGetImmutabilityPolicy.json
func ExampleBlobContainersClient_GetImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobContainersClient().GetImmutabilityPolicy(ctx, "res5221", "sto9177", "container3489", &armstorage.BlobContainersClientGetImmutabilityPolicyOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ImmutabilityPolicy = armstorage.ImmutabilityPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers/immutabilityPolicies"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res5221/providers/Microsoft.Storage/storageAccounts/sto9177/blobServices/default/containers/container3489/immutabilityPolicies/default"),
	// 	Etag: to.Ptr("\"8d59f828e64b75c\""),
	// 	Properties: &armstorage.ImmutabilityPolicyProperty{
	// 		AllowProtectedAppendWrites: to.Ptr(true),
	// 		ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](5),
	// 		State: to.Ptr(armstorage.ImmutabilityPolicyStateUnlocked),
	// 	},
	// }
}
